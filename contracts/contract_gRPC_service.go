package contracts

import (
	"context"
	pb "contract-service/proto"
	"contract-service/storage"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
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
	grpc_health_v1.RegisterHealthServer(server.Server, server)
	log.Println("HealthServer registered")
	go func() {
		log.Println("ContractManager serving clients now")
		defer server.Server.GracefulStop()
		serviceErr := server.Server.Serve(lis)
		if serviceErr != nil {
			server.Channel <- serviceErr.Error()
		} else {
			server.Channel <- "gRPC Service has stopped"
		}
	}()
	return server, nil
}

func (cs *ContractRPCService) Get(ctx context.Context, address *pb.Address) (*pb.Contract, error) {
	log.Printf("Getting contract from address: %s", address.Address)
	contract, err := cs.ContractManager.GetContract(ctx, address.Address)
	if err != nil {
		return nil, err
	}
	return contract.ToRPC(), nil
}

func (cs *ContractRPCService) Store(ctx context.Context, contract *pb.Contract) (*pb.Empty, error) {
	con := &storage.Contract{}
	con.FromRPC(contract)
	log.Printf("Storing contract: %+v\n", con)
	return &pb.Empty{}, cs.ContractManager.StoreContract(ctx, con)
}

func (cs *ContractRPCService) Delete(ctx context.Context, req *pb.AddressOwner) (*pb.Empty, error) {
	log.Printf("Deleting %s's contract: %s\n", req.Owner, req.Address)
	return &pb.Empty{}, cs.ContractManager.DeleteContract(ctx, req.Address, req.Owner)
}

func (cs *ContractRPCService) List(ctx context.Context, owner *pb.Owner) (*pb.Contracts, error) {
	log.Printf("Getting all %s's contracts\n", owner.Owner)
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

func (cs *ContractRPCService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (cs *ContractRPCService) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}
