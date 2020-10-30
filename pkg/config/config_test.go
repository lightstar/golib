package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/goworld/pkg/config"
	"github.com/lightstar/goworld/pkg/config/i2s"
	"github.com/lightstar/goworld/pkg/errors"
)

func TestConfig(t *testing.T) {
	cfg, err := config.NewFromBytes(sampleConfigDataJSON, config.JSONEncoder)
	require.NoError(t, err)

	testSampleConfig(t, cfg, expectedSampleRawDataJSON)
}

func TestErrors(t *testing.T) {
	cfg, err := config.NewFromBytes(sampleConfigDataJSON, config.JSONEncoder)
	require.NoError(t, err)

	_, err = cfg.GetRawByKey("unknown")
	require.Same(t, config.ErrNoSuchKey, err)

	_, err = cfg.GetRawByKey("profile.children.unknown")
	require.Same(t, config.ErrNoSuchKey, err)

	var data sampleConfigType

	err = cfg.Get(data)
	require.Same(t, i2s.ErrOutputNotPointer, err)

	err = cfg.GetByKey("unknown", &data)
	require.Same(t, config.ErrNoSuchKey, err)

	err = cfg.GetByKey("profile", &data)
	require.True(t, errors.Is(err, i2s.ErrUnknownField))
}

func TestMust(t *testing.T) {
	require.NotPanics(t, func() {
		_ = config.Must(config.NewFromBytes(sampleConfigDataJSON, config.JSONEncoder))
	})

	require.Panics(t, func() {
		_ = config.Must(config.NewFromBytes(sampleConfigDataWrongJSON, config.JSONEncoder))
	})
}

func testSampleConfig(t *testing.T, cfg *config.Config, expectedRawData map[string]interface{}) {
	require.Equal(t, expectedRawData, cfg.GetRaw())

	rawDataByEmptyKey, err := cfg.GetRawByKey("")
	require.NoError(t, err)
	require.Equal(t, interface{}(expectedRawData), rawDataByEmptyKey)

	rawChildren, err := cfg.GetRawByKey("profile.children")
	require.NoError(t, err)
	require.Equal(t, expectedRawData["profile"].(map[string]interface{})["children"], rawChildren)

	rawChildren, err = cfg.GetRawByKey(".profile.children")
	require.NoError(t, err)
	require.Equal(t, expectedRawData["profile"].(map[string]interface{})["children"], rawChildren)

	var data sampleConfigType

	require.NoError(t, cfg.Get(&data))
	require.Equal(t, expectedSampleConfig, data)

	var children []childProfile

	require.NoError(t, cfg.GetByKey("profile.children", &children))
	require.Equal(t, expectedSampleConfig.Profile.Children, children)

	require.NoError(t, cfg.GetByKey(".profile.children", &children))
	require.Equal(t, expectedSampleConfig.Profile.Children, children)
}
