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

type ClaimRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedClaimServiceServer
	TransactionService contracts.TransactionClient
}


func NewClaimPRCService(port int, opts []grpc.ServerOption, transactionService contracts.TransactionClient) (*ClaimRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &ClaimRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), TransactionService: transactionService}
	pb.RegisterClaimServiceServer(server.Server, server)
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


func (cRPC *ClaimRPCService) BuildClaimTransaction(ctx context.Context, req *pb.ClaimRequest) (*pb.MintResponse, error) {
	args := [][]byte{}

	nonce, nonceErr := utils.GetNonceBytes()
	if nonceErr != nil {
		return &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: nonceErr.Error()}, nil
	}

	tokenId := new(big.Int).SetInt64(req.TokenId)

	args = append(args, nonce)
	args = append(args, tokenId.Bytes())

	_, err := cRPC.TransactionService.Client.ConstructTransaction(ctx, &pb.TransactionRequest{
		SenderInHash:    true,
		MessageSender:   req.MessageSender,
		FunctionName:    "mint",
		Args:            args,
		ContractAddress: req.ContractAddress,
		Value:           "0",
	})

	if err != nil {
		return &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: err.Error()}, nil
	}

	return &pb.MintResponse{Status: pb.Code_CODE_SUCCESS}, nil
}

func (cRPC *ClaimRPCService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (cRPC *ClaimRPCService) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}
