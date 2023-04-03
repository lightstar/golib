package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/config/encoder/json"
	"github.com/lightstar/golib/pkg/config/i2s"
	"github.com/lightstar/golib/pkg/errors"
)

func TestConfig(t *testing.T) {
	cfg, err := config.NewFromBytes(configtest.SampleConfigDataJSON, json.Encoder)
	require.NoError(t, err)

	configtest.TestSampleConfig(t, cfg, configtest.ExpectedSampleRawDataJSON)
}

func TestErrors(t *testing.T) {
	cfg, err := config.NewFromBytes(configtest.SampleConfigDataJSON, json.Encoder)
	require.NoError(t, err)

	_, err = cfg.GetRawByKey("unknown")
	require.Same(t, config.ErrNoSuchKey, err)

	_, err = cfg.GetRawByKey("profile.children.unknown")
	require.Same(t, config.ErrNoSuchKey, err)

	var data configtest.SampleConfigType

	err = cfg.Get(data)
	require.Same(t, i2s.ErrOutputNotPointer, err)

	err = cfg.GetByKey("unknown", &data)
	require.Same(t, config.ErrNoSuchKey, err)

	err = cfg.GetByKey("profile", &data)
	require.True(t, errors.Is(err, i2s.ErrUnknownField))
}

func TestMust(t *testing.T) {
	require.NotPanics(t, func() {
		_ = config.Must(config.NewFromBytes(configtest.SampleConfigDataJSON, json.Encoder))
	})

	require.Panics(t, func() {
		_ = config.Must(config.NewFromBytes(configtest.SampleConfigDataWrongJSON, json.Encoder))
	})
}

func TestNoSuchKey(t *testing.T) {
	cfg, err := config.NewFromBytes(configtest.SampleConfigDataJSON, json.Encoder)
	require.NoError(t, err)

	require.Equal(t, true, cfg.IsNoSuchKeyError(config.ErrNoSuchKey))
	require.Equal(t, false, cfg.IsNoSuchKeyError(config.ErrNotMap))
}
