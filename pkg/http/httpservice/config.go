package httpservice

import (
	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/log"
)

// DefName is the default service's name.
const DefName = "http-service"

// Config structure with service's configuration. Shouldn't be created manually.
type Config struct {
	name   string
	debug  bool
	logger log.Logger
}

// Option function that is fed to New and MustNew. Obtain them using 'With' functions down below.
type Option func(*Config) error

// WithConfig option retrieves configuration from provided configuration service.
//
// Example JSON configuration with all possible fields (if some are not present, defaults will be used):
//
//	{
//	    "name": "service-name",
//	    "debug": true
//	}
func WithConfig(service config.Interface, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Name  string
			Debug bool
		}{
			Name: DefName,
		}

		err := service.GetByKey(key, &data)
		if err != nil && !errors.Is(err, config.ErrNoSuchKey) {
			return err
		}

		cfg.name = data.Name
		cfg.debug = data.Debug

		return nil
	}
}

// WithName option applies provided service name. Default: "http-service".
func WithName(name string) Option {
	return func(cfg *Config) error {
		cfg.name = name
		return nil
	}
}

// WithDebug option sets debug mode. Default: false.
func WithDebug(debug bool) Option {
	return func(cfg *Config) error {
		cfg.debug = debug
		return nil
	}
}

// WithLogger option applies provided logger. Default: standard logger with name equal to the service's one.
func WithLogger(logger log.Logger) Option {
	return func(cfg *Config) error {
		cfg.logger = logger
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{
		name: DefName,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
