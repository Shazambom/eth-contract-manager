// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package contracts

import (
	"contract-service/proto"
	"contract-service/storage"
	"github.com/aws/aws-sdk-go/aws"
	"google.golang.org/grpc"
)

// Injectors from wire.go:

func InitializeContractTransactionHandler(writer storage.RedisWriter, client pb.SigningServiceClient, tableName string, cfg ...*aws.Config) (ContractTransactionHandler, error) {
	contractRepository, err := storage.NewContractRepository(tableName, cfg...)
	if err != nil {
		return nil, err
	}
	contractTransactionHandler := NewContractTransactionHandler(writer, contractRepository, client)
	return contractTransactionHandler, nil
}

func InitializeContractManagerHandler(tableName string, cfg ...*aws.Config) (ContractManagerHandler, error) {
	contractRepository, err := storage.NewContractRepository(tableName, cfg...)
	if err != nil {
		return nil, err
	}
	contractManagerHandler := NewContractManagerHandler(contractRepository)
	return contractManagerHandler, nil
}

func InitializeTransactionServer(port int, opts []grpc.ServerOption, writer storage.RedisWriter, client pb.SigningServiceClient, tableName string, cfg ...*aws.Config) (*TransactionRPCService, error) {
	contractTransactionHandler, err := InitializeContractTransactionHandler(writer, client, tableName, cfg...)
	if err != nil {
		return nil, err
	}
	transactionRPCService, err := NewTransactionServer(port, opts, contractTransactionHandler)
	if err != nil {
		return nil, err
	}
	return transactionRPCService, nil
}

func InitializeContractServer(port int, opts []grpc.ServerOption, tableName string, cfg ...*aws.Config) (*ContractRPCService, error) {
	contractManagerHandler, err := InitializeContractManagerHandler(tableName, cfg...)
	if err != nil {
		return nil, err
	}
	contractRPCService, err := NewContractServer(port, opts, contractManagerHandler)
	if err != nil {
		return nil, err
	}
	return contractRPCService, nil
}
