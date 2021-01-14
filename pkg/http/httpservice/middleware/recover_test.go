package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/context"
	"github.com/lightstar/golib/pkg/http/httpservice/middleware"
	"github.com/lightstar/golib/pkg/log"
)

func TestRecover(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	ctx := context.New(log.NewNop())
	ctx.Reset(rec, req, nil, nil, nil, "some action")

	err := middleware.Recover()(func(ctx *context.Context) error {
		panic("some panic")
	})(ctx)

	require.Error(t, err)
	require.Equal(t, err.Error(), "recovered from panic: some panic")
}

func TestRecoverNoPanic(t *testing.T) {
	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	ctx := context.New(log.NewNop())
	ctx.Reset(rec, req, nil, nil, nil, "some action")

	err := middleware.Recover()(func(ctx *context.Context) error {
		ctx.SetResult("some result")
		ctx.Response().WriteHeader(http.StatusOK)
		_, err := ctx.Response().Write([]byte("some response"))
		return err
	})(ctx)

	require.NoError(t, err)

	require.Equal(t, "some result", ctx.Result())
	require.Equal(t, http.StatusOK, ctx.Status())
	require.Equal(t, "some response", rec.Body.String())
}
