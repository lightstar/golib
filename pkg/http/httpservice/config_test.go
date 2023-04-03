package httpservice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/configtest"
	"github.com/lightstar/golib/pkg/http/httpservice"
)

func TestConfig(t *testing.T) {
	var service *httpservice.Service

	require.NotPanics(t, func() {
		service = httpservice.MustNew(
			httpservice.WithName("test-service"),
			httpservice.WithDebug(true),
		)
	})

	require.Equal(t, "test-service", service.Name())
	require.True(t, service.Debug())
}

func TestConfigDefault(t *testing.T) {
	var service *httpservice.Service

	require.NotPanics(t, func() {
		service = httpservice.MustNew()
	})

	require.Equal(t, httpservice.DefName, service.Name())
	require.False(t, service.Debug())
}

func TestConfigService(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": struct {
			Name  string
			Debug bool
		}{
			Name:  "test-service",
			Debug: true,
		},
	})

	var service *httpservice.Service

	require.NotPanics(t, func() {
		service = httpservice.MustNew(httpservice.WithConfig(configService, "key"))
	})

	require.Equal(t, "test-service", service.Name())
	require.True(t, service.Debug())
}

func TestConfigServiceDefault(t *testing.T) {
	configService := configtest.New(map[string]interface{}{
		"key": configtest.ErrNoSuchKey,
	})

	var service *httpservice.Service

	require.NotPanics(t, func() {
		service = httpservice.MustNew(httpservice.WithConfig(configService, "key"))
	})

	require.Equal(t, httpservice.DefName, service.Name())
	require.False(t, service.Debug())
}

func TestConfigServiceError(t *testing.T) {
	configService := configtest.New(nil)

	require.Panics(t, func() {
		_ = httpservice.MustNew(httpservice.WithConfig(configService, "key"))
	})
}
