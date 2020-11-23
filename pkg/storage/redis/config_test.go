package redis_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/storage/redis"
	"github.com/lightstar/golib/pkg/test/configtest"
)

func TestConfig(t *testing.T) {
	var client *redis.Client

	require.NotPanics(t, func() {
		client = redis.MustNewClient(
			redis.WithAddress("test_address"),
			redis.WithMaxIdle(10),
			redis.WithIdleTimeout(1800),
		)
	})

	defer client.Close()

	require.Equal(t, "test_address", client.Address())
	require.Equal(t, 10, client.MaxIdle())
	require.Equal(t, 1800*time.Second, client.IdleTimeout())
}

func TestConfigDefault(t *testing.T) {
	var client *redis.Client

	require.NotPanics(t, func() {
		client = redis.MustNewClient()
	})

	defer client.Close()

	require.Equal(t, redis.DefAddress, client.Address())
	require.Equal(t, redis.DefMaxIdle, client.MaxIdle())
	require.Equal(t, redis.DefIdleTimeout*time.Second, client.IdleTimeout())
}

func TestConfigService(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Address     string
			MaxIdle     int
			IdleTimeout int
		}{
			Address:     "test_address",
			MaxIdle:     10,
			IdleTimeout: 1800,
		},
	})

	var client *redis.Client

	require.NotPanics(t, func() {
		client = redis.MustNewClient(redis.WithConfig(configService, "key"))
	})

	defer client.Close()

	require.Equal(t, "test_address", client.Address())
	require.Equal(t, 10, client.MaxIdle())
	require.Equal(t, 1800*time.Second, client.IdleTimeout())
}

func TestConfigServiceDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": config.ErrNoSuchKey,
	})

	var client *redis.Client

	require.NotPanics(t, func() {
		client = redis.MustNewClient(redis.WithConfig(configService, "key"))
	})

	defer client.Close()

	require.Equal(t, redis.DefAddress, client.Address())
	require.Equal(t, redis.DefMaxIdle, client.MaxIdle())
	require.Equal(t, redis.DefIdleTimeout*time.Second, client.IdleTimeout())
}

func TestConfigServiceError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = redis.MustNewClient(redis.WithConfig(configService, "key"))
	})
}
