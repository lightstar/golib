// Package file is a wrapper around config package and allows reading configuration from some file on disk and
// parsing it with some encoder, where file name is defined in environment variable.
//
// Used environment variables:
// CONFIG_FILE - configuration file name. Default is 'configs/config.<encoder>'.
//
// Typical usage:
//
//	cfg := config.Must(file.NewConfig(encoder))
//
//	err = cfg.Get(&myStructure)
//	if err != nil {
//	    panic(err)
//	}
//
// See config package for more details.
package file

import (
	"os"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/config/file"
)

const (
	configFileEnvVar = "CONFIG_FILE"
	defConfigFile    = "configs/config."
)

// NewConfig function creates new configuration service from file.
// Use CONFIG_FILE environment variable to define configuration file name. Default is 'configs/config.<encoder>'.
func NewConfig(encoder config.Encoder) (*config.Config, error) {
	configFile := os.Getenv(configFileEnvVar)
	if configFile == "" {
		configFile = defConfigFile + encoder.Type()
	}

	return file.NewConfig(configFile, encoder)
}
