package contracts

import (
	"context"
	"contract-service/storage"
)

type ContractTransactionHandler interface {
	GetContract(ctx context.Context, address string) (*storage.Contract, error)
	BuildTransaction(ctx context.Context, msgSender, functionName string, arguments [][]byte, contract *storage.Contract) (*storage.Token, error)
	StoreToken(ctx context.Context, token *storage.Token, contract *storage.Contract) error
	GetTransactions(ctx context.Context, address string) ([]*storage.Token, error)
	DeleteTransaction(ctx context.Context, address, hash string) error
	Close()
}

type ContractManagerHandler interface {
	GetContract(ctx context.Context, address string) (*storage.Contract, error)
	StoreContract(ctx context.Context, contract *storage.Contract) error
	DeleteContract(ctx context.Context, address, owner string) error
	ListContracts(ctx context.Context, owner string) ([]*storage.Contract, error)
}
