package config

import (
	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/util/fileutil"
)

// NewFromFile function creates new configuration service using data in some file as a source and chosen encoder.
// Most likely you will use one of the predefined encoders: JSONEncoder, YAMLEncoder or TOMLEncoder.
func NewFromFile(name string, encoder Encoder) (*Config, error) {
	dataBytes, err := fileutil.ReadAll(name)
	if err != nil {
		return nil, errors.New("file error").WithCause(err)
	}

	return NewFromBytes(dataBytes, encoder)
}
