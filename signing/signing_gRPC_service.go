package signing

import (
	"context"
	pb "contract-service/proto"
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
}


func NewSignerServer(port int, opts []grpc.ServerOption, handler SigningService) (*SignerRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &SignerRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), Handler: handler}
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
	hash, signature, err := sRPC.Handler.SignTxn(req.SigningKey, req.Args)
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
