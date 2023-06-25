package adapter

import (
	"time"

	"go.uber.org/zap"
)

type ZapImplement struct {
	logger *zap.Logger
}

func NewZapLogger(logger *zap.Logger) *ZapImplement {
	return &ZapImplement{logger}
}

func ProvideLogger(logger *zap.Logger) *ZapImplement {
	return NewZapLogger(logger)
}

func (z *ZapImplement) Debug(msg string) {
	z.logger.Debug(
		msg,
		zap.Time("time", time.Now()),
	)
}

func (z *ZapImplement) Error(msg string) {
	z.logger.Error(
		msg,
		zap.Time("time", time.Now()),
	)
}

func (z *ZapImplement) Fatal(msg string) {
	z.logger.Fatal(
		msg,
		zap.Time("time", time.Now()),
	)
}

func (z *ZapImplement) Info(msg string) {
	z.logger.Info(
		msg,
		zap.Time("time", time.Now()),
	)
}
