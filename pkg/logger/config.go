package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
)

type ILog interface {
}

type CtxLogger struct {
	writer io.Writer
	Level  zapcore.Level
}

var config = zap.NewProductionEncoderConfig()

func (c *CtxLogger) SetWriter(w io.Writer) {
	c.writer = w
}

func (c *CtxLogger) SetDefaultWriter(filename string) {
	c.writer = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1024 * 1024 * 5,
		MaxBackups: 7,
		MaxAge:     14,
		Compress:   true,
	}
}

func NewLogger() *zap.Logger {
	config.EncodeTime = zapcore.RFC3339TimeEncoder
	config.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewJSONEncoder(config)

	lj := &lumberjack.Logger{
		Filename:   "./logs/test.log",
		MaxSize:    1,
		MaxBackups: 1,
		MaxAge:     2,
		Compress:   true,
	}
	sy := zapcore.AddSync(lj)
	core := zapcore.NewCore(encoder, sy, zapcore.DebugLevel)
	return zap.New(core, zap.AddCaller())
}
