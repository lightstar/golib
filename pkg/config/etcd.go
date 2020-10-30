package config

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"

	"github.com/lightstar/goworld/pkg/errors"
)

// NewFromEtcd function create new configuration service using data under some key in etcd as a source and
// chosen encoder.
// Most likely you will use one of the predefined encoders: JSONEncoder, YAMLEncoder or TOMLEncoder.
func NewFromEtcd(endpoints []string, key string, encoder Encoder) (*Config, error) {
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
		return nil, errors.New("etcd error").WithCause(err)
	}

	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := client.Get(ctx, key)
	if err != nil {
		return nil, errors.New("etcd error").WithCause(err)
	}

	if len(resp.Kvs) == 0 {
		return nil, ErrNoData
	}

	return NewFromBytes(resp.Kvs[0].Value, encoder)
}
