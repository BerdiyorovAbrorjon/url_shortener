package rediscache

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	redisCl    *redis.Client
	expiration time.Duration
}

func NewRedisCache(client *redis.Client, expiration time.Duration) *RedisCache {
	return &RedisCache{
		redisCl:    client,
		expiration: expiration,
	}
}
