package logging

import (
	"time"

	"sync_drive_backend/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Logger 請求日誌記錄中介層
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 記錄開始時間
		start := time.Now()

		// 取得請求路徑
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		// 處理請求
		c.Next()

		// 計算請求時長
		latency := time.Since(start)

		// 記錄日誌
		logger.Info("HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
			zap.String("client_ip", c.ClientIP()),
			zap.String("user_agent", c.Request.UserAgent()),
		)
	}
}
