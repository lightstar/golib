package encoder_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/encoder"
)

func TestJsonEncoder(t *testing.T) {
	data := &struct {
		Key    string `json:"key"`
		Number int    `json:"number"`
	}{
		Key:    "value",
		Number: 5,
	}
	rec := httptest.NewRecorder()

	err := encoder.JSON.Encode(rec, http.StatusOK, data)
	require.NoError(t, err)

	require.Equal(t, "application/json; charset=utf-8", rec.Header().Get("Content-Type"))
	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, `{"key":"value","number":5}`+"\n", rec.Body.String())
}

func TestJsonEncoderError(t *testing.T) {
	err := encoder.JSON.Encode(httptest.NewRecorder(), http.StatusOK, &struct {
		Key func()
	}{})
	require.Error(t, err)
}
