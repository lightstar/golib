// Package log provides simple logger designed to write info messages into stdout, and error messages into stderr.
// It uses zap sugared logger undercover but with very simple API that is enough in almost all situations.
//
// Typical usage:
//
//	logger := log.MustNew(log.WithName("my-logger"), log.WithDebug(true))
//	logger.Debug("Some debug message")
//	logger.Info("Some info message")
//	logger.Error("Some error message")
//	logger.Sync()
//
// Name is optional but highly recommended if you have more than one logger in your application to distinguish them.
package log

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

//nolint:gochecknoglobals // these are actually read-only, so it's ok to use them.
var (
	errorLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	infoLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.InfoLevel
	})
	debugLevelEnabler = zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.InfoLevel
	})

	stdout = zapcore.Lock(os.Stdout)
	stderr = zapcore.Lock(os.Stderr)
)

// Logger interface.
type Logger interface {
	Debug(msg string, v ...interface{})
	Info(msg string, v ...interface{})
	Error(msg string, v ...interface{})
	Fatal(msg string, v ...interface{})
	Sync()
}
