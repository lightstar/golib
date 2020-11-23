// Package redis provides convenient API to redis storage. It uses redigo library undercover.
//
// Typical simplified usage:
//      client := redis.MustNewClient(redis.WithAddress("127.0.0.1:6379"))
//      conn := client.Conn()
//      defer conn.Close()
//      value, err := conn.Get("key").String()
//      if err != nil {
//          ...
//      }
package redis

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

// Client structure provides redis client functionality. Don't create manually, use the functions down below
// instead.
type Client struct {
	address string
	pool    *redis.Pool
}

// NewClient function creates new redis client with provided options.
func NewClient(opts ...Option) (*Client, error) {
	config, err := buildConfig(opts)
	if err != nil {
		return nil, err
	}

	dialFunc := config.dialFunc
	if dialFunc == nil {
		dialFunc = func() (redis.Conn, error) {
			return redis.Dial("tcp", config.address)
		}
	}

	testOnBorrow := func(conn redis.Conn, t time.Time) error {
		_, err := conn.Do("PING")
		return err
	}

	pool := &redis.Pool{
		MaxIdle:      config.maxIdle,
		IdleTimeout:  config.idleTimeout,
		Dial:         dialFunc,
		TestOnBorrow: testOnBorrow,
	}

	return &Client{
		address: config.address,
		pool:    pool,
	}, nil
}

// MustNewClient function creates new redis client with provided options and panics on any error.
func MustNewClient(opts ...Option) *Client {
	client, err := NewClient(opts...)
	if err != nil {
		panic(err)
	}

	return client
}

// Address method retrieves redis server address.
func (client *Client) Address() string {
	return client.address
}

// MaxIdle method retrieves maximum number of idle redis connections in the pool.
func (client *Client) MaxIdle() int {
	return client.pool.MaxIdle
}

// IdleTimeout method retrieves timeout in seconds after which idle connections will be dropped away.
func (client *Client) IdleTimeout() time.Duration {
	return client.pool.IdleTimeout
}

// Conn method retrieves idle connection from the pool (if there are none, it will be created).
func (client *Client) Conn() *Conn {
	conn := client.pool.Get()
	return &Conn{conn: conn}
}

// TransConn method retrieves idle connection from the pool (if there are none, it will be created).
// This connection will be specially wrapped to automatically work in transaction mode.
func (client *Client) TransConn() *TransConn {
	conn := client.pool.Get()
	return &TransConn{conn: conn}
}

// Close method releases all resources used by client. Don't use client object after that.
func (client *Client) Close() error {
	err := client.pool.Close()
	if err != nil {
		return redisError(err)
	}

	return nil
}
