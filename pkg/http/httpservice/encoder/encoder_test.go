package encoder_test

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/encoder"
)

func TestPlainEncoder(t *testing.T) {
	testStringEncoder(t, encoder.Plain, "test data", "text/plain; charset=utf-8")
}

func TestHTMLEncoder(t *testing.T) {
	testStringEncoder(t, encoder.HTML, "test html", "text/html; charset=utf-8")
}

func testStringEncoder(t *testing.T, enc encoder.Encoder, data string, contentType string) {
	t.Helper()

	testEncoder(t, enc, data, contentType, data)
	testEncoder(t, enc, []byte(data), contentType, data)
	testEncoder(t, enc, errors.New(data), contentType, data)
	testEncoder(t, enc, stringer{text: data}, contentType, data)
	testEncoder(t, enc, misc{text: data}, contentType, data)
}

func testEncoder(t *testing.T, enc encoder.Encoder, data interface{}, expectedContentType string, expectedBody string) {
	t.Helper()

	rec := httptest.NewRecorder()

	err := enc.Encode(rec, http.StatusOK, data)
	require.NoError(t, err)

	require.Equal(t, expectedContentType, rec.Header().Get("Content-Type"))
	require.Equal(t, http.StatusOK, rec.Code)
	require.Equal(t, expectedBody, rec.Body.String())
}

type stringer struct {
	text string
}

func (s stringer) String() string {
	return s.text
}

type misc struct {
	text string
}

func (m misc) Format(s fmt.State, _ rune) {
	_, _ = io.WriteString(s, m.text)
}
