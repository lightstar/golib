package response_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/response"
)

func TestWriter(t *testing.T) {
	rec := httptest.NewRecorder()
	writer := &response.Writer{ResponseWriter: rec}

	writer.WriteHeader(http.StatusOK)
	writer.WriteHeader(http.StatusInternalServerError)

	_, err := writer.Write([]byte("some data"))
	require.NoError(t, err)

	require.Equal(t, true, writer.HeaderWritten())
	require.Equal(t, http.StatusOK, writer.Status())
	require.Equal(t, "some data", rec.Body.String())

	recNew := httptest.NewRecorder()
	writer.Reset(recNew)

	require.Equal(t, false, writer.HeaderWritten())
	require.Equal(t, 0, writer.Status())

	_, err = writer.Write([]byte("some data"))
	require.NoError(t, err)

	require.Equal(t, true, writer.HeaderWritten())
	require.Equal(t, http.StatusOK, writer.Status())
	require.Equal(t, "some data", recNew.Body.String())
}
