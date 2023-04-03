package file_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/internal/test/iotest"
	"github.com/lightstar/golib/pkg/config/encoder/yaml"
	"github.com/lightstar/golib/pkg/config/env/file"
)

const testConfigPath = "../../test/config_env2"

func TestFile(t *testing.T) {
	iotest.WriteFile(t, testConfigPath, configtest.SampleConfigDataYAML)
	defer iotest.RemoveFile(t, testConfigPath)

	t.Setenv("CONFIG_FILE", testConfigPath)

	cfg, err := file.NewConfig(yaml.Encoder)
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataYAML)
}
