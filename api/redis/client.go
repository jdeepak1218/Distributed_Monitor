package redis

import (
	"context"
	"fmt"
	"os"

	goredis "github.com/go-redis/redis/v8"
)

var client *goredis.Client

func InitClient() error {
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "localhost:6379"
	}
	if Client == nil {
		Client = goredis.NewClient(
			&goredis.Options{
				Addr:     redisAddr,
				Password: os.Getenv("REDIS_PASSWORD"),
				DB:       0,
			},
		)
	}
	err := Client.Ping(context.Background()).Err()
	if err != nil {
		return fmt.Errorf("redis connection failed: %w", err)
	}
	fmt.Println("Redis connection established")
	return nil
}
func GetClient() *goredis.Client {
	return client
}
