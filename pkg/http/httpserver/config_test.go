package httpserver_test

import (
	"testing"

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
		)
	})

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "test-address", server.Address())
}

func TestConfigDefault(t *testing.T) {
	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew()
	})

	require.Equal(t, httpserver.DefName, server.Name())
	require.Equal(t, httpserver.DefAddress, server.Address())
}

func TestConfigService(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Name    string
			Address string
		}{
			Name:    "test-server",
			Address: "test-address",
		},
	})

	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "test-address", server.Address())
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
}

func TestConfigServiceError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})
}
