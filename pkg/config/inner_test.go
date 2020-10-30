package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/goworld/pkg/config"
)

func TestInner(t *testing.T) {
	cfg, err := config.NewFromBytes(sampleConfigDataJSON, config.JSONEncoder)
	require.NoError(t, err)

	cfgInner, err := config.NewInner("profile", cfg)
	require.NoError(t, err)

	var profile userProfile

	require.NoError(t, cfgInner.Get(&profile))
	require.Equal(t, profile, expectedSampleConfig.Profile)

	var children []childProfile

	require.NoError(t, cfgInner.GetByKey("children", &children))
	require.Equal(t, children, expectedSampleConfig.Profile.Children)
}

func TestInnerErrors(t *testing.T) {
	cfg, err := config.NewFromBytes(sampleConfigDataJSON, config.JSONEncoder)
	require.NoError(t, err)

	_, err = config.NewInner("unknown", cfg)
	require.Same(t, config.ErrNoSuchKey, err)

	_, err = config.NewInner("profile.children", cfg)
	require.Same(t, config.ErrNotMap, err)
}
