package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
)

type Redis struct {
	client *redis.Client
	countKey string
}

type RedisConfig struct {
	Endpoint string
	Password string
	CountKey string
}

func NewRedisWriter(config RedisConfig) RedisWriter {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     config.Endpoint,
		Password: config.Password,
		DB:       0,
	}),
	countKey: config.CountKey,
	}
}

func NewRedisListener(config RedisConfig) RedisListener {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr: config.Endpoint,
		Password: config.Password,
		DB: 0,
	})}
}


func (r *Redis) VerifyValidAddress(ctx context.Context, address, contractAddress string) error {
	rdsRes, rdsErr := r.client.Get(ctx, r.getUserKey(contractAddress, address)).Result()
	if rdsErr != nil && rdsErr != redis.Nil {
		fmt.Println(rdsErr.Error())
		return rdsErr
	}
	if rdsRes != "" {
		rdsErrStr := "registration multiple times is forbidden"
		fmt.Println(rdsErrStr)
		return errors.New(rdsErrStr)
	}
	return nil
}

func (r *Redis) GetReservedCount(ctx context.Context, numRequested, maxMintable int, contractAddress string) error {
	rdsCountResp, rdsCountErr := r.client.Get(ctx, r.getCountKey(contractAddress)).Result()
	if rdsCountErr != nil && rdsCountErr != redis.Nil {
		return rdsCountErr
	}
	if rdsCountResp == "" {
		return nil
	} else {
		count, strconvErr := strconv.Atoi(rdsCountResp)
		if strconvErr != nil {
			return strconvErr
		}
		if count + numRequested > maxMintable {
			return errors.New("the amount of tokens requested exceeds capacity")
		}
		return nil
	}
}

func (r *Redis) MarkAddressAsUsed(ctx context.Context, token *Token) error {
	str, err := token.ToString()
	if err != nil {
		return err
	}
	return r.client.Set(ctx, r.getUserKey(token.ContractAddress, token.UserAddress), str, 0).Err()
}

func (r *Redis) GetQueueNum(ctx context.Context) (int64, error) {
	return r.client.DBSize(ctx).Result()
}

func (r *Redis) IncrementCounter(ctx context.Context, numRequested, maxMintable int, contractAddress string) error {
	newMax, rdsCountIncrErr := r.client.IncrBy(ctx, r.getCountKey(contractAddress), int64(numRequested)).Result()
	if rdsCountIncrErr != nil {
		return rdsCountIncrErr
	}
	if int(newMax) > maxMintable {
		return errors.New("NFTs are sold out, for now")
	}
	return nil
}

func (r *Redis) Ping() (string, error) {
	return r.client.Ping(context.Background()).Result()
}

func (r *Redis) Close() {
	_ = r.client.Close()
}

func (r *Redis) InitEvents() error {
	return r.client.Do(context.Background(), "CONFIG", "SET", "notify-keyspace-events", "KEA").Err()
}
func (r *Redis) Listen(handler func(string, string, error) error) error{
	ctx := context.Background()
	stream := r.client.PSubscribe(ctx, "__keyevent*__:set*")
	if pingErr := stream.Ping(ctx, "test ping"); pingErr != nil {
		return pingErr
	}
	fmt.Printf("Stream: %+v\n", stream)
	for {
		msg, err := stream.ReceiveMessage(ctx)
		val, getErr := r.client.Get(ctx, msg.Payload).Result()
		if err == nil {
			err = getErr
		}
		if handleErr := handler(msg.Payload, val, err); handleErr != nil {
			return handleErr
		}

	}
}

func (r *Redis) getCountKey(contractAddress string) string {
	return r.countKey + "_" + contractAddress
}

func (r *Redis) getUserKey(contractAddress, address string) string {
	return contractAddress + "_" + address
}