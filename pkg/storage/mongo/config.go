package mongo

import (
	"time"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/errors"
)

const (
	// DefAddress is the default mongo server address.
	DefAddress = "127.0.0.1:27017"
	// DefConnectTimeout is the default mongo client connect timeout in seconds.
	DefConnectTimeout = 15
	// DefSocketTimeout is the default mongo client socket timeout in seconds.
	DefSocketTimeout = 30
)

// Config structure with client's configuration. Shouldn't be created manually.
type Config struct {
	address        string
	connectTimeout time.Duration
	socketTimeout  time.Duration
}

// Option function that is fed to NewClient and MustNewClient. Obtain them using 'With' functions down below.
type Option func(*Config) error

// WithConfig option retrieves configuration from provided configuration service.
//
// Example JSON configuration with all possible fields (if some are not present, defaults will be used):
//
//	{
//	    "address": "127.0.0.1:27017",
//	    "connectTimeout": 5,
//	    "socketTimeout": 10
//	}
func WithConfig(service config.Interface, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Address        string
			ConnectTimeout int
			SocketTimeout  int
		}{
			Address:        DefAddress,
			ConnectTimeout: DefConnectTimeout,
			SocketTimeout:  DefSocketTimeout,
		}

		err := service.GetByKey(key, &data)
		if err != nil && !errors.Is(err, config.ErrNoSuchKey) {
			return err
		}

		cfg.address = data.Address
		cfg.connectTimeout = time.Duration(data.ConnectTimeout) * time.Second
		cfg.socketTimeout = time.Duration(data.SocketTimeout) * time.Second

		return nil
	}
}

// WithAddress option applies provided mongo server address. Default: "127.0.0.1:27017".
func WithAddress(address string) Option {
	return func(cfg *Config) error {
		cfg.address = address
		return nil
	}
}

// WithConnectTimeout option applies provided mongo client connect timeout in seconds. Default: 15.
func WithConnectTimeout(idleTimeout int) Option {
	return func(cfg *Config) error {
		cfg.connectTimeout = time.Duration(idleTimeout) * time.Second
		return nil
	}
}

// WithSocketTimeout option applies provided mongo client socket timeout in seconds. Default: 30.
func WithSocketTimeout(socketTimeout int) Option {
	return func(cfg *Config) error {
		cfg.socketTimeout = time.Duration(socketTimeout) * time.Second
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{
		address:        DefAddress,
		connectTimeout: DefConnectTimeout * time.Second,
		socketTimeout:  DefSocketTimeout * time.Second,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
