package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/config"
	"github.com/lightstar/golib/pkg/config/encoder/json"
)

func TestInner(t *testing.T) {
	cfg, err := config.NewFromBytes(configtest.SampleConfigDataJSON, json.Encoder)
	require.NoError(t, err)

	cfgInner, err := config.NewInner("profile", cfg)
	require.NoError(t, err)

	var profile configtest.UserProfile

	require.NoError(t, cfgInner.Get(&profile))
	require.Equal(t, profile, configtest.ExpectedSampleConfig.Profile)

	var children []configtest.ChildProfile

	require.NoError(t, cfgInner.GetByKey("children", &children))
	require.Equal(t, children, configtest.ExpectedSampleConfig.Profile.Children)
}

func TestInnerErrors(t *testing.T) {
	cfg, err := config.NewFromBytes(configtest.SampleConfigDataJSON, json.Encoder)
	require.NoError(t, err)

	_, err = config.NewInner("unknown", cfg)
	require.Same(t, config.ErrNoSuchKey, err)

	_, err = config.NewInner("profile.children", cfg)
	require.Same(t, config.ErrNotMap, err)
}
