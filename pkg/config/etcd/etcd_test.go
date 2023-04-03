package etcd_test

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/config/encoder/json"
	"github.com/lightstar/golib/pkg/config/etcd"
)

const etcdKey = "sample_config"

func TestEtcd(t *testing.T) {
	endpoints := os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS")
	if endpoints == "" {
		t.Log("provide 'TEST_CONFIG_ETCD_ENDPOINTS' environment variable to test etcd source")
		return
	}

	configtest.SetupEtcd(t, etcdKey)
	defer configtest.CleanEtcd(t, etcdKey)

	cfg, err := etcd.NewConfig(strings.Split(endpoints, ","), etcdKey, json.Encoder)
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataJSON)
}

func TestEtcdErrors(t *testing.T) {
	endpoints := os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS")
	if endpoints == "" {
		t.Log("provide 'TEST_CONFIG_ETCD_ENDPOINTS' environment variable to test etcd source")
		return
	}

	_, err := etcd.NewConfig([]string{}, etcdKey, json.Encoder)
	require.Error(t, err)

	_, err = etcd.NewConfig([]string{"unknown_address"}, etcdKey, json.Encoder)
	require.ErrorContains(t, err, "etcd error")

	_, err = etcd.NewConfig(strings.Split(endpoints, ","), etcdKey, json.Encoder)
	require.Same(t, etcd.ErrNoData, err)

	configtest.SetupEtcdWrong(t, etcdKey)
	defer configtest.CleanEtcd(t, etcdKey)

	_, err = etcd.NewConfig(strings.Split(endpoints, ","), etcdKey, json.Encoder)
	require.ErrorContains(t, err, "json error")
}
