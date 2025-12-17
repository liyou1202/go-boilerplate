package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler 健康檢查處理器
type Handler struct{}

// NewHandler 創建健康檢查處理器
func NewHandler() *Handler {
	return &Handler{}
}

// Check 健康檢查端點
// GET /health
func (h *Handler) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"message": "service is healthy",
	})
}
