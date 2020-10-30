package config

import (
	"gopkg.in/yaml.v3"

	"github.com/lightstar/goworld/pkg/errors"
)

// YAMLEncoder is an encoder that treats configuration data as yaml.
func YAMLEncoder(in []byte, out *map[string]interface{}) error {
	err := yaml.Unmarshal(in, out)
	if err != nil {
		return errors.New("yaml error").WithCause(err)
	}

	return nil
}
