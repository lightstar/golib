package config_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/test/iotest"
)

func TestEnvJSON(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataJSON)
	defer iotest.RemoveFile(t, testConfigPath)

	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "")

	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataJSON)
}

func TestEnvYAML(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataYAML)
	defer iotest.RemoveFile(t, testConfigPath)

	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "yaml")

	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataYAML)
}

func TestEnvTOML(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataTOML)
	defer iotest.RemoveFile(t, testConfigPath)

	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "toml")

	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataTOML)
}

func TestEnvEtcd(t *testing.T) {
	endpoints := os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS")
	if endpoints == "" {
		t.Log("provide 'TEST_CONFIG_ETCD_ENDPOINTS' environment variable to test etcd source")
		return
	}

	setupEtcd(t)
	defer cleanEtcd(t)

	t.Setenv("CONFIG_FILE", "")
	t.Setenv("CONFIG_ETCD_ENDPOINTS", os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS"))
	t.Setenv("CONFIG_ETCD_KEY", etcdKey)
	t.Setenv("CONFIG_ENCODER", "json")

	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataJSON)
}

func TestEnvErrors(t *testing.T) {
	t.Setenv("CONFIG_FILE", testConfigPath)
	t.Setenv("CONFIG_ENCODER", "unknown")

	_, err := config.NewFromEnv()
	require.Error(t, err)

	t.Setenv("CONFIG_FILE", "")
	t.Setenv("CONFIG_ETCD_ENDPOINTS", "")
	t.Setenv("CONFIG_ETCD_KEY", "")
	t.Setenv("CONFIG_ENCODER", "")

	_, err = config.NewFromEnv()
	require.Error(t, err)
}
