package main

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

func NewRedis(endpoint, pwd, countKey string) *Redis {
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: pwd,
		DB:       0,
	}),
	countKey: countKey,
	}
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
	return rds.client.Set(ctx, address, token, 0).Err()
}

func (r *Redis) GetQueueNum(ctx context.Context) (int64, error) {
	return rds.client.DBSize(ctx).Result()
}

func (r *Redis) IncrementCounter(ctx context.Context, avatarsRequested, maxMintable int) error {
	newMax, rdsCountIncrErr := rds.client.IncrBy(ctx, r.countKey, int64(avatarsRequested)).Result()
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