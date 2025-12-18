package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	// 初始化應用程式（使用 Wire 依賴注入）
	app, err := InitServer()
	if err != nil {
		fmt.Printf("Failed to initialize app: %v\n", err)
		os.Exit(1)
	}

	// 檢查外部服務連線狀態
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()

	// 檢查 MySQL
	sqlDB, err := app.MySQL.DB()
	if err != nil {
		app.Logger.Error("Failed to get underlying sql.DB from gorm", zap.Error(err))
	} else if err := sqlDB.PingContext(pingCtx); err != nil {
		app.Logger.Error("MySQL connection failed", zap.Error(err))
	} else {
		app.Logger.Info("MySQL connection successful")
	}

	// 檢查 MongoDB
	if err := app.MongoDB.Client().Ping(pingCtx, nil); err != nil {
		app.Logger.Error("MongoDB connection failed", zap.Error(err))
	} else {
		app.Logger.Info("MongoDB connection successful")
	}

	// 檢查 Redis
	if err := app.Redis.Ping(pingCtx).Err(); err != nil {
		app.Logger.Error("Redis connection failed", zap.Error(err))
	} else {
		app.Logger.Info("Redis connection successful")
	}

	// 記錄啟動資訊
	app.Logger.Info("Starting SyncDrive API Server",
		zap.String("env", app.Config.GetString("app.env")),
		zap.Int("port", app.Config.GetInt("app.port")),
	)

	// 創建 HTTP Server
	port := app.Config.GetInt("app.port")
	server := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        app.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// 在 goroutine 中啟動服務器
	go func() {
		app.Logger.Info("HTTP server started",
			zap.Int("port", port),
		)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	// 等待中斷信號以優雅地關閉伺服器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	app.Logger.Info("Shutting down server...")

	// 優雅關閉，最多等待 30 秒
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.Logger.Error("Server forced to shutdown", zap.Error(err))
	}

	// 同步日誌緩衝區
	_ = app.Logger.Sync()

	app.Logger.Info("Server exited")
}
