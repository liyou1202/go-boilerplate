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
		Code:    ErrInternalError,
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

//	 Mapping rules:
//		code 0: Success → 200 OK
//		code 1-999: Application errors → 400 Bad Request (client errors)
//		code >1000: Infrastructure errors → 500 Internal Server Error
//		1000-1099: Database errors (MySQL, MongoDB, Redis)
//		1100-1999: Reserved for other infrastructure
//		2000-2099: AWS services (S3, etc.)
//		2100-2199: External APIs
// getHTTPStatus 粗略劃分若需要更細節則需要維護 error code 與 http code 的 mapping
func getHTTPStatus(code int) int {
	// Success
	if code == CodeSuccess {
		return http.StatusOK
	}

	// Application errors (1-999)
	if code >= 1 && code <= 999 {
		// Default: client error
		return http.StatusBadRequest
	}

	// Unknown error
	return http.StatusInternalServerError // 500
}
