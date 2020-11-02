// nolint: noctx, bodyclose // we don't care much about contexts and resource leaks in tests
package httpserver_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/http/httpserver"
	"github.com/lightstar/golib/pkg/log"
	"github.com/lightstar/golib/pkg/test/configtest"
	"github.com/lightstar/golib/pkg/test/iotest"
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
		httpserver.WithAddress("127.0.0.1:6060"),
		httpserver.WithHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, _ = w.Write([]byte("test response"))
		})),
		httpserver.WithLogger(logger),
	)

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "127.0.0.1:6060", server.Address())

	go server.Run()

	client := &http.Client{
		Timeout: time.Second,
	}

	resp, err := client.Get("http://127.0.0.1:6060")
	require.NoError(t, err)

	body, err := ioutil.ReadAll(resp.Body)
	require.NoError(t, err)

	require.Equal(t, "test response", string(body))

	err = resp.Body.Close()
	require.NoError(t, err)

	server.Shutdown(context.Background())

	_, err = client.Get("http://127.0.0.1:6060")
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

	wg := new(sync.WaitGroup)
	wg.Add(1)

	handlerFinished := false

	server := httpserver.MustNew(
		httpserver.WithName("test-server"),
		httpserver.WithAddress("127.0.0.1:6060"),
		httpserver.WithHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wg.Done()
			time.Sleep(time.Second)
			handlerFinished = true
		})),
		httpserver.WithLogger(logger),
	)

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "127.0.0.1:6060", server.Address())

	go server.Run()

	go func() {
		_, _ = http.Get("http://127.0.0.1:6060")
	}()

	wg.Wait()
	server.Shutdown(context.Background())

	require.True(t, handlerFinished)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}

func TestConfig(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Name    string
			Address string
		}{
			Name:    "test-server",
			Address: "test-address",
		},
	})

	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "test-address", server.Address())
}

func TestConfigDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": config.ErrNoSuchKey,
	})

	var server *httpserver.Server

	require.NotPanics(t, func() {
		server = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})

	require.Equal(t, httpserver.DefName, server.Name())
	require.Equal(t, httpserver.DefAddress, server.Address())
}

func TestConfigError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = httpserver.MustNew(httpserver.WithConfig(configService, "key"))
	})
}
