package storage
import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	client *redis.Client
}

func NewRedis(endpoint, pwd string) *Redis {
	fmt.Printf("Creating client at endpoint: %s\n", endpoint)
	return &Redis{client: redis.NewClient(&redis.Options{
		Addr:     endpoint,
		Password: pwd,
		DB:       0,
	}),
	}
}

func (r *Redis) InitEvents() error {
	return r.client.Do(context.Background(), "CONFIG", "SET", "notify-keyspace-events", "KEA").Err()
}

func (r *Redis) Close() {
	_ = r.client.Close()
}

func (r *Redis) Listen(handler func(string, string, error) error) error{
	ctx := context.Background()
	//stream := r.client.PSubscribe(ctx, "__key*__:*","__keyevent@0__:set")
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
