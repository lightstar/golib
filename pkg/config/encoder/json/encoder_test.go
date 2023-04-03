package json_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/config/encoder/json"
)

func TestJSONEncoder(t *testing.T) {
	var result map[string]interface{}

	err := json.Encoder.Encode(configtest.SampleConfigDataJSON, &result)
	require.NoError(t, err)

	require.Equal(t, configtest.ExpectedSampleRawDataJSON, result)
}

func TestJSONEncoderError(t *testing.T) {
	var result map[string]interface{}

	err := json.Encoder.Encode(configtest.SampleConfigDataWrongJSON, &result)
	require.ErrorContains(t, err, "json error")
}
