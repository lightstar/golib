// Package toml provides TOML encoder to use with config package. See its documentation for more details.
package toml

import (
	"github.com/pelletier/go-toml"

	"github.com/lightstar/golib/pkg/errors"
)

type encoder struct{}

// Encoder is a TOML encoder to use with config package.
//
//nolint:gochecknoglobals // it's intended read-only exported pre-defined variable.
var Encoder = encoder{}

// Type method returns string type of the encoder.
func (encoder) Type() string {
	return "toml"
}

// Encode method treats configuration input data as TOML and parses it into provided map.
func (encoder) Encode(in []byte, out *map[string]interface{}) error {
	err := toml.Unmarshal(in, out)
	if err != nil {
		return errors.NewFmt("toml error (%s)", err.Error()).WithCause(err)
	}

	return nil
}
