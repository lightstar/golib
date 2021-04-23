package grpcserver_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/grpc/grpcserver"
	"github.com/lightstar/golib/pkg/test/configtest"
)

func TestConfig(t *testing.T) {
	var server *grpcserver.Server

	require.NotPanics(t, func() {
		server = grpcserver.MustNew(
			grpcserver.WithName("test-server"),
			grpcserver.WithAddress("test-address"),
			grpcserver.WithRegisterFn(func(*grpc.Server) {}),
		)
	})

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "test-address", server.Address())
}

func TestConfigDefault(t *testing.T) {
	var server *grpcserver.Server

	require.NotPanics(t, func() {
		server = grpcserver.MustNew(
			grpcserver.WithRegisterFn(func(*grpc.Server) {}),
		)
	})

	require.Equal(t, grpcserver.DefName, server.Name())
	require.Equal(t, grpcserver.DefAddress, server.Address())
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

	var server *grpcserver.Server

	require.NotPanics(t, func() {
		server = grpcserver.MustNew(
			grpcserver.WithConfig(configService, "key"),
			grpcserver.WithRegisterFn(func(*grpc.Server) {}),
		)
	})

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "test-address", server.Address())
}

func TestConfigServiceDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": config.ErrNoSuchKey,
	})

	var server *grpcserver.Server

	require.NotPanics(t, func() {
		server = grpcserver.MustNew(
			grpcserver.WithConfig(configService, "key"),
			grpcserver.WithRegisterFn(func(*grpc.Server) {}),
		)
	})

	require.Equal(t, grpcserver.DefName, server.Name())
	require.Equal(t, grpcserver.DefAddress, server.Address())
}

func TestConfigServiceError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = grpcserver.MustNew(
			grpcserver.WithConfig(configService, "key"),
			grpcserver.WithRegisterFn(func(*grpc.Server) {}),
		)
	})
}

func TestConfigWithoutRegisterFn(t *testing.T) {
	require.Panics(t, func() {
		_ = grpcserver.MustNew()
	})
}
