// Package log provides simple logger designed to write info messages into stdout, and error messages into stderr.
// It uses zap sugared logger undercover but with very simple API that is enough in almost all situations.
//
// Typical usage:
//      logger := log.MustNew(log.WithName("my-logger"), log.WithDebug())
//      logger.Debug("Some debug message")
//      logger.Info("Some info message")
//      logger.Error("Some error message")
//      logger.Sync()
// Name is optional but highly recommended if you have more than one logger in your application to distinguish them.
package log

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	timeLayout = "[2006-01-02 15:04:05]"
)

// nolint: gochecknoglobals // these are actually read-only so it's ok to use them.
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

// standardLogger is an internal logger structure implementing Logger interface.
type standardLogger struct {
	zapLogger *zap.SugaredLogger
}

// New function creates new logger with provided options.
func New(opts ...Option) (Logger, error) {
	config, err := buildConfig(opts)
	if err != nil {
		return nil, err
	}

	loggerStdout, loggerStderr := stdout, stderr

	if config.stdout != nil {
		loggerStdout = zapcore.Lock(config.stdout)
	}

	if config.stderr != nil {
		loggerStderr = zapcore.Lock(config.stderr)
	}

	encoder := consoleEncoder()

	var stdoutLevelEnabler zapcore.LevelEnabler
	if config.debug {
		stdoutLevelEnabler = debugLevelEnabler
	} else {
		stdoutLevelEnabler = infoLevelEnabler
	}

	zapLogger := zap.New(
		zapcore.NewTee(
			zapcore.NewCore(encoder, loggerStdout, stdoutLevelEnabler),
			zapcore.NewCore(encoder, loggerStderr, errorLevelEnabler),
		),
		zap.WithCaller(false),
		zap.AddStacktrace(errorLevelEnabler),
		zap.AddCallerSkip(1),
		zap.OnFatal(zapcore.WriteThenPanic),
	)

	if config.name != "" {
		zapLogger = zapLogger.Named(fmt.Sprintf("(%s)", config.name))
	}

	return &standardLogger{
		zapLogger.Sugar(),
	}, nil
}

// MustNew function creates new logger with provided options and panics on any error.
func MustNew(opts ...Option) Logger {
	logger, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return logger
}

// Debug method writes debug message into log.
func (logger *standardLogger) Debug(msg string, v ...interface{}) {
	logger.zapLogger.Debugf(msg, v...)
}

// Info method writes info message into log.
func (logger *standardLogger) Info(msg string, v ...interface{}) {
	logger.zapLogger.Infof(msg, v...)
}

// Error method writes error message into log.
func (logger *standardLogger) Error(msg string, v ...interface{}) {
	logger.zapLogger.Errorf(msg, v...)
}

// Fatal method writes fatal message into log, then calls panic.
func (logger *standardLogger) Fatal(msg string, v ...interface{}) {
	logger.zapLogger.Fatalf(msg, v...)
}

// Sync method synchronizes logger io.
func (logger *standardLogger) Sync() {
	_ = logger.zapLogger.Sync()
}

// consoleEncoder function creates simple generic console encoder.
func consoleEncoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		TimeKey:          "time",
		NameKey:          "name",
		MessageKey:       "message",
		StacktraceKey:    "stacktrace",
		LineEnding:       zapcore.DefaultLineEnding,
		ConsoleSeparator: " ",
		EncodeTime:       zapcore.TimeEncoderOfLayout(timeLayout),
		EncodeName:       zapcore.FullNameEncoder,
	})
}
