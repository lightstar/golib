// Package configtest is used as a helper in tests that need configuration service implementing config.Interface.
package configtest

import (
	"reflect"
)

// Config structure implementing config.Interface.
type Config struct {
	data map[string]interface{}
}

// New function creates configuration service using static predefined map of data.
func New(data map[string]interface{}) *Config {
	return &Config{data: data}
}

// Get method fills structure that 'out' parameter points to with predefined data under empty key.
// It will return an error if 'out' is not a pointer or predefined data is nil.
// If empty key contains an error that error will be returned.
func (config *Config) Get(out interface{}) error {
	return config.GetByKey("", out)
}

// GetByKey method fills structure that 'out' parameter points to with predefined data under the provided key.
// It will return an error if 'out' is not a pointer or predefined data is nil or there is no such key in that data.
// If corresponding key contains an error that error will be returned.
func (config *Config) GetByKey(key string, out interface{}) error {
	if config.data == nil {
		return ErrNoData
	}

	var in interface{}
	var ok bool

	if in, ok = config.data[key]; !ok {
		return ErrNoSuchKey
	}

	if err, ok := in.(error); ok {
		return err
	}

	outValue := reflect.ValueOf(out)
	if outValue.Kind() != reflect.Ptr {
		return ErrOutputNotPointer
	}

	outValue.Elem().Set(reflect.ValueOf(in))

	return nil
}
