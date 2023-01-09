package api

import (
	"context"
	"contract-service/contracts"
	pb "contract-service/proto"
	"contract-service/utils"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"math/big"
	"net"
)

type MintingRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedMintServiceServer
	TransactionService contracts.TransactionClient
}


func NewMintingRPCService(port int, opts []grpc.ServerOption, transactionService contracts.TransactionClient) (*MintingRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &MintingRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), TransactionService: transactionService}
	pb.RegisterMintServiceServer(server.Server, server)
	log.Println("GRPC Server registered")
	grpc_health_v1.RegisterHealthServer(server.Server, server)
	log.Println("HealthServer registered")
	go func() {
		log.Println("SignerServer serving clients now")
		defer server.Server.GracefulStop()
		serviceErr := server.Server.Serve(lis)
		server.Channel <- serviceErr.Error()
	}()
	return server, nil
}


func (mRPC *MintingRPCService) BuildMintTransaction(ctx context.Context, req *pb.MintRequest) (*pb.MintResponse, error) {
	args := [][]byte{}

	nonce, nonceErr := utils.GetNonceBytes()
	if nonceErr != nil {
		return &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: nonceErr.Error()}, nil
	}

	tokensRequested := new(big.Int).SetInt64(req.NumberOfTokens)
	transactionNumber := new(big.Int).SetInt64(req.TransactionNumber)

	args = append(args, nonce)
	args = append(args, tokensRequested.Bytes())
	args = append(args, transactionNumber.Bytes())

	_, err := mRPC.TransactionService.Client.ConstructTransaction(ctx, &pb.TransactionRequest{
		SenderInHash:    true,
		MessageSender:   req.MessageSender,
		FunctionName:    "mint",
		Args:            args,
		ContractAddress: req.ContractAddress,
		Value:           req.Value,
	})

	if err != nil {
		return &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: err.Error()}, nil
	}

	return &pb.MintResponse{Status: pb.Code_CODE_SUCCESS}, nil
}

func (mRPC *MintingRPCService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (mRPC *MintingRPCService) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}
