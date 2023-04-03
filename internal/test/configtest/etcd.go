package configtest

import (
	"context"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
)

// SetupEtcd function prepares etcd for testing by setting provided key to sample config JSON data.
func SetupEtcd(t *testing.T, key string) {
	t.Helper()

	client := etcdClient(t)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Do(ctx, clientv3.OpPut(key, string(SampleConfigDataJSON)))
	require.NoError(t, err)
}

// SetupEtcdWrong function prepares etcd for testing by setting provided key to sample wrong config JSON data.
func SetupEtcdWrong(t *testing.T, key string) {
	t.Helper()

	client := etcdClient(t)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Do(ctx, clientv3.OpPut(key, string(SampleConfigDataWrongJSON)))
	require.NoError(t, err)
}

// CleanEtcd function clears provided key to restore etcd to its original state.
func CleanEtcd(t *testing.T, key string) {
	t.Helper()

	client := etcdClient(t)
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := client.Do(ctx, clientv3.OpDelete(key))
	require.NoError(t, err)
}

func etcdClient(t *testing.T) *clientv3.Client {
	t.Helper()

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
