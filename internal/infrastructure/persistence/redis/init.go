package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// Config Redis 配置
type Config struct {
	Host     string
	Port     int
	Password string
	DB       int
	PoolSize int
}

// Init 初始化 Redis 連接
func Init(cfg *Config) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
		Password: cfg.Password,
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	// Ping 測試連接
	ctx := context.Background()
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect redis: %w", err)
	}

	return client, nil
}
