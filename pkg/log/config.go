package log

import (
	"go.uber.org/zap/zapcore"
)

// Config structure with logger configuration. Shouldn't be created manually.
type Config struct {
	name   string
	debug  bool
	stdout zapcore.WriteSyncer
	stderr zapcore.WriteSyncer
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
//	    "name": "logger-name",
//	    "debug": true
//	}
func WithConfig(service ConfigService, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Name  string
			Debug bool
		}{}

		err := service.GetByKey(key, &data)
		if err != nil && !service.IsNoSuchKeyError(err) {
			return err
		}

		cfg.name = data.Name
		cfg.debug = data.Debug

		return nil
	}
}

// WithName option applies provided logger name.
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

// WithStdout option applies customized stdout, usually for testing.
func WithStdout(stdout zapcore.WriteSyncer) Option {
	return func(cfg *Config) error {
		cfg.stdout = stdout
		return nil
	}
}

// WithStderr option applies customized stderr, usually for testing.
func WithStderr(stderr zapcore.WriteSyncer) Option {
	return func(cfg *Config) error {
		cfg.stderr = stderr
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
