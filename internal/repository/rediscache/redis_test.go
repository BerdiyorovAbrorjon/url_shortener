package rediscache

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/BerdiyorovAbrorjon/url-shortener/config"
	"github.com/BerdiyorovAbrorjon/url-shortener/pkg/database/redis"
)

var testRedisCache *RedisCache

func TestMain(m *testing.M) {
	cfg, err := config.NewConfig("../../..")
	if err != nil {
		log.Fatal("test - RedisCache - config.NewConfig: %w", err)
	}

	// Initialize Redis
	rdb, err := redis.NewClient(cfg.RedisAddress, "", 0)
	if err != nil {
		log.Fatal(fmt.Errorf("test - RedisCache - redis.NewClient: %w", err))
	}

	testRedisCache = NewRedisCache(rdb, cfg.RedisExpiration)

	os.Exit(m.Run())
}
