package config

import (
	"github.com/pelletier/go-toml"

	"github.com/lightstar/golib/pkg/errors"
)

// TOMLEncoder is an encoder that treats configuration data as toml.
func TOMLEncoder(in []byte, out *map[string]interface{}) error {
	err := toml.Unmarshal(in, out)
	if err != nil {
		return errors.New("toml error").WithCause(err)
	}

	return nil
}
