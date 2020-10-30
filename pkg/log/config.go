package log

import (
	"go.uber.org/zap/zapcore"

	"github.com/lightstar/golib/pkg/config"
)

// Config structure with logger configuration. Shouldn't be created manually.
type Config struct {
	name   string
	debug  bool
	stdout zapcore.WriteSyncer
	stderr zapcore.WriteSyncer
}

// Option function that is fed to New and MustNew. Obtain them using 'With' functions down below.
type Option func(*Config) error

// WithConfig option retrieves configuration from provided configuration service.
func WithConfig(service config.Interface, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			Name  string
			Debug bool
		}{}

		err := service.GetByKey(key, &data)
		if err != nil && err != config.ErrNoSuchKey {
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

// WithDebug option enables debug mode.
func WithDebug() Option {
	return func(cfg *Config) error {
		cfg.debug = true
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
