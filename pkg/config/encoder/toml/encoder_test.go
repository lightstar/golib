package toml_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/config/encoder/toml"
)

func TestTOMLEncoder(t *testing.T) {
	var result map[string]interface{}

	err := toml.Encoder.Encode(configtest.SampleConfigDataTOML, &result)
	require.NoError(t, err)

	require.Equal(t, configtest.ExpectedSampleRawDataTOML, result)
}

func TestTOMLEncoderError(t *testing.T) {
	var result map[string]interface{}

	err := toml.Encoder.Encode(configtest.SampleConfigDataWrongTOML, &result)
	require.ErrorContains(t, err, "toml error")
}
