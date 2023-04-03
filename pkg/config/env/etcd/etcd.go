// Package etcd is a wrapper around config package and allows reading configuration from some key in etcd server and
// parsing it with some encoder, where endpoints and key are defined in environment variables.
//
// Used environment variables:
// CONFIG_ETCD_ENDPOINTS - etcd endpoints separated with comma. Default is '127.0.0.1:2379'.
// CONFIG_ETCD_KEY - key in etcd server where configuration data is stored. Default is 'config'.
//
// Typical usage:
//
//	cfg := config.Must(etcd.NewConfig(encoder))
//
//	err = cfg.Get(&myStructure)
//	if err != nil {
//	    panic(err)
//	}
//
// See config package for more details.
package etcd

import (
	"os"
	"strings"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/config/etcd"
)

const (
	configEndpointsEnvVar = "CONFIG_ETCD_ENDPOINTS"
	configKeyEnvVar       = "CONFIG_ETCD_KEY"
	defConfigEndpoints    = "127.0.0.1:2379"
	defConfigKey          = "config"
)

// NewConfig function creates new configuration service from some key in etcd server.
// Use CONFIG_ETCD_ENDPOINTS environment variable to define etcd endpoints. Default is '127.0.0.1:2379'.
// Use CONFIG_ETCD_KEY environment variable to define key where configuration data is stored. Default is 'config'.
func NewConfig(encoder config.Encoder) (*config.Config, error) {
	configEtcdEndpoints := os.Getenv(configEndpointsEnvVar)
	if configEtcdEndpoints == "" {
		configEtcdEndpoints = defConfigEndpoints
	}

	configEtcdKey := os.Getenv(configKeyEnvVar)
	if configEtcdKey == "" {
		configEtcdKey = defConfigKey
	}

	return etcd.NewConfig(strings.Split(configEtcdEndpoints, ","), configEtcdKey, encoder)
}
