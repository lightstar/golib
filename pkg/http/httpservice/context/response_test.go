package context_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/http/httpservice/context"
	"github.com/lightstar/golib/pkg/http/httpservice/encoder"
	"github.com/lightstar/golib/pkg/log"
	"github.com/lightstar/golib/pkg/test/iotest"
)

type Test struct {
	name     string
	fun      func(*context.Context)
	result   string
	status   int
	response string
}

// nolint: funlen // tests slice is too long to pass that linter
func TestResponse(t *testing.T) {
	logger := log.NewNop()

	tests := []Test{
		{
			name: "InternalErrorDefault",
			fun: func(ctx *context.Context) {
				ctx.InternalErrorResponse("")
			},
			result:   "internal error",
			status:   http.StatusInternalServerError,
			response: "internal error",
		},
		{
			name: "InternalError",
			fun: func(ctx *context.Context) {
				ctx.InternalErrorResponse("some error")
			},
			result:   "some error",
			status:   http.StatusInternalServerError,
			response: "internal error",
		},
		{
			name: "UnauthorizedDefault",
			fun: func(ctx *context.Context) {
				ctx.UnauthorizedResponse("")
			},
			result:   "unauthorized",
			status:   http.StatusUnauthorized,
			response: "unauthorized",
		},
		{
			name: "Unauthorized",
			fun: func(ctx *context.Context) {
				ctx.UnauthorizedResponse("some error")
			},
			result:   "some error",
			status:   http.StatusUnauthorized,
			response: "unauthorized",
		},
		{
			name: "BadRequestDefault",
			fun: func(ctx *context.Context) {
				ctx.BadRequestResponse(nil, "")
			},
			result:   "bad request",
			status:   http.StatusBadRequest,
			response: "bad request",
		},
		{
			name: "BadRequest",
			fun: func(ctx *context.Context) {
				ctx.BadRequestResponse(nil, "some error")
			},
			result:   "some error",
			status:   http.StatusBadRequest,
			response: "bad request",
		},
		{
			name: "BadRequestCustomResponse",
			fun: func(ctx *context.Context) {
				ctx.BadRequestResponse("some response", "some error")
			},
			result:   "some error",
			status:   http.StatusBadRequest,
			response: "some response",
		},
		{
			name: "OKDefault",
			fun: func(ctx *context.Context) {
				ctx.OKResponse(nil, "")
			},
			result:   "ok",
			status:   http.StatusOK,
			response: "OK",
		},
		{
			name: "OK",
			fun: func(ctx *context.Context) {
				ctx.OKResponse(nil, "some result")
			},
			result:   "some result",
			status:   http.StatusOK,
			response: "OK",
		},
		{
			name: "OKCustomResponse",
			fun: func(ctx *context.Context) {
				ctx.OKResponse("some response", "some result")
			},
			result:   "some result",
			status:   http.StatusOK,
			response: "some response",
		},
	}

	for _, test := range tests {
		func(test Test) {
			t.Run(test.name, func(t *testing.T) {
				rec := httptest.NewRecorder()
				req := httptest.NewRequest("GET", "/", nil)

				ctx := context.New(logger)

				ctx.Reset(rec, req, nil,
					encoder.Func(func(w http.ResponseWriter, status int, data interface{}) error {
						w.WriteHeader(status)
						_, err := w.Write([]byte(fmt.Sprintf("%s", data)))
						return err
					}), nil, "")

				test.fun(ctx)

				require.Equal(t, test.result, ctx.Result())
				require.Equal(t, test.status, ctx.Status())
				require.Equal(t, test.response, rec.Body.String())
			})
		}(test)
	}
}

func TestResponseEncodeError(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	rec := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	ctx := context.New(logger)

	ctx.Reset(rec, req, nil, encoder.Func(func(w http.ResponseWriter, status int, data interface{}) error {
		return errors.New("encode error")
	}), nil, "")

	ctx.DoResponse("some data", http.StatusOK, "some result")

	require.Equal(t, "some result", ctx.Result())
	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] `+
		`can't encode data 'some data' \(encode error\)\n`, stderr.String())
}
