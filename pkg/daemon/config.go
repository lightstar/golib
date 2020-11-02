package daemon

import (
	"time"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/log"
)

const (
	// DefDelay is the default process delay interval in milliseconds.
	DefDelay = 1
	// DefName is the default daemon's name.
	DefName = "daemon"
)

// Config structure with daemon configuration. Shouldn't be created manually.
type Config struct {
	name      string
	delay     time.Duration
	logger    log.Logger
	processor Processor
}

// Option function that is fed to New and MustNew. Obtain them using 'With' functions down below.
type Option func(*Config) error

// WithConfig option retrieves configuration from provided configuration service.
//
// Example JSON configuration with all possible fields (if some are not present, defaults will be used):
//     {
//         "name": "daemon-name",
//         "delay": 2000
//     }
func WithConfig(service config.Interface, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Name  string
			Delay int
		}{
			Name:  DefName,
			Delay: DefDelay,
		}

		err := service.GetByKey(key, &data)
		if err != nil && err != config.ErrNoSuchKey {
			return err
		}

		cfg.name = data.Name
		cfg.delay = time.Duration(data.Delay) * time.Millisecond

		return nil
	}
}

// WithName option applies provided daemon name. Default: "daemon".
func WithName(name string) Option {
	return func(cfg *Config) error {
		cfg.name = name
		return nil
	}
}

// WithDelay option applies provided process delay in milliseconds. Default: 1.
func WithDelay(delay int) Option {
	return func(cfg *Config) error {
		cfg.delay = time.Duration(delay) * time.Millisecond
		return nil
	}
}

// WithLogger option applies provided logger. Default: standard logger with name equal to daemon's one.
func WithLogger(logger log.Logger) Option {
	return func(cfg *Config) error {
		cfg.logger = logger
		return nil
	}
}

// WithProcessor option applies provided implementation of Processor interface. Default: none.
func WithProcessor(processor Processor) Option {
	return func(cfg *Config) error {
		cfg.processor = processor
		return nil
	}
}

// WithProcessFunc option applies provided function as processor. Default: none.
func WithProcessFunc(processFunc func()) Option {
	return func(cfg *Config) error {
		cfg.processor = ProcessFunc(processFunc)
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{
		name:  DefName,
		delay: DefDelay * time.Millisecond,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
