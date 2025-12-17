package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Log *zap.Logger

// Init 初始化 Zap Logger
// level: debug, info, warn, error
// format: json, console
func Init(level, format string) error {
	// 設定日誌級別
	var zapLevel zapcore.Level
	switch level {
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
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// 設定編碼器
	var encoder zapcore.Encoder
	if format == "json" {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 設定輸出
	core := zapcore.NewCore(
		encoder,
		zapcore.AddSync(os.Stdout),
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
