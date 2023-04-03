// Package json provides JSON encoder to use with config package. See its documentation for more details.
package json

import (
	"encoding/json"

	"github.com/lightstar/golib/pkg/errors"
)

type encoder struct{}

// Encoder is a JSON encoder to use with config package.
//
//nolint:gochecknoglobals // it's intended read-only exported pre-defined variable.
var Encoder = encoder{}

// Type method returns string type of the encoder.
func (encoder) Type() string {
	return "json"
}

// Encode method treats configuration input data as JSON and parses it into provided map.
func (encoder) Encode(in []byte, out *map[string]interface{}) error {
	err := json.Unmarshal(in, out)
	if err != nil {
		return errors.NewFmt("json error (%s)", err.Error()).WithCause(err)
	}

	return nil
}
