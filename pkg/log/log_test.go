package log_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/log"
	"github.com/lightstar/golib/pkg/test/configtest"
	"github.com/lightstar/golib/pkg/test/iotest"
)

func TestLog(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()

	logger, err := log.New(
		log.WithName("test"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)
	require.NoError(t, err)

	defer logger.Sync()

	logger.Debug("Test debug message")
	logger.Info("Test info message")
	logger.Error("Test error message")

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test info message\n$`,
		stdout.String())
	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test error message\n`+
		`github\.com/lightstar/golib/pkg/log_test\.TestLog\n`, stderr.String())
}

func TestDebug(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()

	logger, err := log.New(
		log.WithName("test"),
		log.WithDebug(true),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)
	require.NoError(t, err)

	defer logger.Sync()

	logger.Debug("Test debug message")
	logger.Info("Test info message")
	logger.Error("Test error message")

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test debug message\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test info message\n$`, stdout.String())
	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test error message\n`+
		`github\.com/lightstar/golib/pkg/log_test\.TestDebug\n`, stderr.String())
}

func TestFatal(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()

	logger, err := log.New(
		log.WithName("test"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)
	require.NoError(t, err)

	defer logger.Sync()

	require.Panics(t, func() {
		logger.Fatal("Test fatal message")
	})

	require.Empty(t, stdout.String())
	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test fatal message\n`+
		`github\.com/lightstar/golib/pkg/log_test\.TestFatal\..+\n`, stderr.String())
}

func TestNop(t *testing.T) {
	logger := log.NewNop()
	defer logger.Sync()

	logger.Debug("Test debug message")
	logger.Info("Test info message")
	logger.Error("Test error message")

	require.Panics(t, func() {
		logger.Fatal("Test fatal message")
	})
}

func TestConfig(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()

	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Name  string
			Debug bool
		}{
			Name:  "test",
			Debug: true,
		},
	})
	var logger log.Logger

	require.NotPanics(t, func() {
		logger = log.MustNew(
			log.WithConfig(configService, "key"),
			log.WithStdout(stdout),
			log.WithStderr(stderr),
		)
	})

	defer logger.Sync()

	logger.Debug("Test debug message")
	logger.Info("Test info message")
	logger.Error("Test error message")

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test debug message\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test info message\n$`, stdout.String())
	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] \(test\) Test error message\n`+
		`github\.com/lightstar/golib/pkg/log_test\.TestConfig\n`, stderr.String())
}

func TestConfigDefault(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()

	configService := configtest.New(map[string]interface{}{
		"key": configtest.ErrNoSuchKey,
	})
	var logger log.Logger

	require.NotPanics(t, func() {
		logger = log.MustNew(
			log.WithConfig(configService, "key"),
			log.WithStdout(stdout),
			log.WithStderr(stderr),
		)
	})

	defer logger.Sync()

	logger.Debug("Test debug message")
	logger.Info("Test info message")
	logger.Error("Test error message")

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] Test info message\n$`, stdout.String())
	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}] Test error message\n`+
		`github\.com/lightstar/golib/pkg/log_test\.TestConfigDefault\n`, stderr.String())
}

func TestConfigError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = log.MustNew(log.WithConfig(configService, "key"))
	})
}
