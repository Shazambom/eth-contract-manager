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

//TODO Add Ping route to gRPC to check if service is alive

type TransactionRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedTransactionServiceServer
	TransactionManager ContractTransactionHandler
}

func NewTransactionServer(port int, opts []grpc.ServerOption, handler ContractTransactionHandler) (*TransactionRPCService, error) {
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

func (ts *TransactionRPCService) ConstructTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.Error, error) {
	contract := &storage.Contract{}
	contract.FromRPC(req.Contract)

	isNotValidErr := ts.TransactionManager.CheckIfValidRequest(ctx, req.MessageSender, int(req.NumRequested), contract)
	if isNotValidErr != nil {
		return &pb.Error{Err: isNotValidErr.Error()}, isNotValidErr
	}

	token, tokenErr := ts.TransactionManager.BuildTransaction(ctx, req.MessageSender, req.FunctionName, int(req.NumRequested), req.Args, contract)
	if tokenErr != nil {
		return &pb.Error{Err: tokenErr.Error()}, tokenErr
	}

	storeErr := ts.TransactionManager.StoreToken(ctx, token, contract)
	if storeErr != nil {
		return &pb.Error{Err: storeErr.Error()}, storeErr
	}
	return nil, nil
}

func (ts *TransactionRPCService) Check(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &pb.HealthCheckResponse{Status: pb.HealthCheckResponse_SERVING}, nil
}

func (ts *TransactionRPCService) Watch(req *pb.HealthCheckRequest, server pb.TransactionService_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&pb.HealthCheckResponse{Status: pb.HealthCheckResponse_SERVING})
}
