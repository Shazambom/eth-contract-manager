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
	grpc_health_v1.RegisterHealthServer(server.Server, server)
	log.Println("HealthServer registered")
	go func() {
		log.Println("TransactionService serving clients now")
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

func (ts *TransactionRPCService) ConstructTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.Transaction, error) {
	contract := &storage.Contract{}
	contract.FromRPC(req.Contract)

	token, tokenErr := ts.TransactionManager.BuildTransaction(ctx, req.MessageSender, req.FunctionName, req.Args, contract)
	if tokenErr != nil {
		log.Println(tokenErr.Error())
		return nil, tokenErr
	}

	storeErr := ts.TransactionManager.StoreToken(ctx, token, contract)
	if storeErr != nil {
		log.Println(storeErr.Error())
		return nil, storeErr
	}
	return token.ToRPC(), nil
}

func (ts *TransactionRPCService) GetTransactions(ctx context.Context, address *pb.Address) (*pb.Transactions, error) {
	tokens, err := ts.TransactionManager.GetTransactions(ctx, address.Address)
	if err != nil {
		return nil, err
	}
	txns := &pb.Transactions{Transactions: []*pb.Transaction{}}
	for _, token := range tokens {
		txns.Transactions = append(txns.Transactions, token.ToRPC())
	}
	return txns, nil
}

func (ts *TransactionRPCService) GetAllTransactions(ctx context.Context, address *pb.Address) (*pb.Transactions, error) {
	tokens, err := ts.TransactionManager.GetAllTransactions(ctx, address.Address)
	if err != nil {
		return nil, err
	}
	txns := &pb.Transactions{Transactions: []*pb.Transaction{}}
	for _, token := range tokens {
		txns.Transactions = append(txns.Transactions, token.ToRPC())
	}
	return txns, nil
}

func (ts *TransactionRPCService) CompleteTransaction(ctx context.Context, req *pb.KeyTransactionRequest) (*pb.Empty, error) {
	return &pb.Empty{}, ts.TransactionManager.CompleteTransaction(ctx, req.Address, req.Hash)
}

func (ts *TransactionRPCService) DeleteTransaction(ctx context.Context, req *pb.KeyTransactionRequest) (*pb.Empty, error) {
	return &pb.Empty{}, ts.TransactionManager.DeleteTransaction(ctx, req.Address, req.Hash)
}

func (ts *TransactionRPCService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (ts *TransactionRPCService) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}
