package httpservice_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/iotest"
	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/http/httpservice"
	"github.com/lightstar/golib/pkg/http/httpservice/context"
	"github.com/lightstar/golib/pkg/http/httpservice/decoder"
	"github.com/lightstar/golib/pkg/http/httpservice/encoder"
	"github.com/lightstar/golib/pkg/log"
)

type (
	SetupFunc func(*httpservice.Service, httpservice.HandlerFunc, map[string]string)
	TestFunc  func(*testing.T, *httpservice.Service, map[string]string)
)

type Test struct {
	name      string
	method    string
	status    int
	response  string
	stdout    string
	stderr    string
	setupFunc SetupFunc
	testFunc  TestFunc
}

//nolint:funlen // tests slice is too long to pass that linter
func TestService(t *testing.T) {
	middlewareSetupFunc := func(useFunc func(httpservice.MiddlewareFunc)) SetupFunc {
		return func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
			param["test"] = ""

			useFunc(func(h httpservice.HandlerFunc) httpservice.HandlerFunc {
				return func(ctx *context.Context) error {
					param["test"] += " first"
					return h(ctx)
				}
			})

			useFunc(func(h httpservice.HandlerFunc) httpservice.HandlerFunc {
				return func(ctx *context.Context) error {
					param["test"] += " second"
					return h(ctx)
				}
			})

			service.GET("/test", "", handler)
		}
	}

	tests := []Test{
		{
			name:   "GET",
			method: "GET",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				service.GET("/test", "", handler)
			},
		},
		{
			name:   "POST",
			method: "POST",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				service.POST("/test", "", handler)
			},
		},
		{
			name:   "PUT",
			method: "PUT",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				service.PUT("/test", "", handler)
			},
		},
		{
			name:   "DELETE",
			method: "DELETE",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				service.DELETE("/test", "", handler)
			},
		},
		{
			name:   "OPTIONS",
			method: "OPTIONS",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				service.OPTIONS("/test", "", handler)
			},
		},
		{
			name:   "Middleware",
			method: "GET",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				middlewareSetupFunc(service.UseMiddleware)(service, handler, param)
			},
			testFunc: func(t *testing.T, service *httpservice.Service, param map[string]string) {
				t.Helper()
				require.Equal(t, " first second", param["test"])
			},
		},
		{
			name:   "MiddlewareBefore",
			method: "GET",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				middlewareSetupFunc(service.UseMiddlewareBefore)(service, handler, param)
			},
			testFunc: func(t *testing.T, service *httpservice.Service, param map[string]string) {
				t.Helper()
				require.Equal(t, " second first", param["test"])
			},
		},
		{
			name:     "Error",
			method:   "GET",
			status:   http.StatusInternalServerError,
			response: "internal error",
			stderr: `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \[\d+\.\d+\.\d+\.\d+:\d+\] ` +
				`test action error: test error` + "\n",
			setupFunc: func(service *httpservice.Service, handler httpservice.HandlerFunc, param map[string]string) {
				service.UseMiddleware(func(h httpservice.HandlerFunc) httpservice.HandlerFunc {
					return func(ctx *context.Context) error {
						return errors.New("test error")
					}
				})

				service.GET("/test", "test action", handler)
			},
		},
	}

	for _, test := range tests {
		func(test Test) {
			t.Run(test.name, func(t *testing.T) {
				testService(t, test)
			})
		}(test)
	}
}

func testService(t *testing.T, test Test) {
	t.Helper()

	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	enc := encoder.Func(func(w http.ResponseWriter, status int, data interface{}) error {
		w.WriteHeader(status)
		_, err := w.Write([]byte(fmt.Sprintf("%s", data)))
		return err
	})

	dec := decoder.Func(func(r *http.Request, data interface{}) error {
		return json.NewDecoder(r.Body).Decode(data)
	})

	handler := func(ctx *context.Context) error {
		data := struct {
			Message string
		}{}

		if err := ctx.Decode(&data); err != nil {
			return err
		}

		return ctx.Encode(http.StatusOK, data.Message)
	}

	param := make(map[string]string)

	service, err := httpservice.New(
		httpservice.WithName("test-service"),
		httpservice.WithDebug(true),
		httpservice.WithLogger(logger),
	)
	require.NoError(t, err)

	service.UseEncoder(enc)
	service.UseDecoder(dec)

	require.Equal(t, "test-service", service.Name())
	require.Equal(t, true, service.Debug())
	require.Same(t, logger, service.Logger())

	test.setupFunc(service, handler, param)

	rec := httptest.NewRecorder()
	r := httptest.NewRequest(test.method, "/test", strings.NewReader(`{"message":"OK"}`))

	service.ServeHTTP(rec, r)

	status := http.StatusOK
	if test.status != 0 {
		status = test.status
	}

	response := "OK"
	if test.response != "" {
		response = test.response
	}

	require.Equal(t, status, rec.Code)
	require.Equal(t, response, rec.Body.String())
	require.Regexp(t, test.stdout, stdout.String())
	require.Regexp(t, test.stderr, stderr.String())

	if test.testFunc != nil {
		test.testFunc(t, service, param)
	}
}
