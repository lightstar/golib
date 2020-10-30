package config

import (
	"os"
	"strings"

	"github.com/lightstar/golib/pkg/errors"
)

const (
	configFileEnvVar          = "CONFIG_FILE"
	configEtcdEndpointsEnvVar = "CONFIG_ETCD_ENDPOINTS"
	configEtcdKeyEnvVar       = "CONFIG_ETCD_KEY"
	configEncoderEnvVar       = "CONFIG_ENCODER"
	configEncoderNameDef      = "json"
)

// NewFromEnv function creates new configuration service using source and encoder defined in environment variables.
// Use CONFIG_FILE to define configuration file.
// Use CONFIG_ETCD_ENDPOINTS and CONFIG_ETCD_KEY to define etcd deployment as a source.
// Use CONFIG_ENCODER to define one of the supported encoders: json, yaml or toml. Default is json.
func NewFromEnv() (*Config, error) {
	var configEncoder Encoder

	configEncoderName := os.Getenv(configEncoderEnvVar)
	if configEncoderName == "" {
		configEncoderName = configEncoderNameDef
	}

	switch configEncoderName {
	case "json":
		configEncoder = JSONEncoder
	case "yaml":
		configEncoder = YAMLEncoder
	case "toml":
		configEncoder = TOMLEncoder
	default:
		return nil, errors.NewFmt("unknown encoder '%s'", configEncoderName).WithCause(ErrUnknownEncoder)
	}

	configFile := os.Getenv(configFileEnvVar)
	if configFile != "" {
		return NewFromFile(configFile, configEncoder)
	}

	configEtcdEndpoints := os.Getenv(configEtcdEndpointsEnvVar)
	configEtcdKey := os.Getenv(configEtcdKeyEnvVar)

	if configEtcdEndpoints != "" && configEtcdKey != "" {
		return NewFromEtcd(strings.Split(configEtcdEndpoints, ","), configEtcdKey, configEncoder)
	}

	return nil, ErrNoSource
}
