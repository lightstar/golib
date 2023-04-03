package file_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/internal/test/iotest"
	"github.com/lightstar/golib/pkg/config/encoder/json"
	"github.com/lightstar/golib/pkg/config/file"
)

const testConfigPath = "../../test/config"

func TestFile(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, configtest.SampleConfigDataJSON)
	defer iotest.RemoveFile(t, testConfigPath)

	cfg, err := file.NewConfig(testConfigPath, json.Encoder)
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataJSON)
}

func TestFileErrors(t *testing.T) {
	_, err := file.NewConfig(testConfigPath, json.Encoder)
	require.ErrorContains(t, err, "can't read from file")

	iotest.WriteFile(t, testConfigPath, configtest.SampleConfigDataWrongJSON)
	defer iotest.RemoveFile(t, testConfigPath)

	_, err = file.NewConfig(testConfigPath, json.Encoder)
	require.ErrorContains(t, err, "json error")
}
