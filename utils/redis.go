package utils

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
	"sync"
	"user_system/config"
)

var (
	redisConn *redis.Client
	redisOnce sync.Once
)

// openDB 连接db
func initRedis() {
	redisConfig := config.GetGlobalConf().RedisConfig
	addr := fmt.Sprintf("%s:%d", redisConfig.Host, redisConfig.Port)
	redisConn = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisConfig.PassWord,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	if redisConn == nil {
		panic("failed to call redis.NewClient")
	}

	_, err := redisConn.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to ping redis, err:%s")
	}
}

func CloseRedis() {
	redisConn.Close()
}

// GetDB 获取数据库连接
func GetRedisCli() *redis.Client {
	redisOnce.Do(initRedis)
	return redisConn
}
