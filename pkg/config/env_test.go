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

	os.Setenv("CONFIG_FILE", testConfigPath)
	os.Setenv("CONFIG_ENCODER", "")

	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataJSON)
}

func TestEnvYAML(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataYAML)
	defer iotest.RemoveFile(t, testConfigPath)

	os.Setenv("CONFIG_FILE", testConfigPath)
	os.Setenv("CONFIG_ENCODER", "yaml")

	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataYAML)
}

func TestEnvTOML(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataTOML)
	defer iotest.RemoveFile(t, testConfigPath)

	os.Setenv("CONFIG_FILE", testConfigPath)
	os.Setenv("CONFIG_ENCODER", "toml")

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

	os.Setenv("CONFIG_FILE", "")
	os.Setenv("CONFIG_ETCD_ENDPOINTS", os.Getenv("TEST_CONFIG_ETCD_ENDPOINTS"))
	os.Setenv("CONFIG_ETCD_KEY", etcdKey)
	os.Setenv("CONFIG_ENCODER", "json")

	cfg, err := config.NewFromEnv()
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataJSON)
}

func TestEnvErrors(t *testing.T) {
	os.Setenv("CONFIG_FILE", testConfigPath)
	os.Setenv("CONFIG_ENCODER", "unknown")

	_, err := config.NewFromEnv()
	require.Error(t, err)

	os.Setenv("CONFIG_FILE", "")
	os.Setenv("CONFIG_ETCD_ENDPOINTS", "")
	os.Setenv("CONFIG_ETCD_KEY", "")
	os.Setenv("CONFIG_ENCODER", "")

	_, err = config.NewFromEnv()
	require.Error(t, err)
}
