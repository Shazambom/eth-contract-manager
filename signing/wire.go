//go:build wireinject
// +build wireinject

package signing

import (
	"bitbucket.org/artie_inc/contract-service/storage"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeSigningServer(host int, opts []grpc.ServerOption, config storage.PrivateKeyConfig) (*SignerRPCService, error) {
	wire.Build(NewSignerServer, NewSigningService, storage.NewPrivateKeyRepository)
	return &SignerRPCService{}, nil
}

func InitializeVerifierServer(host int, opts []grpc.ServerOption) (*VerifierRPCService, error) {
	wire.Build(NewVerifierServer, NewSigningService)
	return &VerifierRPCService{}, nil
}
