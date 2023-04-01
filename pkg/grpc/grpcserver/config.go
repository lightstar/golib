package grpcserver

import (
	"google.golang.org/grpc"

	"github.com/lightstar/golib/pkg/log"
)

const (
	// DefAddress is the default server's address that will be listened to.
	DefAddress = "127.0.0.1:50051"
	// DefName is the default server's name.
	DefName = "grpc-server"
)

// Config structure with server configuration. Shouldn't be created manually.
type Config struct {
	name       string
	address    string
	logger     log.Logger
	registerFn RegisterFn
}

// ConfigService interface used to obtain configuration from somewhere into some specific structure.
type ConfigService interface {
	GetByKey(string, interface{}) error
	IsNoSuchKeyError(error) bool
}

// Option function that is fed to New and MustNew. Obtain them using 'With' functions down below.
type Option func(*Config) error

// RegisterFn callback function that is called before listening in order to register any user-defined grpc service.
type RegisterFn func(*grpc.Server)

// WithConfig option retrieves configuration from provided configuration service.
//
// Example JSON configuration with all possible fields (if some are not present, defaults will be used):
//
//	{
//	    "name": "server-name",
//	    "address": "127.0.0.1:50051"
//	}
func WithConfig(service ConfigService, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Name    string
			Address string
		}{
			Name:    DefName,
			Address: DefAddress,
		}

		err := service.GetByKey(key, &data)
		if err != nil && !service.IsNoSuchKeyError(err) {
			return err
		}

		cfg.name = data.Name
		cfg.address = data.Address

		return nil
	}
}

// WithName option applies provided server name. Default: "grpc-server".
func WithName(name string) Option {
	return func(cfg *Config) error {
		cfg.name = name
		return nil
	}
}

// WithAddress option applies provided address that will be listened to. Default: "127.0.0.1:50051".
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

// WithRegisterFn option applies provided register function.
func WithRegisterFn(registerFn RegisterFn) Option {
	return func(cfg *Config) error {
		cfg.registerFn = registerFn
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{
		name:    DefName,
		address: DefAddress,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
