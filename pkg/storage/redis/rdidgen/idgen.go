// Package rdidgen provides implementation of id generation using redis server.
//
// Typical usage:
//
//	idgen := rdidgen.MustNew(
//	    rdidgen.WithRedisClient(client),
//	    rdidgen.WithKeyPrefix("object"),
//	)
//
//	id, err := idgen.NextID()
//	if err != nil {
//	    ...
//	}
package rdidgen

import "github.com/lightstar/golib/pkg/storage/redis"

// IDGen structure provides id generation functionality using redis client. Don't create manually, use the functions
// down below instead.
type IDGen struct {
	key   string
	redis *redis.Client
}

// New function creates new redis id generator with provided options.
func New(opts ...Option) (*IDGen, error) {
	config, err := buildConfig(opts)
	if err != nil {
		return nil, err
	}

	redisClient := config.redisClient
	if redisClient == nil {
		redisClient, err = redis.NewClient()
		if err != nil {
			return nil, err
		}
	}

	return &IDGen{
		key:   config.keyPrefix + ":id",
		redis: redisClient,
	}, nil
}

// MustNew function creates new redis id generator with provided options and panics on any error.
func MustNew(opts ...Option) *IDGen {
	idGen, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return idGen
}

// Key method retrieves redis key that stores next id.
func (idGen *IDGen) Key() string {
	return idGen.key
}

// NextID method increments and retrieves the next id.
func (idGen *IDGen) NextID() (int64, error) {
	conn := idGen.redis.Conn()
	defer conn.Close()

	return conn.Incr(idGen.key).Int64()
}
