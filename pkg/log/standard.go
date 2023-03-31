package log

import (
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	timeLayout = "[2006-01-02 15:04:05]"
)

// StandardLogger which uses standard logging format and implements Logger interface.
type StandardLogger struct {
	zapLogger *zap.SugaredLogger
}

// New function creates new standard logger with provided options.
func New(opts ...Option) (*StandardLogger, error) {
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
		zap.WithFatalHook(zapcore.WriteThenPanic),
	)

	if config.name != "" {
		zapLogger = zapLogger.Named(fmt.Sprintf("(%s)", config.name))
	}

	return &StandardLogger{
		zapLogger.Sugar(),
	}, nil
}

// MustNew function creates new standard logger with provided options and panics on any error.
func MustNew(opts ...Option) *StandardLogger {
	logger, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return logger
}

// Debug method writes debug message into log.
func (logger *StandardLogger) Debug(msg string, v ...interface{}) {
	logger.zapLogger.Debugf(msg, v...)
}

// Info method writes info message into log.
func (logger *StandardLogger) Info(msg string, v ...interface{}) {
	logger.zapLogger.Infof(msg, v...)
}

// Error method writes error message into log.
func (logger *StandardLogger) Error(msg string, v ...interface{}) {
	logger.zapLogger.Errorf(msg, v...)
}

// Fatal method writes fatal message into log, then calls panic.
func (logger *StandardLogger) Fatal(msg string, v ...interface{}) {
	logger.zapLogger.Fatalf(msg, v...)
}

// Sync method synchronizes logger io.
func (logger *StandardLogger) Sync() {
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
