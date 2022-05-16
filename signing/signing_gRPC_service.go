package signing

import (
	"context"
	pb "contract-service/proto"
	"contract-service/storage"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

type SignerRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedSigningServiceServer
	Handler SigningService
	Repo storage.PrivateKeyRepository
	KeyManager KeyManagerService
}


func NewSignerServer(port int, opts []grpc.ServerOption, handler SigningService, repo storage.PrivateKeyRepository, keyManager KeyManagerService) (*SignerRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &SignerRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), Handler: handler, Repo: repo, KeyManager: keyManager}
	pb.RegisterSigningServiceServer(server.Server, server)
	log.Println("GRPC Server registered")
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
	privKey, addr, err := sRPC.KeyManager.GenerateKey()
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
