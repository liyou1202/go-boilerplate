package request

import (
	"sync_drive_backend/pkg/tools"

	"github.com/gin-gonic/gin"
)

const (
	// HeaderRequestID Request ID header key
	HeaderRequestID = "X-Request-ID"
	// ContextKeyRequestID Request ID context key
	ContextKeyRequestID = "request_id"
)

// RequestID 請求 ID 追蹤中介層
// 為每個請求生成唯一的 Request ID，用於日誌追蹤和問題排查
func RequestID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 優先使用客戶端提供的 Request ID
		requestID := c.GetHeader(HeaderRequestID)

		// 如果客戶端沒有提供，則自動生成
		if requestID == "" {
			requestID = tools.GenerateUUID()
		}

		// 將 Request ID 存入 Context
		c.Set(ContextKeyRequestID, requestID)

		// 將 Request ID 寫入 Response Header
		c.Header(HeaderRequestID, requestID)

		c.Next()
	}
}

// GetRequestID 從 Context 取得 Request ID
func GetRequestID(c *gin.Context) string {
	if requestID, exists := c.Get(ContextKeyRequestID); exists {
		return requestID.(string)
	}
	return ""
}
