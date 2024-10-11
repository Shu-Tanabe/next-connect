package utils

import (
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LogConfig struct {
	LogLevel string
}

var Logger *zap.Logger

func init() {
	Logger = NewLogger(LogConfig{
		LogLevel: "info",
	})
}

func NewLogger(cfg LogConfig) *zap.Logger {
	logLevel := getLogLevel(cfg)
	encoder := getEncoder()
	w := zapcore.Lock(os.Stdout)

	core := zapcore.NewCore(encoder, w, logLevel)
	logger := zap.New(core, zap.AddCaller())
	return logger
}

func getLogLevel(cfg LogConfig) zapcore.Level {
	switch {
		case strings.ToLower(cfg.LogLevel) == "debug":
			return zapcore.DebugLevel
		case strings.ToLower(cfg.LogLevel) == "info":
			return zapcore.InfoLevel
		case strings.ToLower(cfg.LogLevel) == "warn":
			return zapcore.WarnLevel
		case strings.ToLower(cfg.LogLevel) == "error":
			return zapcore.ErrorLevel
		default:
			return zapcore.InfoLevel
		}
	}

func getEncoder() zapcore.Encoder {
	encCfg := zapcore.EncoderConfig{
		LevelKey:   "level",
		TimeKey:    "timestamp",
		MessageKey: "message",
		FunctionKey: "function",
		CallerKey:  "caller",
		LineEnding: "\n",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	}
	return zapcore.NewJSONEncoder(encCfg)
}