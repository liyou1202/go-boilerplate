package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 統一錯誤回應格式
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// HandleError 處理錯誤並回應
func HandleError(c *gin.Context, err error) {
	if appErr, ok := err.(*AppError); ok {
		// 自定義錯誤
		c.JSON(getHTTPStatus(appErr.Code), Response{
			Code:    appErr.Code,
			Message: appErr.Message,
		})
		return
	}

	// 一般錯誤
	c.JSON(http.StatusInternalServerError, Response{
		Code:    CodeInternalError,
		Message: err.Error(),
	})
}

// Success 回應成功
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    CodeSuccess,
		Message: "success",
		Data:    data,
	})
}

// getHTTPStatus 根據錯誤碼取得 HTTP 狀態碼
func getHTTPStatus(code int) int {
	switch {
	case code >= 20001 && code <= 20099:
		return http.StatusUnauthorized
	case code == CodeForbidden:
		return http.StatusForbidden
	case code == CodeNotFound:
		return http.StatusNotFound
	case code == CodeInvalidParams:
		return http.StatusBadRequest
	case code == CodeAlreadyExists:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
