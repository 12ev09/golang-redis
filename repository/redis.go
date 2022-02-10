package repository

import (
	"github.com/go-redis/redis/v8"
)

var Cache *redis.Client

// Redisに接続する
func SetupRedis() {
	Cache = redis.NewClient(
		&redis.Options{
			Addr: "redis:6379",
			DB:   0,
		},
	)
}
