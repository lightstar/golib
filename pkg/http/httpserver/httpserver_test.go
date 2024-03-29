//nolint:noctx,bodyclose // we don't care much about contexts and resource leaks in tests
package httpserver_test

import (
	"context"
	"io"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/iotest"
	"github.com/lightstar/golib/pkg/http/httpserver"
	"github.com/lightstar/golib/pkg/log"
)

func TestServer(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-server"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	server := httpserver.MustNew(
		httpserver.WithName("test-server"),
		httpserver.WithAddress("127.0.0.1:9090"),
		httpserver.WithHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("test response"))
		})),
		httpserver.WithLogger(logger),
	)

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "127.0.0.1:9090", server.Address())

	ctx, cancel := context.WithCancel(context.Background())
	stopChan := make(chan struct{})

	go func() {
		server.Run(ctx)
		close(stopChan)
	}()

	time.Sleep(2 * time.Second)

	client := &http.Client{
		Timeout: time.Second,
	}

	resp, err := client.Get("http://127.0.0.1:9090")
	require.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, "test response", string(body))

	err = resp.Body.Close()
	require.NoError(t, err)

	cancel()
	<-stopChan

	_, err = client.Get("http://127.0.0.1:9090")
	require.Error(t, err)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}

func TestServerWait(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-server"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	handlerEnteredChan := make(chan struct{})
	handlerFinished := false

	server := httpserver.MustNew(
		httpserver.WithName("test-server"),
		httpserver.WithAddress("127.0.0.1:9090"),
		httpserver.WithHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			close(handlerEnteredChan)
			time.Sleep(time.Second)
			handlerFinished = true
		})),
		httpserver.WithLogger(logger),
	)

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "127.0.0.1:9090", server.Address())

	ctx, cancel := context.WithCancel(context.Background())
	stopChan := make(chan struct{})

	go func() {
		server.Run(ctx)
		close(stopChan)
	}()

	time.Sleep(2 * time.Second)

	go func() {
		_, _ = http.Get("http://127.0.0.1:9090")
	}()

	<-handlerEnteredChan
	cancel()
	<-stopChan

	require.True(t, handlerFinished)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}

func TestServerError(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-server"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	server := httpserver.MustNew(
		httpserver.WithName("test-server"),
		httpserver.WithAddress("-1:80"),
		httpserver.WithHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})),
		httpserver.WithLogger(logger),
	)

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "-1:80", server.Address())

	ctx, cancel := context.WithCancel(context.Background())
	stopChan := make(chan struct{})

	go func() {
		<-time.After(15 * time.Second)
		cancel()
	}()

	go func() {
		server.Run(ctx)
		close(stopChan)
	}()

	<-stopChan

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) started\n$`, stdout.String())
	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) listen tcp: `+
		`lookup -1: no such host`, stderr.String())
}
