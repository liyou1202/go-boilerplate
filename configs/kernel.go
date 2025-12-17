package configs

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

var Config *viper.Viper

// Load 載入配置文件
// 根據環境變數 ENV 決定載入哪個配置檔案 (loc, dev, prod)
func Load() error {
	Config = viper.New()

	// 從環境變數讀取環境名稱，預設為 loc (本地開發)
	env := os.Getenv("ENV")
	if env == "" {
		env = "loc"
	}

	// 設定配置檔案名稱和路徑
	Config.SetConfigName(env)
	Config.SetConfigType("toml")
	Config.AddConfigPath("./configs")
	Config.AddConfigPath("../configs")
	Config.AddConfigPath("../../configs")

	// 讀取配置檔案
	if err := Config.ReadInConfig(); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	// 允許環境變數覆蓋配置
	Config.AutomaticEnv()

	return nil
}

// GetString 取得字串配置值
func GetString(key string) string {
	return Config.GetString(key)
}

// GetInt 取得整數配置值
func GetInt(key string) int {
	return Config.GetInt(key)
}

// GetBool 取得布林配置值
func GetBool(key string) bool {
	return Config.GetBool(key)
}

// GetStringMap 取得字串 map 配置值
func GetStringMap(key string) map[string]interface{} {
	return Config.GetStringMap(key)
}
