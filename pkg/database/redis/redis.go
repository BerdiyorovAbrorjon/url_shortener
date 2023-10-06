package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

const timeout = 5 * time.Second

func NewClient(redisAddr string, redisPass string, redisDB int32) (*redis.Client, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPass,
		DB:       int(redisDB),
	})

	err := client.Ping(ctx).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to ping to redis: %w", err)
	}

	return client, nil
}
