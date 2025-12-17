package webserver

import (
	"sync_drive_backend/internal/common/middleware/logging"
	"sync_drive_backend/internal/common/middleware/request"
	"sync_drive_backend/internal/infrastructure/webserver/health"

	"github.com/gin-gonic/gin"
)

// SetupRouter 設定路由
func SetupRouter() *gin.Engine {
	// 創建 Gin Engine
	router := gin.New()

	// 全局中介層
	router.Use(gin.Recovery())           // 恢復 panic
	router.Use(request.RequestID())      // Request ID 追蹤
	router.Use(logging.Logger())         // 請求日誌記錄

	// 健康檢查端點（不需要認證）
	healthHandler := health.NewHandler()
	router.GET("/health", healthHandler.Check)

	// API 路由群組
	api := router.Group("/api/v1")
	{
		// TODO: 註冊業務路由
		// 例如：
		// auth := api.Group("/auth")
		// {
		//     auth.POST("/login", authController.Login)
		//     auth.POST("/register", authController.Register)
		// }

		_ = api // 避免未使用變數錯誤
	}

	return router
}
