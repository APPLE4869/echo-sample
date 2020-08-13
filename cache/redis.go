package cache

import (
	"context"

	"github.com/go-redis/redis/v8" // @link https://redis.uptrace.dev/
)

func client() *redis.Client {
	opt, err := redis.ParseURL("redis://redis:6379/0")
	if err != nil {
		panic(err)
	}
	return redis.NewClient(opt)
}

var ctx = context.Background()

// Get .
func Get(key string) string {
	val, err := client().Get(ctx, key).Result()
	if err != nil {
		panic("failed to get result from redis.")
	}
	return val
}

// Set .
func Set(key string, val string) error {
	err := client().Set(ctx, key, val, 0).Err()
	return err
}
