// Package grpcserver provides API to grpc server with graceful shutdown.
//
// Typical usage:
//
//	grpcserver.New(
//	    httpserver.WithName("my-server"),
//	    httpserver.WithAddress("127.0.0.1:8080"),
//	    httpserver.WithRegisterFn(func(s *grpc.Server) {
//	        // Register your service here
//	    }),
//	).Run(ctx)
package grpcserver

import (
	"context"
	"net"

	"google.golang.org/grpc"

	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/log"
)

// Server structure that provides grpc server functionality. Don't create manually, use the functions down below
// instead.
type Server struct {
	name       string
	address    string
	logger     log.Logger
	registerFn RegisterFn
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

	if config.registerFn == nil {
		return nil, errors.New("need register function")
	}

	return &Server{
		name:       config.name,
		address:    config.address,
		logger:     logger,
		registerFn: config.registerFn,
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
	return server.address
}

// Run method runs server listen loop. It is blocking so you probably want to run it in a separate goroutine.
// If you pass cancellable context here, you will be able to gracefully shutdown server that waits for all requests
// to complete.
func (server *Server) Run(ctx context.Context) {
	stopChan := make(chan struct{})
	grpcServer := grpc.NewServer()

	server.logger.Info("started")

	go func() {
		defer close(stopChan)

		listener, err := net.Listen("tcp", server.address)
		if err != nil {
			server.logger.Error(err.Error())
			return
		}

		server.registerFn(grpcServer)

		if err := grpcServer.Serve(listener); err != nil {
			server.logger.Error(err.Error())
		} else {
			server.logger.Info("stopped")
		}
	}()

	select {
	case <-stopChan:
	case <-ctx.Done():
	}

	grpcServer.GracefulStop()

	<-stopChan

	server.logger.Sync()
}
