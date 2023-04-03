package rdidgen_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/storage/redis/rdidgen"
)

func TestConfig(t *testing.T) {
	var idGen *rdidgen.IDGen

	require.NotPanics(t, func() {
		idGen = rdidgen.MustNew(
			rdidgen.WithKeyPrefix("test"),
		)
	})

	require.Equal(t, "test:id", idGen.Key())
}

func TestConfigDefault(t *testing.T) {
	var idGen *rdidgen.IDGen

	require.NotPanics(t, func() {
		idGen = rdidgen.MustNew()
	})

	require.Equal(t, rdidgen.DefKeyPrefix+":id", idGen.Key())
}

func TestConfigService(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			KeyPrefix string
		}{
			KeyPrefix: "test",
		},
	})

	var idGen *rdidgen.IDGen

	require.NotPanics(t, func() {
		idGen = rdidgen.MustNew(rdidgen.WithConfig(configService, "key"))
	})

	require.Equal(t, "test:id", idGen.Key())
}

func TestConfigServiceDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": configtest.ErrNoSuchKey,
	})

	var idGen *rdidgen.IDGen

	require.NotPanics(t, func() {
		idGen = rdidgen.MustNew(rdidgen.WithConfig(configService, "key"))
	})

	require.Equal(t, rdidgen.DefKeyPrefix+":id", idGen.Key())
}

func TestConfigServiceError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = rdidgen.MustNew(rdidgen.WithConfig(configService, "key"))
	})
}
