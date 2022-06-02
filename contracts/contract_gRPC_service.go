package contracts

import (
	"context"
	pb "contract-service/proto"
	"contract-service/storage"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type ContractRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedContractManagementServer
	ContractManager ContractManagerHandler
}

func NewContractServer(port int, opts []grpc.ServerOption, handler ContractManagerHandler) (*ContractRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &ContractRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), ContractManager: handler}
	pb.RegisterContractManagementServer(server.Server, server)
	log.Println("GRPC Server registered")
	go func() {
		log.Println("SignerServer serving clients now")
		defer server.Server.GracefulStop()
		serviceErr := server.Server.Serve(lis)
		server.Channel <- serviceErr.Error()
	}()
	return server, nil
}

func (cs *ContractRPCService) Get(ctx context.Context, address *pb.Address) (*pb.Contract, error) {
	contract, err := cs.ContractManager.GetContract(ctx, address.Address)
	if err != nil {
		return nil, err
	}
	return contract.ToRPC(), nil
}

func (cs *ContractRPCService) Store(ctx context.Context, contract *pb.Contract) (*pb.Error, error) {
	con := &storage.Contract{}
	con.FromRPC(contract)
	//TODO vvvvv_Can this just be moved to the return statement and get rid of all this error logic_vvvvv
	err := cs.ContractManager.StoreContract(ctx, con)
	if err != nil {
		return &pb.Error{Err: err.Error()}, err
	}
	return nil, nil
}

func (cs *ContractRPCService) Delete(ctx context.Context, req *pb.AddressOwner) (*pb.Error, error) {
	err := cs.ContractManager.DeleteContract(ctx, req.Address, req.Owner)
	if err != nil {
		return &pb.Error{Err: err.Error()}, err
	}
	return nil, nil
}

func (cs *ContractRPCService) List(ctx context.Context, owner *pb.Owner) (*pb.Contracts, error) {
	contracts, err := cs.ContractManager.ListContracts(ctx, owner.Owner)
	if err != nil {
		return nil, err
	}
	protoContracts := []*pb.Contract{}
	for _, contract := range contracts {
		protoContracts = append(protoContracts, contract.ToRPC())
	}
	return &pb.Contracts{Contracts: protoContracts}, nil
}
