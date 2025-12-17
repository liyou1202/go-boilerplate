package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

// Config Logger 配置
type Config struct {
	Level      string // 日誌級別：debug, info, warn, error
	Format     string // 日誌格式：json, console
	OutputPath string // 日誌檔案路徑（空字串則不輸出到檔案）
	MaxSize    int    // 單個日誌檔案最大大小（MB）
	MaxBackups int    // 保留的舊日誌檔案最大數量
	MaxAge     int    // 保留舊日誌檔案的最大天數
	Compress   bool   // 是否壓縮舊日誌檔案
}

// Init 初始化 Zap Logger
func Init(level, format string) error {
	cfg := &Config{
		Level:  level,
		Format: format,
	}
	return InitWithConfig(cfg)
}

// InitWithConfig 使用完整配置初始化 Logger
func InitWithConfig(cfg *Config) error {
	// 設定日誌級別
	var zapLevel zapcore.Level
	switch cfg.Level {
	case "debug":
		zapLevel = zapcore.DebugLevel
	case "info":
		zapLevel = zapcore.InfoLevel
	case "warn":
		zapLevel = zapcore.WarnLevel
	case "error":
		zapLevel = zapcore.ErrorLevel
	default:
		zapLevel = zapcore.InfoLevel
	}

	// 設定編碼器配置
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 設定編碼器（根據格式決定是否使用顏色）
	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		// Console 格式使用顏色輸出
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 設定輸出目標
	var writeSyncer zapcore.WriteSyncer

	// 如果設定了檔案輸出路徑，則同時輸出到 stdout 和檔案
	if cfg.OutputPath != "" {
		// 檔案輸出（使用 lumberjack 實現日誌輪轉）
		fileWriter := &lumberjack.Logger{
			Filename:   cfg.OutputPath,
			MaxSize:    cfg.MaxSize,    // MB
			MaxBackups: cfg.MaxBackups,
			MaxAge:     cfg.MaxAge,     // days
			Compress:   cfg.Compress,
		}

		// 同時輸出到 stdout 和檔案
		writeSyncer = zapcore.NewMultiWriteSyncer(
			zapcore.AddSync(os.Stdout),
			zapcore.AddSync(fileWriter),
		)
	} else {
		// 只輸出到 stdout
		writeSyncer = zapcore.AddSync(os.Stdout)
	}

	// 建立 core
	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		zapLevel,
	)

	// 建立 logger
	Log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))

	return nil
}

// Debug 記錄 debug 級別日誌
func Debug(msg string, fields ...zap.Field) {
	Log.Debug(msg, fields...)
}

// Info 記錄 info 級別日誌
func Info(msg string, fields ...zap.Field) {
	Log.Info(msg, fields...)
}

// Warn 記錄 warn 級別日誌
func Warn(msg string, fields ...zap.Field) {
	Log.Warn(msg, fields...)
}

// Error 記錄 error 級別日誌
func Error(msg string, fields ...zap.Field) {
	Log.Error(msg, fields...)
}

// Fatal 記錄 fatal 級別日誌並退出程式
func Fatal(msg string, fields ...zap.Field) {
	Log.Fatal(msg, fields...)
}

// Sync 同步日誌緩衝區
func Sync() error {
	return Log.Sync()
}
