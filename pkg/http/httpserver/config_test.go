package httpserver_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/http/httpserver"
	"github.com/lightstar/golib/pkg/test/configtest"
)

func TestConfig(t *testing.T) {
	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew(
			httpserver.WithName("test-server"),
			httpserver.WithAddress("test-address"),
			httpserver.WithReadHeaderTimeout(10),
			httpserver.WithReadTimeout(4),
			httpserver.WithWriteTimeout(5),
		)
	})

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "test-address", server.Address())
	require.Equal(t, 10*time.Second, server.ReadHeaderTimeout())
	require.Equal(t, 4*time.Second, server.ReadTimeout())
	require.Equal(t, 5*time.Second, server.WriteTimeout())
}

func TestConfigDefault(t *testing.T) {
	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew()
	})

	require.Equal(t, httpserver.DefName, server.Name())
	require.Equal(t, httpserver.DefAddress, server.Address())
	require.Equal(t, httpserver.DefReadHeaderTimeout*time.Second, server.ReadHeaderTimeout())
	require.Equal(t, httpserver.DefReadTimeout*time.Second, server.ReadTimeout())
	require.Equal(t, httpserver.DefWriteTimeout*time.Second, server.WriteTimeout())
}

func TestConfigService(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Name              string
			Address           string
			ReadHeaderTimeout int64
			ReadTimeout       int64
			WriteTimeout      int64
		}{
			Name:              "test-server",
			Address:           "test-address",
			ReadHeaderTimeout: 10,
			ReadTimeout:       4,
			WriteTimeout:      5,
		},
	})

	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "test-address", server.Address())
	require.Equal(t, 10*time.Second, server.ReadHeaderTimeout())
	require.Equal(t, 4*time.Second, server.ReadTimeout())
	require.Equal(t, 5*time.Second, server.WriteTimeout())
}

func TestConfigServiceDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": config.ErrNoSuchKey,
	})

	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})

	require.Equal(t, httpserver.DefName, server.Name())
	require.Equal(t, httpserver.DefAddress, server.Address())
	require.Equal(t, httpserver.DefReadHeaderTimeout*time.Second, server.ReadHeaderTimeout())
	require.Equal(t, httpserver.DefReadTimeout*time.Second, server.ReadTimeout())
	require.Equal(t, httpserver.DefWriteTimeout*time.Second, server.WriteTimeout())
}

func TestConfigServiceError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})
}
