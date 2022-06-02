package contracts

import (
	"context"
	"contract-service/storage"
)

type ContractTransactionHandler interface {
	GetContract(ctx context.Context, address string) (*storage.Contract, error)
	BuildTransaction(ctx context.Context, msgSender, functionName string, numRequested int, arguments []interface{}, contract storage.Contract) (*storage.Token, error)
	StoreToken(ctx context.Context, token *storage.Token, contract *storage.Contract) error
	CheckIfValidRequest(ctx context.Context, msgSender string, numRequested int, contract *storage.Contract) error
}

type ContractManagerHandler interface {
	GetContract(ctx context.Context, address string) (*storage.Contract, error)
	StoreContract(ctx context.Context, contract *storage.Contract) error
	DeleteContract(ctx context.Context, address, owner string) error
	ListContracts(ctx context.Context, owner string) ([]*storage.Contract, error)
}
