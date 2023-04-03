// Package etcd is a wrapper around config package and allows reading configuration from some key in etcd server.
//
// Typical usage:
//
//	cfg := config.Must(etcd.NewConfig(endpoints, key, encoder))
//
//	err = cfg.Get(&myStructure)
//	if err != nil {
//	    panic(err)
//	}
//
// See config package for more details.
package etcd

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"

	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/errors"
)

// ErrNoData error is returned when etcd doesn't have any configuration data by given key.
var ErrNoData = errors.New("no data")

// NewConfig function creates new configuration service using data stored in some key in etcd server and
// chosen encoder.
// Most likely you will use one of the predefined encoders: json.Encoder, yaml.Encoder or toml.Encoder.
func NewConfig(endpoints []string, key string, encoder config.Encoder) (*config.Config, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: time.Second,
		LogConfig: &zap.Config{
			Level:            zap.NewAtomicLevel(),
			Encoding:         "json",
			OutputPaths:      nil,
			ErrorOutputPaths: nil,
		},
	})
	if err != nil {
		return nil, errors.NewFmt("etcd error (%s)", err.Error()).WithCause(err)
	}

	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Get(ctx, key)
	if err != nil {
		return nil, errors.NewFmt("etcd error (%s)", err.Error()).WithCause(err)
	}

	if len(resp.Kvs) == 0 {
		return nil, ErrNoData
	}

	return config.NewFromBytes(resp.Kvs[0].Value, encoder)
}
