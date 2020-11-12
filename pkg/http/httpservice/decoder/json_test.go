package decoder_test

import (
	"bytes"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/decoder"
)

func TestJsonDecoder(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{"key":"value", "number": 5}`))
	req := httptest.NewRequest("GET", "/", body)

	data := &struct {
		Key    string `json:"key"`
		Number int    `json:"number"`
	}{}

	err := decoder.JSON.Decode(req, data)
	require.NoError(t, err)

	require.Equal(t, &struct {
		Key    string `json:"key"`
		Number int    `json:"number"`
	}{
		Key:    "value",
		Number: 5,
	}, data)
}

func TestJsonDecoderError(t *testing.T) {
	body := bytes.NewBuffer([]byte(`{"key":"value",`))
	req := httptest.NewRequest("GET", "/", body)

	data := &struct {
		Key    string
		Number int
	}{}

	err := decoder.JSON.Decode(req, data)
	require.Error(t, err)
}
