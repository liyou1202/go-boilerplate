package errors

import "fmt"

// AppError 自定義應用程式錯誤類型
type AppError struct {
	Code    int    // 錯誤碼
	Message string // 錯誤訊息
	Err     error  // 原始錯誤
}

// Error 實作 error 介面
func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

// New 建立新的 AppError
func New(code int, message string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
	}
}

// Wrap 包裝原始錯誤
func Wrap(code int, message string, err error) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}
