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

func NewRedisWriter(endpoint, pwd, countKey string) RedisWriter {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: pwd,
		DB:       0,
	}),
	countKey: countKey,
	}
}

func NewRedisListener(endpoint, pwd string) RedisListener {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr: endpoint,
		Password: pwd,
		DB: 0,
	})}
}

func (r *Redis) VerifyValidAddress(ctx context.Context, address string) error {
	rdsRes, rdsErr := r.client.Get(ctx, address).Result()
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

func (r *Redis) GetReservedCount(ctx context.Context, avatarsRequested, maxMintable int) (error) {
	rdsCountResp, rdsCountErr := r.client.Get(ctx, r.countKey).Result()
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
		if count + avatarsRequested > maxMintable {
			return errors.New("the amount of tokens requested exceeds capacity")
		}
		return nil
	}
}

func (r *Redis) MarkAddressAsUsed(ctx context.Context, address, token string) error {
	return r.client.Set(ctx, address, token, 0).Err()
}

func (r *Redis) GetQueueNum(ctx context.Context) (int64, error) {
	return r.client.DBSize(ctx).Result()
}

func (r *Redis) IncrementCounter(ctx context.Context, avatarsRequested, maxMintable int) error {
	newMax, rdsCountIncrErr := r.client.IncrBy(ctx, r.countKey, int64(avatarsRequested)).Result()
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