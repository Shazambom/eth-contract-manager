package storage

import "context"

type RedisListener interface {
	InitEvents() error
	Close()
	Listen(handler func(string, string, error) error) error
	Ping() (string, error)
}

type RedisWriter interface {
	VerifyValidAddress(ctx context.Context, address string) error
	GetReservedCount(ctx context.Context, avatarsRequested, maxMintable int) (error)
	MarkAddressAsUsed(ctx context.Context, token *Token) error
	GetQueueNum(ctx context.Context) (int64, error)
	IncrementCounter(ctx context.Context, avatarsRequested, maxMintable int) error
	Ping() (string, error)
	Close()
}

type PrivateKeyRepository interface {
	GetPrivateKey(ctx context.Context, contractAddress string) (string, error)
	UpsertPrivateKey(ctx context.Context, contractAddress, key string) error
	DeletePrivateKey(ctx context.Context, contractAddress string) error
}