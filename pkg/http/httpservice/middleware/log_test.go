package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/context"
	"github.com/lightstar/golib/pkg/http/httpservice/middleware"
	"github.com/lightstar/golib/pkg/log"
	"github.com/lightstar/golib/pkg/test/iotest"
)

func TestLog(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithDebug(true),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	ctx := context.New(logger)
	ctx.Reset(rec, req, nil, nil, "some action")

	err := middleware.Log(logger)(func(ctx *context.Context) error {
		ctx.SetResult("some result")
		ctx.Response().WriteHeader(http.StatusOK)
		_, err := ctx.Response().Write([]byte("some response"))
		return err
	})(ctx)

	require.NoError(t, err)

	require.Equal(t, "some result", ctx.Result())
	require.Equal(t, http.StatusOK, ctx.Status())
	require.Equal(t, "some response", rec.Body.String())

	require.Regexp(t, `^\[\d{4}\-\d{2}\-\d{2} \d{2}:\d{2}:\d{2}\] \[\d+\.\d+.\d+.\d+:\d+\] `+
		`/ \[200\] some action: some result \(\d+\.\d+ms\)\n$`, stdout.String())
	require.Empty(t, stderr.String())
}
