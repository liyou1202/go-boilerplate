package util

import (
	"strings"
)

// IsEmpty 檢查字串是否為空（包含只有空白字元的情況）
func IsEmpty(s string) bool {
	return strings.TrimSpace(s) == ""
}

// IsNotEmpty 檢查字串是否不為空
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// DefaultIfEmpty 如果字串為空則返回預設值
func DefaultIfEmpty(s, defaultValue string) string {
	if IsEmpty(s) {
		return defaultValue
	}
	return s
}

// Truncate 截斷字串到指定長度
func Truncate(s string, maxLength int) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength]
}

// TruncateWithSuffix 截斷字串並添加後綴（例如：...）
func TruncateWithSuffix(s string, maxLength int, suffix string) string {
	if len(s) <= maxLength {
		return s
	}
	return s[:maxLength-len(suffix)] + suffix
}
