package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

const (
	// DefAddress is the default redis server address.
	DefAddress = "127.0.0.1:6379"
	// DefMaxIdle is the default maximum number of idle redis connections in the pool.
	DefMaxIdle = 3
	// DefIdleTimeout is the default timeout in seconds after which idle connections will be dropped away.
	DefIdleTimeout = 600
)

// Config structure with client's configuration. Shouldn't be created manually.
type Config struct {
	address     string
	maxIdle     int
	idleTimeout time.Duration
	dialFunc    func() (redis.Conn, error)
}

// ConfigService interface used to obtain configuration from somewhere into some specific structure.
type ConfigService interface {
	GetByKey(string, interface{}) error
	IsNoSuchKeyError(error) bool
}

// Option function that is fed to NewClient and MustNewClient. Obtain them using 'With' functions down below.
type Option func(*Config) error

// WithConfig option retrieves configuration from provided configuration service.
//
// Example JSON configuration with all possible fields (if some are not present, defaults will be used):
//
//	{
//	    "address": "127.0.0.1:6379",
//	    "maxIdle": 5,
//	    "idleTimeout": 300
//	}
func WithConfig(service ConfigService, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Address     string
			MaxIdle     int
			IdleTimeout int
		}{
			Address:     DefAddress,
			MaxIdle:     DefMaxIdle,
			IdleTimeout: DefIdleTimeout,
		}

		err := service.GetByKey(key, &data)
		if err != nil && !service.IsNoSuchKeyError(err) {
			return err
		}

		cfg.address = data.Address
		cfg.maxIdle = data.MaxIdle
		cfg.idleTimeout = time.Duration(data.IdleTimeout) * time.Second

		return nil
	}
}

// WithAddress option applies provided redis server address. Default: "127.0.0.1:6379".
func WithAddress(address string) Option {
	return func(cfg *Config) error {
		cfg.address = address
		return nil
	}
}

// WithMaxIdle option applies provided maximum number of idle redis connections in the pool. Default: 3.
func WithMaxIdle(maxIdle int) Option {
	return func(cfg *Config) error {
		cfg.maxIdle = maxIdle
		return nil
	}
}

// WithIdleTimeout option applies provided timeout in seconds after which idle connections will be dropped away.
// Default: 600.
func WithIdleTimeout(idleTimeout int) Option {
	return func(cfg *Config) error {
		cfg.idleTimeout = time.Duration(idleTimeout) * time.Second
		return nil
	}
}

// WithDealFunc option applies provided custom dial function instead of the default one.
// Used primarily in tests to mock redis connection.
func WithDealFunc(dialFunc func() (redis.Conn, error)) Option {
	return func(cfg *Config) error {
		cfg.dialFunc = dialFunc
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{
		address:     DefAddress,
		maxIdle:     DefMaxIdle,
		idleTimeout: DefIdleTimeout * time.Second,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
