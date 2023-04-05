package utils

import (
	"sync"
	"user_system/config"

	"github.com/go-redis/redis"
)

var (
	redisConn *redis.Client
	redisOnce sync.Once
)

// openDB 连接db
func initRedis() {
	redisConfig := config.GetGlobalConf().RedisConfig
	redisConn = redis.NewClient(&redis.Options{
		Addr:     redisConfig.Addr,
		Password: redisConfig.PassWord,
		DB:       redisConfig.DB,
		PoolSize: redisConfig.PoolSize,
	})
	if redisConn == nil {
		panic("failed to call redis.NewClient")
	}

	_, err := redisConn.Ping().Result()
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
