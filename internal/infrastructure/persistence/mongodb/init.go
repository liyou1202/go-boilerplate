package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Config MongoDB 配置
type Config struct {
	URI      string
	Database string
	Timeout  int // 連接超時（秒）
}

// Init 初始化 MongoDB 連接
func Init(cfg *Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Timeout)*time.Second)
	defer cancel()

	// 創建客戶端選項
	clientOptions := options.Client().ApplyURI(cfg.URI)

	// 連接 MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect mongodb: %w", err)
	}

	// Ping 測試連接
	if err := client.Ping(ctx, nil); err != nil {
		return nil, fmt.Errorf("failed to ping mongodb: %w", err)
	}

	// 返回資料庫實例
	return client.Database(cfg.Database), nil
}
