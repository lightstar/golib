package env_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/internal/test/iotest"
	"github.com/lightstar/golib/pkg/config/env"
)

const (
	testConfigPath = "../../test/config_env"
	etcdKey        = "sample_env_config"
)

func TestEnvJSON(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, configtest.SampleConfigDataJSON)
	defer iotest.RemoveFile(t, testConfigPath)

	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "json")

	cfg, err := env.NewConfig()
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataJSON)
}

func TestEnvYAML(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, configtest.SampleConfigDataYAML)
	defer iotest.RemoveFile(t, testConfigPath)

	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "")

	cfg, err := env.NewConfig()
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataYAML)
}

func TestEnvTOML(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, configtest.SampleConfigDataTOML)
	defer iotest.RemoveFile(t, testConfigPath)

	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "toml")

	cfg, err := env.NewConfig()
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataTOML)
}

func TestEnvEtcd(t *testing.T) {
	endpoints := os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS")
	if endpoints == "" {
		t.Log("provide 'TEST_CONFIG_ETCD_ENDPOINTS' environment variable to test etcd source")
		return
	}

	configtest.SetupEtcd(t, etcdKey)
	defer configtest.CleanEtcd(t, etcdKey)

	t.Setenv("CONFIG_FILE", "")
	t.Setenv("CONFIG_ETCD_ENDPOINTS", os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS"))
	t.Setenv("CONFIG_ETCD_KEY", etcdKey)
	t.Setenv("CONFIG_ENCODER", "json")

	cfg, err := env.NewConfig()
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataJSON)
}

func TestEnvErrors(t *testing.T) {
	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "unknown")

	_, err := env.NewConfig()
	require.Error(t, err)

	t.Setenv("CONFIG_FILE", "")
	t.Setenv("CONFIG_ETCD_ENDPOINTS", "")
	t.Setenv("CONFIG_ETCD_KEY", "")
	t.Setenv("CONFIG_ENCODER", "")

	_, err = env.NewConfig()
	require.Error(t, err)
}
