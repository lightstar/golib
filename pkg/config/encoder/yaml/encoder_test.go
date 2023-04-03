package yaml_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/config/encoder/yaml"
)

func TestTOMLEncoder(t *testing.T) {
	var result map[string]interface{}

	err := yaml.Encoder.Encode(configtest.SampleConfigDataYAML, &result)
	require.NoError(t, err)

	require.Equal(t, configtest.ExpectedSampleRawDataYAML, result)
}

func TestTOMLEncoderError(t *testing.T) {
	var result map[string]interface{}

	err := yaml.Encoder.Encode(configtest.SampleConfigDataWrongYAML, &result)
	require.ErrorContains(t, err, "yaml error")
}
