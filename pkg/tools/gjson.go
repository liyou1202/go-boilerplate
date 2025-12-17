package tools

import "github.com/tidwall/gjson"

// GetString 從 JSON 字串中取得字串值
func GetString(json, path string) string {
	return gjson.Get(json, path).String()
}

// GetInt 從 JSON 字串中取得整數值
func GetInt(json, path string) int64 {
	return gjson.Get(json, path).Int()
}

// GetFloat 從 JSON 字串中取得浮點數值
func GetFloat(json, path string) float64 {
	return gjson.Get(json, path).Float()
}

// GetBool 從 JSON 字串中取得布林值
func GetBool(json, path string) bool {
	return gjson.Get(json, path).Bool()
}

// Exists 檢查 JSON 路徑是否存在
func Exists(json, path string) bool {
	return gjson.Get(json, path).Exists()
}
