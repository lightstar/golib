// Package config provides configuration service that can use arbitrary encoders to fill external structures with
// configuration data.
//
// Typical usage might be such as that:
//
//	cfg := config.Must(config.NewFromBytes(data, encoder))
//
//	err = cfg.Get(&myStructure)
//	if err != nil {
//	    panic(err)
//	}
//
//	err = cfg.GetByKey("myKey", &myStructure)
//	if err != nil {
//	    panic(err)
//	}
//
// Argument 'data' is raw configuration bytes in some format that provided encoder can parse.
// Argument 'encoder' can be one of predefined ones - json.Encoder, yaml.Encoder, toml.Encoder, or some custom one.
//
// Variable 'myStructure' is your custom structure designed to hold configuration data.
// For example if JSON configuration data is '{"myKey":{"age":20, "name":"Peter"}}', then myStructure could be of type
// struct { Age int, Name string }
//
// More likely you will use some of more specialized packages that work with specific sources of configuration data,
// like file on disk or etcd service.
package config

import (
	"strings"

	"github.com/lightstar/golib/pkg/config/i2s"
	"github.com/lightstar/golib/pkg/errors"
)

// Encoder is an object used to convert source bytes into structured representation of configuration.
type Encoder interface {
	Type() string
	Encode([]byte, *map[string]interface{}) error
}

// Config structure that provides configuration service. Don't create it manually, use the functions down below instead.
type Config struct {
	data map[string]interface{}
	i2s  *i2s.Convertor
}

// NewFromBytes function creates new configuration service using source bytes and chosen encoder.
// Most likely you will use one of the predefined encoders: json.Encoder, yaml.Encoder or toml.Encoder.
func NewFromBytes(dataBytes []byte, encoder Encoder) (*Config, error) {
	var data map[string]interface{}

	err := encoder.Encode(dataBytes, &data)
	if err != nil {
		return nil, err
	}

	return NewFromRaw(data), nil
}

// NewFromRaw function creates new configuration service using already parsed raw configuration data.
// You will rarely use this function yourself.
func NewFromRaw(data map[string]interface{}) *Config {
	return &Config{
		data: data,
		i2s:  i2s.Instance(),
	}
}

// Must function panics on any error that can rise after creating configuration service.
// Use this like that:
//
//	cfg := config.Must(config.NewFromBytes(...))
func Must(config *Config, err error) *Config {
	if err != nil {
		panic(err)
	}

	return config
}

// GetRaw method retrieves raw representation of configuration data.
// It should be used only internally or in very special cases.
func (config *Config) GetRaw() map[string]interface{} {
	return config.data
}

// GetRawByKey method retrieves raw representation of data that lies inside configuration under some key.
// You can use empty key to retrieve all data or use a composite key like 'key1.key2.key3' to retrieve some deep data.
// It should be used only internally or in very special cases.
func (config *Config) GetRawByKey(key string) (interface{}, error) {
	if key == "" {
		return config.data, nil
	}

	var value interface{}
	var ok bool

	valueElem := config.data
	keySlice := strings.Split(key, ".")

	for index, keyElem := range keySlice {
		if keyElem == "" {
			value = valueElem
		} else if value, ok = valueElem[keyElem]; !ok {
			return nil, ErrNoSuchKey
		}

		if index+1 < len(keySlice) {
			if valueAsMap, ok := value.(map[string]interface{}); ok {
				valueElem = valueAsMap
			} else {
				return nil, ErrNoSuchKey
			}
		}
	}

	return value, nil
}

// Get method fills structure that 'out' parameter points to with all configuration data.
// It will return an error if that structure doesn't have some field, or it is not of an appropriate type.
func (config *Config) Get(out interface{}) error {
	err := config.i2s.Convert(config.data, out)
	if err != nil {
		return err
	}

	return nil
}

// GetByKey method fills structure that 'out' parameter points to with configuration data lying under some key.
// Rules are the same as for Get method.
// You can use empty key to retrieve all data or use a composite key like 'key1.key2.key3' to retrieve some deep data.
func (config *Config) GetByKey(key string, out interface{}) error {
	var data interface{}
	var err error

	data, err = config.GetRawByKey(key)
	if err != nil {
		return err
	}

	err = config.i2s.Convert(data, out)
	if err != nil {
		return err
	}

	return nil
}

// IsNoSuchKeyError checks if provided error is 'NoSuchKey' one.
func (config *Config) IsNoSuchKeyError(err error) bool {
	return errors.Is(err, ErrNoSuchKey)
}
