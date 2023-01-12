package contracts

import (
	"context"
	"contract-service/storage"
)

type ContractTransactionHandler interface {
	GetContract(ctx context.Context, address string) (*storage.Contract, error)
	BuildTransaction(ctx context.Context, senderInHash bool, msgSender, functionName string, arguments [][]byte, value string, contract *storage.Contract) (*storage.Token, error)
	StoreToken(ctx context.Context, token *storage.Token) error
	GetTransactions(ctx context.Context, address string) ([]*storage.Token, error)
	DeleteTransaction(ctx context.Context, address, hash string) error
	GetAllTransactions(ctx context.Context, address string) ([]*storage.Token, error)
	CompleteTransaction(ctx context.Context, address, hash string) error
}

type ContractManagerHandler interface {
	GetContract(ctx context.Context, address string) (*storage.Contract, error)
	StoreContract(ctx context.Context, contract *storage.Contract) error
	DeleteContract(ctx context.Context, address, owner string) error
	ListContracts(ctx context.Context, owner string) ([]*storage.Contract, error)
}
