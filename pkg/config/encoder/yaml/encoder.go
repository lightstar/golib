// Package yaml provides YAML encoder to use with config package. See its documentation for more details.
package yaml

import (
	"gopkg.in/yaml.v3"

	"github.com/lightstar/golib/pkg/errors"
)

type encoder struct{}

// Encoder is a YAML encoder to use with config package.
//
//nolint:gochecknoglobals // it's intended read-only exported pre-defined variable.
var Encoder = encoder{}

// Type method returns string type of the encoder.
func (encoder) Type() string {
	return "yaml"
}

// Encode method treats configuration input data as YAML and parses it into provided map.
func (encoder) Encode(in []byte, out *map[string]interface{}) error {
	err := yaml.Unmarshal(in, out)
	if err != nil {
		return errors.NewFmt("yaml error (%s)", err.Error()).WithCause(err)
	}

	return nil
}
