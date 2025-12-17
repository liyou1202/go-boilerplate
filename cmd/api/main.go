package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"sync_drive_backend/configs"
	"sync_drive_backend/internal/infrastructure/webserver"
	"sync_drive_backend/pkg/logger"
)

func main() {
	// 載入配置
	if err := configs.Load(); err != nil {
		fmt.Printf("Failed to load config: %v\n", err)
		os.Exit(1)
	}

	// 初始化 Logger
	logLevel := configs.GetString("log.level")
	logFormat := configs.GetString("log.format")
	if err := logger.Init(logLevel, logFormat); err != nil {
		fmt.Printf("Failed to init logger: %v\n", err)
		os.Exit(1)
	}
	defer logger.Sync()

	// 設定路由
	router := webserver.SetupRouter()

	// 取得服務端口
	port := configs.GetInt("app.port")
	if port == 0 {
		port = 8080 // 預設端口
	}

	// 啟動服務器
	addr := fmt.Sprintf(":%d", port)
	logger.Info(fmt.Sprintf("Starting server on %s", addr))

	// 在 goroutine 中啟動服務器
	go func() {
		if err := router.Run(addr); err != nil {
			logger.Fatal(fmt.Sprintf("Failed to start server: %v", err))
		}
	}()

	// 等待中斷信號以優雅關閉
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("Shutting down server...")
}
