package daemon_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/daemon"
	"github.com/lightstar/golib/pkg/test/configtest"
)

func TestConfig(t *testing.T) {
	var dmn *daemon.Daemon

	require.NotPanics(t, func() {
		dmn = daemon.MustNew(
			daemon.WithName("test-daemon"),
			daemon.WithDelay(2000),
		)
	})

	require.Equal(t, "test-daemon", dmn.Name())
	require.Equal(t, 2*time.Second, dmn.Delay())
}

func TestConfigDefault(t *testing.T) {
	var dmn *daemon.Daemon

	require.NotPanics(t, func() {
		dmn = daemon.MustNew()
	})

	require.Equal(t, "daemon", dmn.Name())
	require.Equal(t, 1*time.Millisecond, dmn.Delay())
}

func TestConfigService(t *testing.T) {
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

func TestConfigServiceDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": configtest.ErrNoSuchKey,
	})
	var dmn *daemon.Daemon

	require.NotPanics(t, func() {
		dmn = daemon.MustNew(daemon.WithConfig(configService, "key"))
	})

	require.Equal(t, "daemon", dmn.Name())
	require.Equal(t, 1*time.Millisecond, dmn.Delay())
}

func TestConfigServiceError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = daemon.MustNew(daemon.WithConfig(configService, "key"))
	})
}
