package redis

import (
	"github.com/gomodule/redigo/redis"
)

// Reply structure represents some reply returned from redis server.
type Reply struct {
	err  error
	data interface{}
}

// IsError method checks if reply contains an error.
func (reply Reply) IsError() bool {
	return reply.err != nil
}

// Error method retrieves an error arisen on reply receiving.
func (reply Reply) Error() error {
	return reply.err
}

// IsNil method checks if returned data is nil.
func (reply Reply) IsNil() bool {
	return reply.data == nil
}

// String method tries to represent returned data as string.
func (reply Reply) String() (string, error) {
	result, err := redis.String(reply.data, reply.err)
	if err != nil {
		return "", redisError(err)
	}

	return result, nil
}

// Bytes method tries to represent returned data as slice of bytes.
func (reply Reply) Bytes() ([]byte, error) {
	result, err := redis.Bytes(reply.data, reply.err)
	if err != nil {
		return []byte{}, redisError(err)
	}

	return result, nil
}

// Float64 method tries to represent returned data as a number with a float point.
func (reply Reply) Float64() (float64, error) {
	result, err := redis.Float64(reply.data, reply.err)
	if err != nil {
		return 0, redisError(err)
	}

	return result, nil
}

// Int64 method tries to represent returned data as int64.
func (reply Reply) Int64() (int64, error) {
	result, err := redis.Int64(reply.data, reply.err)
	if err != nil {
		return 0, redisError(err)
	}

	return result, nil
}

// Int method tries to represent returned data as int.
func (reply Reply) Int() (int, error) {
	result, err := redis.Int(reply.data, reply.err)
	if err != nil {
		return 0, redisError(err)
	}

	return result, nil
}

// Bool method tries to represent returned data as boolean.
func (reply Reply) Bool() (bool, error) {
	result, err := redis.Bool(reply.data, reply.err)
	if err != nil {
		return false, redisError(err)
	}

	return result, nil
}

// List method tries to represent returned data as a list of strings.
func (reply Reply) List() ([]string, error) {
	result, err := redis.Strings(reply.data, reply.err)
	if err != nil {
		return nil, redisError(err)
	}

	return result, nil
}

// Map method tries to represent returned data as a map of strings.
func (reply Reply) Map() (map[string]string, error) {
	result, err := redis.StringMap(reply.data, reply.err)
	if err != nil {
		return nil, redisError(err)
	}

	return result, nil
}

// Scan method tries to put returned data into provided arguments (that must be pointers) in order.
func (reply Reply) Scan(args ...interface{}) error {
	if reply.err != nil {
		return reply.err
	}

	values, err := redis.Values(reply.data, nil)
	if err != nil {
		return redisError(err)
	}

	if len(values) == 0 {
		return ErrNoValues
	}

	_, err = redis.Scan(values, args...)
	if err != nil {
		return redisError(err)
	}

	return nil
}

// ScanStruct method tries to put returned data into a structure that is pointed to by the dest argument.
// For this method to work properly, the returned data must be a map of some sort (such as the HGETALL command returns).
//
// You can use 'redis' tag and RedisScan methods to control scanning behavior. Nil values in reply data are silently
// ignored.
func (reply Reply) ScanStruct(dest interface{}) error {
	if reply.err != nil {
		return reply.err
	}

	values, err := redis.Values(reply.data, nil)
	if err != nil {
		return redisError(err)
	}

	if len(values) == 0 {
		return ErrNoValues
	}

	if err = redis.ScanStruct(values, dest); err != nil {
		return redisError(err)
	}

	return nil
}
