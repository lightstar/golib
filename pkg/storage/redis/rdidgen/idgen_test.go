package rdidgen_test

import (
	"testing"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/internal/test/redistest"
	"github.com/lightstar/golib/pkg/storage/redis"
	"github.com/lightstar/golib/pkg/storage/redis/rdidgen"
)

func TestIDGen(t *testing.T) {
	mockConn := redistest.NewMockConn()

	mockConn.On("Do", "INCR", []interface{}{"test:id"}).
		Return([]byte("1"), nil).
		Once()

	client, err := redis.NewClient(redis.WithDealFunc(func() (redigo.Conn, error) {
		return mockConn, nil
	}))
	require.NoError(t, err)

	idgen, err := rdidgen.New(
		rdidgen.WithKeyPrefix("test"),
		rdidgen.WithRedisClient(client),
	)
	require.NoError(t, err)

	id, err := idgen.NextID()
	require.NoError(t, err)

	require.Equal(t, int64(1), id)

	err = client.Close()
	require.NoError(t, err)

	mockConn.AssertExpectations(t)
}

func TestIDGenDefaultClient(t *testing.T) {
	require.NotPanics(t, func() {
		rdidgen.MustNew()
	})
}
