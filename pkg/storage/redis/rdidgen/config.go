package rdidgen

import (
	"github.com/lightstar/golib/pkg/storage/redis"
)

// DefKeyPrefix is the default prefix of the redis key that stores next id.
const DefKeyPrefix = "entity"

// Config structure with generator's configuration. Shouldn't be created manually.
type Config struct {
	redisClient *redis.Client
	keyPrefix   string
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
//	    "keyPrefix": "user"
//	}
func WithConfig(service ConfigService, key string) Option {
	return func(cfg *Config) error {
		data := struct {
			KeyPrefix string
		}{
			KeyPrefix: DefKeyPrefix,
		}

		err := service.GetByKey(key, &data)
		if err != nil && !service.IsNoSuchKeyError(err) {
			return err
		}

		cfg.keyPrefix = data.KeyPrefix

		return nil
	}
}

// WithRedisClient option applies provided instance of redis client. If not provided, the default client will be used.
func WithRedisClient(redisClient *redis.Client) Option {
	return func(cfg *Config) error {
		cfg.redisClient = redisClient
		return nil
	}
}

// WithKeyPrefix option applies provided redis key prefix that stores next id. Default: "entity".
func WithKeyPrefix(keyPrefix string) Option {
	return func(cfg *Config) error {
		cfg.keyPrefix = keyPrefix
		return nil
	}
}

// buildConfig function builds configuration using list of provided options.
func buildConfig(opts []Option) (*Config, error) {
	cfg := &Config{
		keyPrefix: DefKeyPrefix,
	}

	for _, opt := range opts {
		if err := opt(cfg); err != nil {
			return nil, err
		}
	}

	return cfg, nil
}
