// Package httpservice provides simple http framework for generic http service. In general you will include instance of
// Service structure into your own specific service to reduce amount of written code.
//
// Typical simplified usage (in tandem with httpserver package):
//
//	srv := httpservice.New(httpservice.WithLogger(logger))
//
//	srv.UseMiddleware(middleware.Recover)
//	srv.UseMiddleware(middleware.Log)
//
//	srv.GET("/", "index", func (ctx *context.Context) error {
//	    // ... Handle request
//	})
//
//	// ... Define other routes
//
//	httpserver.New(httpserver.WithHandler(srv)).Run(ctx)
package httpservice

import (
	"sync"

	"github.com/julienschmidt/httprouter"

	"github.com/lightstar/golib/pkg/http/httpservice/context"
	"github.com/lightstar/golib/pkg/http/httpservice/decoder"
	"github.com/lightstar/golib/pkg/http/httpservice/encoder"
	"github.com/lightstar/golib/pkg/log"
)

// Service structure that provides http service functionality. Don't create manually, use the functions down below
// instead.
type Service struct {
	name  string
	debug bool

	logger     log.Logger
	encoder    encoder.Encoder
	decoder    decoder.Decoder
	middleware MiddlewareFunc
	router     *httprouter.Router

	pool sync.Pool
}

// New function creates new service with provided options.
func New(opts ...Option) (*Service, error) {
	config, err := buildConfig(opts)
	if err != nil {
		return nil, err
	}

	logger := config.logger
	if logger == nil {
		var err error

		logger, err = log.New(
			log.WithName(config.name),
			log.WithDebug(config.debug),
		)
		if err != nil {
			return nil, err
		}
	}

	service := &Service{
		name:    config.name,
		debug:   config.debug,
		logger:  logger,
		encoder: encoder.JSON,
		decoder: decoder.JSON,
		router:  httprouter.New(),
	}

	service.pool.New = func() interface{} {
		return context.New(service.logger)
	}

	return service, nil
}

// MustNew function creates new service with provided options and panics on any error.
func MustNew(opts ...Option) *Service {
	service, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return service
}

// Name method retrieves service's name.
func (service *Service) Name() string {
	return service.name
}

// Debug method retrieves service's debug flag.
func (service *Service) Debug() bool {
	return service.debug
}

// Logger method retrieves service's logger object.
func (service *Service) Logger() log.Logger {
	return service.logger
}

// UseEncoder method applies provided encoder that will be used for each route defined after that.
// Example:
//
//	srv.UseEncoder(enc1)
//	srv.GET(...) // will use encoder enc1
//	srv.GET(...) // will use encoder enc1 too
//	srv.UseEncoder(enc2)
//	srv.GET(...) // will use encoder enc2
func (service *Service) UseEncoder(enc encoder.Encoder) {
	service.encoder = enc
}

// UseDecoder method applies provided decoder that will be used for each route defined after that.
// Example:
//
//	srv.UseDecoder(dec1)
//	srv.GET(...) // will use decoder dec1
//	srv.GET(...) // will use decoder dec1 too
//	srv.UseDecoder(dec2)
//	srv.GET(...) // will use decoder dec2
func (service *Service) UseDecoder(dec decoder.Decoder) {
	service.decoder = dec
}

// error method handles error returned from handler function or middleware function if any.
func (service *Service) error(err error, c *context.Context) {
	service.logger.Error("[%s] %s error: %s", c.RemoteAddr(), c.Action(), err.Error())

	if !c.Response().HeaderWritten() {
		c.InternalErrorResponse("")
	}
}
