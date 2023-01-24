package signing

import (
	"context"
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"bitbucket.org/artie_inc/contract-service/storage"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

type SignerRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedSigningServiceServer
	Handler SigningService
	Repo storage.PrivateKeyRepository
}


func NewSignerServer(port int, opts []grpc.ServerOption, handler SigningService, repo storage.PrivateKeyRepository) (*SignerRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &SignerRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), Handler: handler, Repo: repo}
	pb.RegisterSigningServiceServer(server.Server, server)
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


func (sRPC *SignerRPCService) SignTxn(ctx context.Context, req *pb.SignatureRequest) (*pb.SignatureResponse, error) {
	log.Printf("Signing transaction\ncontext: %+v\nargs: %+v\n", ctx, req.Args)
	key, keyErr := sRPC.Repo.GetPrivateKey(ctx, req.ContractAddress)
	if keyErr != nil {
		log.Println(keyErr)
		return nil, keyErr
	}
	hash, signature, err := sRPC.Handler.SignTxn(key, req.Args)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.SignatureResponse{Signature: signature, Hash: hash}, nil
}

func (sRPC *SignerRPCService) BatchSignTxn(ctx context.Context, req *pb.BatchSignatureRequest) (*pb.BatchSignatureResponse, error) {
	log.Printf("Starting batch with context: %+v\nProcessing %d requests\n", ctx, len(req.SignatureRequests))
	responses := []*pb.SignatureResponse{}
	for _, request := range req.SignatureRequests {
		signedResponse, err := sRPC.SignTxn(ctx, request)
		if err != nil {
			continue
		}
		responses = append(responses, signedResponse)
	}
	if len(responses) == 0 {
		return nil, errors.New("none of the signing requests could be fulfilled")
	}
	return &pb.BatchSignatureResponse{SignatureResponses: responses}, nil
}

func (sRPC *SignerRPCService) GenerateNewKey(ctx context.Context, req *pb.KeyManagementRequest)  (*pb.KeyManagementResponse, error) {
	privKey, addr, err := sRPC.Handler.GenerateKey()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	putErr := sRPC.Repo.UpsertPrivateKey(ctx, req.ContractAddress, privKey)
	if putErr != nil {
		log.Println(putErr)
		return nil, putErr
	}
	return &pb.KeyManagementResponse{ContractAddress: req.ContractAddress, PublicKey: addr}, nil
}

func (sRPC *SignerRPCService) DeleteKey(ctx context.Context, req *pb.KeyManagementRequest)  (*pb.KeyManagementResponse, error) {
	err := sRPC.Repo.DeletePrivateKey(ctx, req.ContractAddress)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.KeyManagementResponse{ContractAddress: req.ContractAddress, PublicKey: ""}, nil
}

func (sRPC *SignerRPCService) GetKey(ctx context.Context, req *pb.KeyManagementRequest)  (*pb.KeyManagementResponse, error) {
	privKey, getErr := sRPC.Repo.GetPrivateKey(ctx, req.ContractAddress)
	if getErr != nil {
		log.Println(getErr)
		return nil, getErr
	}

	address, err := sRPC.Handler.PrivateKeyToAddress(privKey)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.KeyManagementResponse{ContractAddress: req.ContractAddress, PublicKey: address}, nil
}

func (sRPC *SignerRPCService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (sRPC *SignerRPCService) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}
