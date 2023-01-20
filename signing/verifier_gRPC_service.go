package signing

import (
	"context"
	pb "contract-service/proto"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
)

type VerifierRPCService struct {
	Server *grpc.Server
	Channel chan string
	pb.UnimplementedVerificationServiceServer
	Handler SigningService
}


func NewVerifierServer(port int, opts []grpc.ServerOption, handler SigningService) (*VerifierRPCService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", port))
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully listening on port: %d\n", port)

	server := &VerifierRPCService{Server: grpc.NewServer(opts...), Channel: make(chan string), Handler: handler}
	pb.RegisterVerificationServiceServer(server.Server, server)
	log.Println("GRPC Server registered")
	grpc_health_v1.RegisterHealthServer(server.Server, server)
	log.Println("HealthServer registered")
	go func() {
		log.Println("SignerServer serving clients now")
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


func (vRPC *VerifierRPCService) Verify(ctx context.Context, req *pb.SignatureVerificationRequest) (*pb.SignatureVerificationResponse, error) {
	log.Printf("Verifying\ncontext: %+v\nmessage: %s\nsignature: %s\naddress: %s\n", ctx, req.Message, req.Signature, req.Address)
	if err := vRPC.Handler.Verify(req.Message, req.Signature, req.Address); err != nil {
		log.Println(err)
		return &pb.SignatureVerificationResponse{Success: false}, nil
	}
	return &pb.SignatureVerificationResponse{Success: true}, nil
}

func (vRPC *VerifierRPCService) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	log.Println("Health check ping to: " + req.Service)
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}

func (vRPC *VerifierRPCService) Watch(req *grpc_health_v1.HealthCheckRequest, server grpc_health_v1.Health_WatchServer) error {
	log.Println("Health watcher for: " + req.Service)
	return server.SendMsg(&grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING})
}
