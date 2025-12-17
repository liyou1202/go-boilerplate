package mysql

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config MySQL 配置
type Config struct {
	Host         string
	Port         int
	Database     string
	Username     string
	Password     string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

// Init 初始化 MySQL 連接
func Init(cfg *Config) (*gorm.DB, error) {
	// 組建 DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%t&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
		cfg.Charset,
		cfg.ParseTime,
	)

	// 連接 MySQL
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to connect mysql: %w", err)
	}

	// 取得底層 *sql.DB 並設定連接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get sql.DB: %w", err)
	}

	// 設定連接池參數
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return db, nil
}
