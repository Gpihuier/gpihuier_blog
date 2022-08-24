package initialize

import (
	"context"
	"fmt"

	"github.com/Gpihuier/gpihuier_blog/global"

	"github.com/go-redis/redis/v8"
)

func RedisDrive() *redis.Client {
	config := global.CONFIG.Redis

	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Errorf("cache drive init is panic: %s", err))
	}
	return client
}
