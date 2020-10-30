package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/goworld/pkg/config"
	"github.com/lightstar/goworld/pkg/test/iotest"
)

const testConfigPath = "../../test/config"

func TestFileJSON(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataJSON)
	defer iotest.RemoveFile(t, testConfigPath)

	cfg, err := config.NewFromFile(testConfigPath, config.JSONEncoder)
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataJSON)
}

func TestFileYAML(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataYAML)
	defer iotest.RemoveFile(t, testConfigPath)

	cfg, err := config.NewFromFile(testConfigPath, config.YAMLEncoder)
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataYAML)
}

func TestFileToml(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, sampleConfigDataTOML)
	defer iotest.RemoveFile(t, testConfigPath)

	cfg, err := config.NewFromFile(testConfigPath, config.TOMLEncoder)
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataTOML)
}

func TestFileErrors(t *testing.T) {
	_, err := config.NewFromFile(testConfigPath, config.JSONEncoder)
	require.EqualError(t, err, "file error")

	iotest.WriteFile(t, testConfigPath, sampleConfigDataWrongJSON)
	defer iotest.RemoveFile(t, testConfigPath)

	_, err = config.NewFromFile(testConfigPath, config.JSONEncoder)
	require.EqualError(t, err, "json error")

	iotest.WriteFile(t, testConfigPath, sampleConfigDataWrongYAML)

	_, err = config.NewFromFile(testConfigPath, config.YAMLEncoder)
	require.EqualError(t, err, "yaml error")

	iotest.WriteFile(t, testConfigPath, sampleConfigDataWrongTOML)

	_, err = config.NewFromFile(testConfigPath, config.TOMLEncoder)
	require.EqualError(t, err, "toml error")
}
