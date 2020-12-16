package mongo_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/storage/mongo"
	"github.com/lightstar/golib/pkg/test/configtest"
)

func TestConfig(t *testing.T) {
	var client *mongo.Client

	require.NotPanics(t, func() {
		client = mongo.MustNewClient(
			mongo.WithAddress("test_address"),
			mongo.WithConnectTimeout(120),
			mongo.WithSocketTimeout(180),
		)
	})

	defer client.Close()

	require.Equal(t, "test_address", client.Address())
	require.Equal(t, 120*time.Second, client.ConnectTimeout())
	require.Equal(t, 180*time.Second, client.SocketTimeout())
}

func TestConfigDefault(t *testing.T) {
	var client *mongo.Client

	require.NotPanics(t, func() {
		client = mongo.MustNewClient()
	})

	defer client.Close()

	require.Equal(t, mongo.DefAddress, client.Address())
	require.Equal(t, mongo.DefConnectTimeout*time.Second, client.ConnectTimeout())
	require.Equal(t, mongo.DefSocketTimeout*time.Second, client.SocketTimeout())
}

func TestConfigService(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Address        string
			ConnectTimeout int
			SocketTimeout  int
		}{
			Address:        "test_address",
			ConnectTimeout: 120,
			SocketTimeout:  180,
		},
	})

	var client *mongo.Client

	require.NotPanics(t, func() {
		client = mongo.MustNewClient(mongo.WithConfig(configService, "key"))
	})

	defer client.Close()

	require.Equal(t, "test_address", client.Address())
	require.Equal(t, 120*time.Second, client.ConnectTimeout())
	require.Equal(t, 180*time.Second, client.SocketTimeout())
}

func TestConfigServiceDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": config.ErrNoSuchKey,
	})

	var client *mongo.Client

	require.NotPanics(t, func() {
		client = mongo.MustNewClient(mongo.WithConfig(configService, "key"))
	})

	defer client.Close()

	require.Equal(t, mongo.DefAddress, client.Address())
	require.Equal(t, mongo.DefConnectTimeout*time.Second, client.ConnectTimeout())
	require.Equal(t, mongo.DefSocketTimeout*time.Second, client.SocketTimeout())
}

func TestConfigError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = mongo.MustNewClient(mongo.WithConfig(configService, "key"))
	})
}
