package context_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/http/httpservice/context"
	"github.com/lightstar/golib/pkg/http/httpservice/decoder"
	"github.com/lightstar/golib/pkg/http/httpservice/encoder"
	"github.com/lightstar/golib/pkg/log"
)

func TestContext(t *testing.T) {
	logger := log.NewNop()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", strings.NewReader(`{"message":"test"}`))
	params := NewParams()

	ctx := context.New(logger)

	ctx.Reset(rec, req, params, encoder.Func(func(w http.ResponseWriter, status int, data interface{}) error {
		w.WriteHeader(status)
		_, err := w.Write([]byte(fmt.Sprintf("%s", data)))
		return err
	}), decoder.Func(func(r *http.Request, data interface{}) error {
		return json.NewDecoder(r.Body).Decode(data)
	}), "test action")

	require.Same(t, logger, ctx.Logger())
	require.Same(t, rec, ctx.Response().ResponseWriter)
	require.Same(t, req, ctx.Request())
	require.Same(t, params, ctx.Params())

	require.Equal(t, "unknown", ctx.Result())
	require.Equal(t, "test action", ctx.Action())
	require.Regexp(t, `^\d+\.\d+\.\d+\.\d+:\d+$`, ctx.RemoteAddr())
	require.Equal(t, "/", ctx.RequestURI())

	data := struct {
		Message string
	}{}

	err := ctx.Decode(&data)
	require.NoError(t, err)

	require.Equal(t, &struct {
		Message string
	}{
		Message: "test",
	}, &data)

	ctx.SetResult("test result")
	require.Equal(t, "test result", ctx.Result())

	err = ctx.Encode(http.StatusOK, "some data")
	require.NoError(t, err)

	require.Equal(t, "some data", rec.Body.String())
	require.Equal(t, http.StatusOK, ctx.Status())
}

func TestContextEncodeError(t *testing.T) {
	logger := log.NewNop()
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	ctx := context.New(logger)

	encodeError := errors.New("test encode error")

	ctx.Reset(rec, req, nil, encoder.Func(func(w http.ResponseWriter, status int, data interface{}) error {
		return encodeError
	}), nil, "")

	err := ctx.Encode(http.StatusOK, "some data")

	require.Error(t, err)
	require.True(t, errors.Is(err, encodeError))
}

func TestContextDecodeError(t *testing.T) {
	logger := log.NewNop()

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	ctx := context.New(logger)

	decodeError := errors.New("test decode error")

	ctx.Reset(rec, req, nil, nil, decoder.Func(func(r *http.Request, data interface{}) error {
		return decodeError
	}), "")

	data := struct {
		Message string
	}{}

	err := ctx.Decode(&data)

	require.Error(t, err)
	require.True(t, errors.Is(err, decodeError))
}
