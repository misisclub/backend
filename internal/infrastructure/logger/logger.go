package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger() *Logger {
	config := zap.NewDevelopmentEncoderConfig()
	config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(config),
		zapcore.Lock(os.Stdout),
		zapcore.DebugLevel,
	)
	logger := zap.New(core).Sugar()
	return &Logger{logger}
}
