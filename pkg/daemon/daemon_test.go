package daemon_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/daemon"
	"github.com/lightstar/golib/pkg/log"
	"github.com/lightstar/golib/pkg/test/configtest"
)

func TestDaemon(t *testing.T) {
	testProcessor := &processor{}

	dmn, err := daemon.New(
		daemon.WithName("test-daemon"),
		daemon.WithDelay(2000),
		daemon.WithProcessor(testProcessor),
		daemon.WithLogger(log.NewNop()),
	)
	require.NoError(t, err)

	go func() {
		<-time.After(5 * time.Second)
		dmn.Terminate()
	}()

	dmn.Run()

	require.Equal(t, "test-daemon", dmn.Name())
	require.Equal(t, 2, testProcessor.processCalled)
}

func TestProcessFunc(t *testing.T) {
	processCalled := 0

	dmn, err := daemon.New(
		daemon.WithName("test-daemon"),
		daemon.WithDelay(2000),
		daemon.WithProcessFunc(func() {
			processCalled++
		}),
		daemon.WithLogger(log.NewNop()),
	)
	require.NoError(t, err)

	go func() {
		<-time.After(3 * time.Second)
		dmn.Terminate()
	}()

	dmn.Run()

	require.Equal(t, "test-daemon", dmn.Name())
	require.Equal(t, 1, processCalled)
}

func TestNopDaemon(t *testing.T) {
	dmn, err := daemon.New(daemon.WithLogger(log.NewNop()))
	require.NoError(t, err)

	var stopCalled bool

	go func() {
		<-time.After(1 * time.Second)

		stopCalled = true

		dmn.Terminate()
	}()

	dmn.Run()

	require.True(t, stopCalled)
}

func TestConfig(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Name  string
			Delay int
		}{
			Name:  "test-daemon",
			Delay: 2000,
		},
	})
	var dmn *daemon.Daemon

	require.NotPanics(t, func() {
		dmn = daemon.MustNew(daemon.WithConfig(configService, "key"))
	})

	require.Equal(t, "test-daemon", dmn.Name())
	require.Equal(t, 2*time.Second, dmn.Delay())
}

func TestConfigDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": config.ErrNoSuchKey,
	})
	var dmn *daemon.Daemon

	require.NotPanics(t, func() {
		dmn = daemon.MustNew(daemon.WithConfig(configService, "key"))
	})

	require.Equal(t, "daemon", dmn.Name())
	require.Equal(t, 1*time.Millisecond, dmn.Delay())
}

func TestConfigError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = daemon.MustNew(daemon.WithConfig(configService, "key"))
	})
}
