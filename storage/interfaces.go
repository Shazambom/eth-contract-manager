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
	MarkAddressAsUsed(ctx context.Context, address, token string) error
	GetQueueNum(ctx context.Context) (int64, error)
	IncrementCounter(ctx context.Context, avatarsRequested, maxMintable int) error
	Ping() (string, error)
	Close()
}