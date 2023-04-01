package httpserver

import (
	"net/http"

	"github.com/lightstar/golib/pkg/log"
)

const (
	// DefAddress is the default server's address that will be listened to.
	DefAddress = "127.0.0.1:8080"
	// DefName is the default server's name.
	DefName = "http-server"
	// DefReadHeaderTimeout is the default maximum time in seconds to read http header.
	DefReadHeaderTimeout = 2
	// DefReadTimeout is the default maximum time in seconds to read request.
	DefReadTimeout = 3
	// DefWriteTimeout is the maximum time in seconds to write request response.
	DefWriteTimeout = 3
)

// Config structure with server configuration. Shouldn't be created manually.
type Config struct {
	name              string
	address           string
	readHeaderTimeout int64
	readTimeout       int64
	writeTimeout      int64

	logger  log.Logger
	handler http.Handler
}

// ConfigService interface used to obtain configuration from somewhere into some specific structure.
type ConfigService interface {
	GetByKey(string, interface{}) error
	IsNoSuchKeyError(error) bool
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
func WithConfig(service ConfigService, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Name              string
			Address           string
			ReadHeaderTimeout int64
			ReadTimeout       int64
			WriteTimeout      int64
		}{
			Name:              DefName,
			Address:           DefAddress,
			ReadHeaderTimeout: DefReadHeaderTimeout,
			ReadTimeout:       DefReadTimeout,
			WriteTimeout:      DefWriteTimeout,
		}

		err := service.GetByKey(key, &data)
		if err != nil && !service.IsNoSuchKeyError(err) {
			return err
		}

		cfg.name = data.Name
		cfg.address = data.Address
		cfg.readHeaderTimeout = data.ReadHeaderTimeout
		cfg.readTimeout = data.ReadTimeout
		cfg.writeTimeout = data.WriteTimeout

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

// WithReadHeaderTimeout option applies provided maximum time in seconds to read http header. Default: 2.
func WithReadHeaderTimeout(readHeaderTimeout int64) Option {
	return func(cfg *Config) error {
		cfg.readHeaderTimeout = readHeaderTimeout
		return nil
	}
}

// WithReadTimeout option applies provided maximum time in seconds to read request. Default: 3.
func WithReadTimeout(readTimeout int64) Option {
	return func(cfg *Config) error {
		cfg.readTimeout = readTimeout
		return nil
	}
}

// WithWriteTimeout option applies provided maximum time in seconds to write response. Default: 3.
func WithWriteTimeout(writeTimeout int64) Option {
	return func(cfg *Config) error {
		cfg.writeTimeout = writeTimeout
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
		name:              DefName,
		address:           DefAddress,
		readHeaderTimeout: DefReadHeaderTimeout,
		readTimeout:       DefReadTimeout,
		writeTimeout:      DefWriteTimeout,
		handler:           http.NotFoundHandler(),
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
