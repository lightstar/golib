package httpserver

import (
	"net/http"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/log"
)

const (
	// DefAddress is the default server's address that will be listened to.
	DefAddress = "127.0.0.1:8080"
	// DefName is the default server's name.
	DefName = "http-server"
	// ReadHeaderTimeout is maximum time in seconds to read http header.
	ReadHeaderTimeout = 2
	// ReadTimeout is maximum time in seconds to read request.
	ReadTimeout = 3
	// WriteTimeout is maximum time in seconds to write request response.
	WriteTimeout = 5
)

// Config structure with server configuration. Shouldn't be created manually.
type Config struct {
	name    string
	address string
	logger  log.Logger
	handler http.Handler
}

// Option function that is fed to New and MustNew. Obtain them using 'With' functions down below.
type Option func(*Config) error

// WithConfig option retrieves configuration from provided configuration service.
//
// Example JSON configuration with all possible fields (if some are not present, defaults will be used):
//
//	{
//	    "name": "server-name",
//	    "address": "127.0.0.1:8080"
//	}
func WithConfig(service config.Interface, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Name    string
			Address string
		}{
			Name:    DefName,
			Address: DefAddress,
		}

		err := service.GetByKey(key, &data)
		if err != nil && !errors.Is(err, config.ErrNoSuchKey) {
			return err
		}

		cfg.name = data.Name
		cfg.address = data.Address

		return nil
	}
}

// WithName option applies provided server name. Default: "http-server".
func WithName(name string) Option {
	return func(cfg *Config) error {
		cfg.name = name
		return nil
	}
}

// WithAddress option applies provided address that will be listened to. Default: "127.0.0.1:8080".
func WithAddress(address string) Option {
	return func(cfg *Config) error {
		cfg.address = address
		return nil
	}
}

// WithLogger option applies provided logger. Default: standard logger with name equal to server's one.
func WithLogger(logger log.Logger) Option {
	return func(cfg *Config) error {
		cfg.logger = logger
		return nil
	}
}

// WithHandler option applies provided handler. Default: http.NotFoundHandler().
func WithHandler(handler http.Handler) Option {
	return func(cfg *Config) error {
		cfg.handler = handler
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{
		name:    DefName,
		address: DefAddress,
		handler: http.NotFoundHandler(),
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
