package redis

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

// Cache Redis 快取操作封裝
type Cache struct {
	client *redis.Client
}

// NewCache 創建快取實例
func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}

// Set 設定快取
func (c *Cache) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	return c.client.Set(ctx, key, value, expiration).Err()
}

// Get 取得快取
func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.client.Get(ctx, key).Result()
}

// GetBytes 取得快取（bytes）
func (c *Cache) GetBytes(ctx context.Context, key string) ([]byte, error) {
	return c.client.Get(ctx, key).Bytes()
}

// Del 刪除快取
func (c *Cache) Del(ctx context.Context, keys ...string) error {
	return c.client.Del(ctx, keys...).Err()
}

// Exists 檢查 key 是否存在
func (c *Cache) Exists(ctx context.Context, keys ...string) (int64, error) {
	return c.client.Exists(ctx, keys...).Result()
}

// Expire 設定過期時間
func (c *Cache) Expire(ctx context.Context, key string, expiration time.Duration) error {
	return c.client.Expire(ctx, key, expiration).Err()
}

// TTL 取得剩餘過期時間
func (c *Cache) TTL(ctx context.Context, key string) (time.Duration, error) {
	return c.client.TTL(ctx, key).Result()
}

// Incr 自增
func (c *Cache) Incr(ctx context.Context, key string) (int64, error) {
	return c.client.Incr(ctx, key).Result()
}

// Decr 自減
func (c *Cache) Decr(ctx context.Context, key string) (int64, error) {
	return c.client.Decr(ctx, key).Result()
}
