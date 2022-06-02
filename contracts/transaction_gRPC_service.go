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

type TransactionRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedTransactionServiceServer
	TransactionManager ContractTransactionHandler
}

func NewTransactionServer(port int, opts []grpc.ServerOption, handler ContractTransactionHandler, repo storage.PrivateKeyRepository) (*TransactionRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &TransactionRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), TransactionManager: handler}
	pb.RegisterTransactionServiceServer(server.Server, server)
	log.Println("GRPC Server registered")
	go func() {
		log.Println("SignerServer serving clients now")
		defer server.Server.GracefulStop()
		serviceErr := server.Server.Serve(lis)
		server.Channel <- serviceErr.Error()
	}()
	return server, nil
}


func (ts *TransactionRPCService) GetContract(ctx context.Context, address *pb.Address) (*pb.Contract, error) {
	contract, err := ts.TransactionManager.GetContract(ctx, address.Address)
	if err != nil {
		return nil, err
	}
	return contract.ToRPC(), nil
}

func (ts *TransactionRPCService) CreateAndStoreTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.Error, error) {
	//TODO build out the logic here:
	//break down the request
	//check if the request is valid
	//build the transaction
	//store the token

	return nil, nil
}
