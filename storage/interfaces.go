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
	Ping() (string, error)
	Close()
}

type PrivateKeyRepository interface {
	GetPrivateKey(ctx context.Context, contractAddress string) (string, error)
	UpsertPrivateKey(ctx context.Context, contractAddress, key string) error
	DeletePrivateKey(ctx context.Context, contractAddress string) error
}

type ContractRepository interface {
	GetContract(ctx context.Context, contractAddress string) (*Contract, error)
	UpsertContract(ctx context.Context, contract *Contract) error
	DeleteContract(ctx context.Context, contractAddress string) error
	GetContractsByOwner(ctx context.Context, owner string) ([]*Contract, error)
}