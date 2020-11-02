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
	"sync"

	"github.com/lightstar/golib/pkg/log"
)

// Server structure that provides http server functionality.
type Server struct {
	name   string
	logger log.Logger
	server *http.Server

	wgRun *sync.WaitGroup
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

	wgRun := new(sync.WaitGroup)
	wgRun.Add(1)

	return &Server{
		name:   config.name,
		logger: logger,
		server: server,
		wgRun:  wgRun,
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

// Run method runs server listen loop. It should be called in a separate goroutine, otherwise it will be blocked.
// Call Shutdown method when you wish listening loop to stop.
func (server *Server) Run() {
	server.logger.Info("started")

	if err := server.server.ListenAndServe(); err != nil {
		if err != http.ErrServerClosed {
			server.logger.Error(err.Error())
		} else {
			server.logger.Info("stopped")
		}
	}

	server.wgRun.Done()
}

// Shutdown method gracefully shutdowns server waiting for all connections to complete.
// It should be called after Run method in another goroutine, otherwise it will hang forever waiting for Run to finish.
func (server *Server) Shutdown(ctx context.Context) {
	err := server.server.Shutdown(ctx)
	if err != nil {
		server.logger.Error(err.Error())
	}

	server.wgRun.Wait()
	server.logger.Sync()
}
