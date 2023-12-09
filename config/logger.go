package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger
func init(){
	lumberjackLogger := &lumberjack.Logger{
		Filename:   "app.log", // 设置日志文件名
		MaxSize:    10,        // 设置日志文件的最大大小，单位是MB
		MaxBackups: 5,         // 设置最多保留的旧日志文件数
		MaxAge:     30,        // 设置保留的最大天数
		Compress:   true,      // 是否启用压缩
	}

	// 创建一个Encoder配置
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:   "message",
		LevelKey:     "level",
		TimeKey:      "time",
		CallerKey:    "caller",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeLevel:  zapcore.LowercaseLevelEncoder,
		EncodeTime:   zapcore.ISO8601TimeEncoder,
	}
	encoder := zapcore.NewJSONEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.AddSync(lumberjackLogger), zapcore.DebugLevel)
	// 创建Logger
	Logger = zap.New(core)
}