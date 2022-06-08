//+build wireinject

package contracts

import (
	pb "contract-service/proto"
	"contract-service/storage"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeContractTransactionHandler(writer storage.RedisWriter, client pb.SigningServiceClient, tableName string, cfg ...*aws.Config) (ContractTransactionHandler, error) {
	wire.Build(NewContractTransactionHandler, storage.NewContractRepository)
	return &ContractManagerService{}, nil
}

func InitializeContractManagerHandler(tableName string, cfg ...*aws.Config) (ContractManagerHandler, error) {
	wire.Build(NewContractManagerHandler, storage.NewContractRepository)
	return &ContractManagerService{}, nil
}

func InitializeTransactionServer(port int, opts []grpc.ServerOption, writer storage.RedisWriter, client pb.SigningServiceClient, tableName string, cfg ...*aws.Config) (*TransactionRPCService, error) {
	wire.Build(NewTransactionServer, InitializeContractTransactionHandler)
	return &TransactionRPCService{}, nil
}

func InitializeContractServer(port int, opts []grpc.ServerOption,tableName string, cfg ...*aws.Config) (*ContractRPCService, error) {
	wire.Build(NewContractServer, InitializeContractManagerHandler)
	return &ContractRPCService{}, nil
}