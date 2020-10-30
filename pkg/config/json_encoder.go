package config

import (
	"encoding/json"

	"github.com/lightstar/golib/pkg/errors"
)

// JSONEncoder is an encoder that treats configuration data as json.
func JSONEncoder(in []byte, out *map[string]interface{}) error {
	err := json.Unmarshal(in, out)
	if err != nil {
		return errors.New("json error").WithCause(err)
	}

	return nil
}
