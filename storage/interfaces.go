package storage

import "context"

type RedisListener interface {
	InitEvents() error
	Close()
	Listen(handler func(string, string, error) error) error
	Ping() (string, error)
}

type RedisWriter interface {
	VerifyValidAddress(ctx context.Context, address, contractAddress string) error
	GetReservedCount(ctx context.Context, numRequested, maxMintable int, contractAddress string) error
	MarkAddressAsUsed(ctx context.Context, token *Token) error
	GetQueueNum(ctx context.Context) (int64, error)
	IncrementCounter(ctx context.Context, numRequested, maxMintable int, contractAddress string) error
	Get(ctx context.Context, address, contractAddres string) (*Token, error)
	Ping() (string, error)
	Close()
}

type TransactionRepository interface {
	StoreTransaction(ctx context.Context, token Token) error
	GetTransactions(ctx context.Context, address string) ([]*Token, error)
	DeleteTransaction(ctx context.Context, address, hash string) error
	CompleteTransaction(ctx context.Context, address, hash string) error
	GetAllTransactions(ctx context.Context, address string) ([]*Token, error)
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