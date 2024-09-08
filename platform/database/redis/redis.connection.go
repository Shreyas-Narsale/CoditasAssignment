package redis

import (
	"SOCIAL_MEDIA_APP/pkg/config"
	"SOCIAL_MEDIA_APP/pkg/logger"
	"context"
	"fmt"
	"sync"

	"github.com/redis/go-redis/v9"
)

var (
	rdb       *redis.Client
	rdbError  error
	redisOnce sync.Once
	ctx       = context.Background()
)

func initializeRedis() {
	DbConfig := config.GetDBConfig()
	logs := logger.GetLogger()

	ip := DbConfig.Redis.IP
	port := DbConfig.Redis.Port
	password := DbConfig.Redis.Password
	addr := fmt.Sprintf("%s:%d", ip, port)

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       0,
	})

	// Ping the Redis server to check the connection
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		rdbError = fmt.Errorf("error connecting to redis: %w", err)
		return
	}
	logs.Info().Msg("Redis connected Successfully!")
}

func GetRedisClient() (*redis.Client, error) {
	logs := logger.GetLogger()

	redisOnce.Do(initializeRedis)
	if rdbError != nil {
		logs.Error().Err(rdbError).Msg("redis error:")
		return nil, rdbError
	}

	return rdb, nil
}
