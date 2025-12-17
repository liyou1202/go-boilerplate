package tools

import "github.com/google/uuid"

// GenerateUUID 生成 UUID
func GenerateUUID() string {
	return uuid.New().String()
}

// IsValidUUID 驗證 UUID 格式是否正確
func IsValidUUID(str string) bool {
	_, err := uuid.Parse(str)
	return err == nil
}
