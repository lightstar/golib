// Package env is a wrapper around config package and allows reading configuration from some source and parsing it
// with some encoder, both defined in environment variables.
//
// Used environment variables:
// CONFIG_FILE - configuration file name. Default is 'configs/config.<encoder>'.
// CONFIG_ETCD_ENDPOINTS - etcd endpoints separated with comma. Such as '127.0.0.1:2379'.
// CONFIG_ETCD_KEY - key in etcd server where configuration data is stored.
// CONFIG_ENCODER - one of the supported encoders: json, yaml or toml. Default is 'yaml'.
//
// Typical usage:
//
//	cfg := config.Must(env.NewConfig())
//
//	err = cfg.Get(&myStructure)
//	if err != nil {
//	    panic(err)
//	}
//
// See config package for more details.
package env

import (
	"os"
	"strings"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/config/encoder/json"
	"github.com/lightstar/golib/pkg/config/encoder/toml"
	"github.com/lightstar/golib/pkg/config/encoder/yaml"
	"github.com/lightstar/golib/pkg/config/etcd"
	"github.com/lightstar/golib/pkg/config/file"
	"github.com/lightstar/golib/pkg/errors"
)

const (
	configFileEnvVar          = "CONFIG_FILE"
	configEtcdEndpointsEnvVar = "CONFIG_ETCD_ENDPOINTS"
	configEtcdKeyEnvVar       = "CONFIG_ETCD_KEY"
	configEncoderEnvVar       = "CONFIG_ENCODER"
	configEncoderNameDef      = "yaml"
	configFileDefPrefix       = "configs/config."
)

// ErrUnknownEncoder error is returned when encoder defined in environment variable CONFIG_ENCODER is unsupported.
var ErrUnknownEncoder = errors.New("unknown encoder")

// NewConfig function creates new configuration service using source and encoder defined in environment variables.
// Use CONFIG_FILE to define configuration file. Default is 'configs/config.<encoder>'.
// Use CONFIG_ETCD_ENDPOINTS and CONFIG_ETCD_KEY to define etcd deployment as a source.
// Use CONFIG_ENCODER to define one of the supported encoders: json, yaml or toml. Default is 'yaml'.
func NewConfig() (*config.Config, error) {
	var configEncoder config.Encoder

	configEncoderName := os.Getenv(configEncoderEnvVar)
	if configEncoderName == "" {
		configEncoderName = configEncoderNameDef
	}

	switch configEncoderName {
	case "json":
		configEncoder = json.Encoder
	case "yaml":
		configEncoder = yaml.Encoder
	case "toml":
		configEncoder = toml.Encoder
	default:
		return nil, ErrUnknownEncoder
	}

	configEtcdEndpoints := os.Getenv(configEtcdEndpointsEnvVar)
	configEtcdKey := os.Getenv(configEtcdKeyEnvVar)

	if configEtcdEndpoints != "" && configEtcdKey != "" {
		return etcd.NewConfig(strings.Split(configEtcdEndpoints, ","), configEtcdKey, configEncoder)
	}

	configFile := os.Getenv(configFileEnvVar)
	if configFile == "" {
		configFile = configFileDefPrefix + configEncoderName
	}

	return file.NewConfig(configFile, configEncoder)
}
