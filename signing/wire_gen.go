// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package signing

import (
	"bitbucket.org/artie_inc/contract-service/storage"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitializeSigningServer(host int, opts []grpc.ServerOption, config storage.PrivateKeyConfig) (*SignerRPCService, error) {
	signingService := NewSigningService()
	privateKeyRepository, err := storage.NewPrivateKeyRepository(config)
	if err != nil {
		return nil, err
	}
	signerRPCService, err := NewSignerServer(host, opts, signingService, privateKeyRepository)
	if err != nil {
		return nil, err
	}
	return signerRPCService, nil
}

func InitializeVerifierServer(host int, opts []grpc.ServerOption) (*VerifierRPCService, error) {
	signingService := NewSigningService()
	verifierRPCService, err := NewVerifierServer(host, opts, signingService)
	if err != nil {
		return nil, err
	}
	return verifierRPCService, nil
}
