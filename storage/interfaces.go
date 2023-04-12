package storage

import "context"

type TransactionRepository interface {
	StoreTransaction(ctx context.Context, token Transaction) error
	GetTransactions(ctx context.Context, address string) ([]*Transaction, error)
	DeleteTransaction(ctx context.Context, address, hash string) error
	CompleteTransaction(ctx context.Context, address, hash string) error
	GetAllTransactions(ctx context.Context, address string) ([]*Transaction, error)
}

type PrivateKeyRepository interface {
	GetPrivateKey(ctx context.Context, contractAddress string) (string, error)
	UpsertPrivateKey(ctx context.Context, contractAddress, key string) error
	DeletePrivateKey(ctx context.Context, contractAddress string) error
}

type ContractRepository interface {
	GetContract(ctx context.Context, contractAddress string) (*Contract, error)
	UpsertContract(ctx context.Context, contract *Contract) error
	DeleteContract(ctx context.Context, contractAddress, owner string) error
	GetContractsByOwner(ctx context.Context, owner string) ([]*Contract, error)
}