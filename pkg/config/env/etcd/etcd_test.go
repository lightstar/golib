package etcd_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/config/encoder/json"
	"github.com/lightstar/golib/pkg/config/env/etcd"
)

const etcdKey = "sample_env2_config"

func TestEtcd(t *testing.T) {
	endpoints := os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS")
	if endpoints == "" {
		t.Log("provide 'TEST_CONFIG_ETCD_ENDPOINTS' environment variable to test etcd source")
		return
	}

	configtest.SetupEtcd(t, etcdKey)
	defer configtest.CleanEtcd(t, etcdKey)

	t.Setenv("CONFIG_ETCD_ENDPOINTS", os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS"))
	t.Setenv("CONFIG_ETCD_KEY", etcdKey)

	cfg, err := etcd.NewConfig(json.Encoder)
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataJSON)
}
