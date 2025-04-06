// Package file is a wrapper around config package and allows reading configuration from some file on disk.
//
// Typical usage:
//
//	cfg := config.Must(file.NewConfig(fileName, encoder))
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
	"github.com/lightstar/golib/pkg/errors"
)

// NewConfig function creates new configuration service using data in some file as a source and chosen encoder.
// Most likely you will use one of the predefined encoders: json.Encoder, yaml.Encoder or toml.Encoder.
func NewConfig(name string, encoder config.Encoder) (*config.Config, error) {
	dataBytes, err := os.ReadFile(name)
	if err != nil {
		return nil, errors.NewFmt("can't read from file '%s' (%s)", name, err.Error()).WithCause(err)
	}

	return config.NewFromBytes(dataBytes, encoder)
}
