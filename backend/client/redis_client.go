package client

import (
	"backend/config"
	"context"
	"sync"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

// 单例模式的 Redis 客户端
var (
	redisClient *redis.Client
	redisOnce   sync.Once
)

// 获取 Redis 客户端
func GetRedisClient() *redis.Client {
	redisOnce.Do(func() {
		logrus.Infof("Connecting to Redis %s", config.GetString("redis.host"))
		redisClient = redis.NewClient(&redis.Options{
			Addr:     config.GetString("redis.host"),
			Password: "",
			DB:       0,
		})

		ctx := context.Background()
		if err := redisClient.Ping(ctx).Err(); err != nil {
			logrus.Fatalf("Failed to connect to Redis: %v", err)
		}
	})
	return redisClient
}
