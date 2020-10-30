package log

import "go.uber.org/zap"

// NewNop function creates dummy logger that produces no output at all.
func NewNop() Logger {
	return &standardLogger{
		zap.NewNop().Sugar(),
	}
}
