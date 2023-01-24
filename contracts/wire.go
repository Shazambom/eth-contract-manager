//+build wireinject

package contracts

import (
	pb "bitbucket.org/artie_inc/contract-service/proto"
	"bitbucket.org/artie_inc/contract-service/storage"
	"github.com/google/wire"
	"google.golang.org/grpc"
)

func InitializeContractTransactionHandler(client pb.SigningServiceClient, contractConfig storage.ContractConfig, transactionConfig storage.TransactionConfig) (ContractTransactionHandler, error) {
	wire.Build(NewContractTransactionHandler, storage.NewContractRepository, storage.NewTransactionRepo)
	return &ContractManagerService{}, nil
}

func InitializeContractManagerHandler(contractConfig storage.ContractConfig) (ContractManagerHandler, error) {
	wire.Build(NewContractManagerHandler, storage.NewContractRepository)
	return &ContractManagerService{}, nil
}

func InitializeTransactionServer(port int, opts []grpc.ServerOption, client pb.SigningServiceClient, contractConfig storage.ContractConfig, transactionConfig storage.TransactionConfig) (*TransactionRPCService, error) {
	wire.Build(NewTransactionServer, InitializeContractTransactionHandler)
	return &TransactionRPCService{}, nil
}

func InitializeContractServer(port int, opts []grpc.ServerOption, contractConfig storage.ContractConfig) (*ContractRPCService, error) {
	wire.Build(NewContractServer, InitializeContractManagerHandler)
	return &ContractRPCService{}, nil
}