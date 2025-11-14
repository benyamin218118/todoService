package repositories

import (
	"context"

	"github.com/benyamin218118/todoService/domain"
	"github.com/benyamin218118/todoService/domain/contracts"
	"github.com/go-redis/redis/v8"
)

type redisStreamPublisher struct {
	client *redis.Client
}

func NewRedisPubSubRepository(conf *domain.Config) contracts.IPubSub {
	opts, err := redis.ParseURL(conf.RedisUrl)
	if err != nil {
		panic(err)
	}
	redisClient := redis.NewClient(opts)
	return &redisStreamPublisher{
		client: redisClient,
	}
}

func (r *redisStreamPublisher) Publish(stream string, message map[string]any) error {
	return r.client.XAdd(context.Background(), &redis.XAddArgs{
		Stream: stream,
		Values: message,
	},
	).Err()
}
