// Package httpserver provides API to standard http server with graceful shutdown.
//
// Typical usage:
//      server := httpserver.New(
//          httpserver.WithName("my-server"),
//          httpserver.WithAddress("127.0.0.1:8080"),
//          httpserver.WithHandler(myHandler),
//      )
//      go server.Run()
//      // ... Initialize other services, run daemon loop, etc
//      server.Shutdown()
package httpserver

import (
	"context"
	"net/http"

	"github.com/lightstar/golib/pkg/log"
)

// Server structure that provides http server functionality.
type Server struct {
	name   string
	logger log.Logger
	server *http.Server
}

// New function creates new server with provided options.
func New(opts ...Option) (*Server, error) {
	config, err := buildConfig(opts)
	if err != nil {
		return nil, err
	}

	logger := config.logger
	if logger == nil {
		var err error

		logger, err = log.New(log.WithName(config.name))
		if err != nil {
			return nil, err
		}
	}

	server := &http.Server{
		Addr:    config.address,
		Handler: config.handler,
	}

	return &Server{
		name:   config.name,
		logger: logger,
		server: server,
	}, nil
}

// MustNew function creates new server with provided options and panics on any error.
func MustNew(opts ...Option) *Server {
	server, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return server
}

// Name method gets server's name.
func (server *Server) Name() string {
	return server.name
}

// Address method gets server's address that it listens to.
func (server *Server) Address() string {
	return server.server.Addr
}

// Run method runs server listen loop. It is blocking so you probably want to run it in a separate goroutine.
// If you pass cancellable context here, you will be able to gracefully shutdown server that waits for all requests
// to complete.
//
// Only upgraded connections (such as websocket ones) will not be waited for, you will need to shutdown
// them manually.
func (server *Server) Run(ctx context.Context) {
	stopChan := make(chan struct{})

	server.logger.Info("started")

	go func() {
		if err := server.server.ListenAndServe(); err != nil {
			if err == http.ErrServerClosed {
				server.logger.Info("stopped")
			} else {
				server.logger.Error(err.Error())
			}
		}

		close(stopChan)
	}()

	select {
	case <-stopChan:
	case <-ctx.Done():
	}

	err := server.server.Shutdown(context.Background())
	if err != nil {
		server.logger.Error(err.Error())
	}

	<-stopChan

	server.logger.Sync()
}
