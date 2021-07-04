package config_test

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"

	"github.com/lightstar/golib/pkg/config"
)

const etcdKey = "sample_config"

func TestEtcd(t *testing.T) {
	endpoints := os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS")
	if endpoints == "" {
		t.Log("provide 'TEST_CONFIG_ETCD_ENDPOINTS' environment variable to test etcd source")
		return
	}

	setupEtcd(t)
	defer cleanEtcd(t)

	cfg, err := config.NewFromEtcd(strings.Split(endpoints, ","), etcdKey, config.JSONEncoder)
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataJSON)
}

func TestEtcdErrors(t *testing.T) {
	endpoints := os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS")
	if endpoints == "" {
		t.Log("provide 'TEST_CONFIG_ETCD_ENDPOINTS' environment variable to test etcd source")
		return
	}

	_, err := config.NewFromEtcd([]string{}, etcdKey, config.JSONEncoder)
	require.Error(t, err)

	_, err = config.NewFromEtcd([]string{"unknown_address"}, etcdKey, config.JSONEncoder)
	require.EqualError(t, err, "etcd error")

	_, err = config.NewFromEtcd(strings.Split(endpoints, ","), etcdKey, config.JSONEncoder)
	require.Same(t, config.ErrNoData, err)

	setupEtcdWrong(t)
	defer cleanEtcd(t)

	_, err = config.NewFromEtcd(strings.Split(endpoints, ","), etcdKey, config.JSONEncoder)
	require.EqualError(t, err, "json error")
}

func setupEtcd(t *testing.T) {
	client := etcdClient(t)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Do(ctx, clientv3.OpPut(etcdKey, string(sampleConfigDataJSON)))
	require.NoError(t, err)
}

func setupEtcdWrong(t *testing.T) {
	client := etcdClient(t)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Do(ctx, clientv3.OpPut(etcdKey, string(sampleConfigDataWrongJSON)))
	require.NoError(t, err)
}

func cleanEtcd(t *testing.T) {
	client := etcdClient(t)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Do(ctx, clientv3.OpDelete(etcdKey))
	require.NoError(t, err)
}

func etcdClient(t *testing.T) *clientv3.Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS"), ","),
		DialTimeout: time.Second,
		LogConfig: &zap.Config{
			Level:            zap.NewAtomicLevel(),
			Encoding:         "json",
			OutputPaths:      nil,
			ErrorOutputPaths: nil,
		},
	})
	require.NoError(t, err)

	return client
}
