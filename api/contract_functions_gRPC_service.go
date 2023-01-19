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
	"net"
)

type ContractIntegrationRPC struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedContractIntegrationServer
	TransactionService *contracts.TransactionClient
}


func NewContractIntegrationRPCService(port int, opts []grpc.ServerOption, transactionService *contracts.TransactionClient) (*ContractIntegrationRPC, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &ContractIntegrationRPC{Server: grpc.NewServer(opts...), Channel: make(chan string), TransactionService: transactionService}
	pb.RegisterContractIntegrationServer(server.Server, server)
	log.Println("GRPC Server registered")
	grpc_health_v1.RegisterHealthServer(server.Server, server)
	log.Println("HealthServer registered")
	go func() {
		log.Println("Contract Implementation Service serving clients now")
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


func (cRPC *ContractIntegrationRPC) BuildClaimTransaction(ctx context.Context, req *pb.ClaimRequest) (*pb.MintResponse, error) {
	args := [][]byte{}

	nonce, nonceErr := utils.GetNonceBytes()
	if nonceErr != nil {
		return &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: nonceErr.Error()}, nil
	}

	tokenId := fmt.Sprintf("%d", req.TokenId)

	args = append(args, nonce)
	args = append(args, []byte(tokenId))

	log.Printf("%+v\n", args)

	_, err := cRPC.TransactionService.Client.ConstructTransaction(ctx, &pb.TransactionRequest{
		SenderInHash:    true,
		MessageSender:   req.MessageSender,
		FunctionName:    "mintArtie",
		Args:            args,
		ContractAddress: req.ContractAddress,
		Value:           "0",
	})

	if err != nil {
		return &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: err.Error()}, nil
	}

	return &pb.MintResponse{Status: pb.Code_CODE_SUCCESS}, nil
}

func (cRPC *ContractIntegrationRPC) BuildMintTransaction(ctx context.Context, req *pb.MintRequest) (*pb.MintResponse, error) {
	args := [][]byte{}

	nonce, nonceErr := utils.GetNonceBytes()
	if nonceErr != nil {
		return &pb.MintResponse{Status: pb.Code_CODE_INTERNAL_SERVER_ERROR, Message: nonceErr.Error()}, nil
	}

	tokensRequested := fmt.Sprintf("%d", req.NumberOfTokens)
	transactionNumber := fmt.Sprintf("%d", req.TransactionNumber)

	args = append(args, nonce)
	args = append(args, []byte(tokensRequested))
	args = append(args, []byte(transactionNumber))

	_, err := cRPC.TransactionService.Client.ConstructTransaction(ctx, &pb.TransactionRequest{
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

func (cRPC *ContractIntegrationRPC) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (cRPC *ContractIntegrationRPC) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}
