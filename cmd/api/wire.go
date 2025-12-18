//go:build wireinject
// +build wireinject

package main

import (
	"sync_drive_backend/configs"
	"sync_drive_backend/internal/infrastructure/persistence/mongodb"
	"sync_drive_backend/internal/infrastructure/persistence/mysql"
	redisinfra "sync_drive_backend/internal/infrastructure/persistence/redis"
	"sync_drive_backend/internal/infrastructure/webserver"
	"sync_drive_backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	redisclient "github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ProvideConfig 提供配置
func ProvideConfig() (*viper.Viper, error) {
	if err := configs.Load(); err != nil {
		return nil, err
	}
	return configs.Config, nil
}

// ProvideLogger 提供 Logger
func ProvideLogger(cfg *viper.Viper) (*zap.Logger, error) {
	loggerCfg := &logger.Config{
		Level:      cfg.GetString("log.level"),
		Format:     cfg.GetString("log.format"),
		OutputPath: cfg.GetString("log.outputPath"),
		MaxSize:    cfg.GetInt("log.maxSize"),
		MaxBackups: cfg.GetInt("log.maxBackups"),
		MaxAge:     cfg.GetInt("log.maxAge"),
		Compress:   cfg.GetBool("log.compress"),
	}

	if err := logger.InitWithConfig(loggerCfg); err != nil {
		return nil, err
	}

	return logger.Log, nil
}

// ProvideMySQL 提供 MySQL 連接
func ProvideMySQL(cfg *viper.Viper) (*gorm.DB, error) {
	mysqlCfg := &mysql.Config{
		Host:         cfg.GetString("mysql.host"),
		Port:         cfg.GetInt("mysql.port"),
		Database:     cfg.GetString("mysql.database"),
		Username:     cfg.GetString("mysql.username"),
		Password:     cfg.GetString("mysql.password"),
		Charset:      cfg.GetString("mysql.charset"),
		ParseTime:    cfg.GetBool("mysql.parseTime"),
		MaxIdleConns: cfg.GetInt("mysql.maxIdleConns"),
		MaxOpenConns: cfg.GetInt("mysql.maxOpenConns"),
	}

	return mysql.Init(mysqlCfg)
}

// ProvideMongoDB 提供 MongoDB 連接
func ProvideMongoDB(cfg *viper.Viper) (*mongo.Database, error) {
	mongoCfg := &mongodb.Config{
		URI:      cfg.GetString("mongodb.uri"),
		Database: cfg.GetString("mongodb.database"),
		Timeout:  cfg.GetInt("mongodb.timeout"),
	}

	return mongodb.Init(mongoCfg)
}

// ProvideRedis 提供 Redis 連接
func ProvideRedis(cfg *viper.Viper) (*redisclient.Client, error) {
	redisCfg := &redisinfra.Config{
		Host:     cfg.GetString("redis.host"),
		Port:     cfg.GetInt("redis.port"),
		Password: cfg.GetString("redis.password"),
		DB:       cfg.GetInt("redis.db"),
		PoolSize: cfg.GetInt("redis.poolSize"),
	}

	return redisinfra.Init(redisCfg)
}

// ProvideRouter 提供 Gin Router
func ProvideRouter() *gin.Engine {
	return webserver.SetupRouter()
}

// App 應用程式結構
type App struct {
	Config  *viper.Viper
	Logger  *zap.Logger
	MySQL   *gorm.DB
	MongoDB *mongo.Database
	Redis   *redisclient.Client
	Router  *gin.Engine
}

// newApp 創建 App 實例
func newApp(
	config *viper.Viper,
	logger *zap.Logger,
	mysql *gorm.DB,
	mongodb *mongo.Database,
	redis *redisclient.Client,
	router *gin.Engine,
) *App {
	return &App{
		Config:  config,
		Logger:  logger,
		MySQL:   mysql,
		MongoDB: mongodb,
		Redis:   redis,
		Router:  router,
	}
}

// InitServer 初始化應用程式（Wire 會生成此函數的實作）
func InitServer() (*App, error) {
	panic(wire.Build(
		// Config
		ProvideConfig,

		// Logger
		ProvideLogger,

		// Database
		ProvideMySQL,
		ProvideMongoDB,
		ProvideRedis,

		// Router
		ProvideRouter,

		// App
		newApp,
	))
}
