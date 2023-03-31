package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewNop function creates dummy logger that produces no output at all.
func NewNop() *StandardLogger {
	return &StandardLogger{
		zap.NewNop().WithOptions(zap.WithFatalHook(zapcore.WriteThenPanic)).Sugar(),
	}
}
