package redis_test

import (
	"testing"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/storage/redis"
	"github.com/lightstar/golib/pkg/test/redistest"
)

type baseHelper struct {
	mockConn *redistest.MockConn
	client   *redis.Client
}

func newBaseHelper(t *testing.T) *baseHelper {
	t.Helper()

	mockConn := redistest.NewMockConn()

	client, err := redis.NewClient(redis.WithDealFunc(func() (redigo.Conn, error) {
		return mockConn, nil
	}))
	require.NoError(t, err)

	return &baseHelper{
		mockConn: mockConn,
		client:   client,
	}
}

func (helper *baseHelper) Close(t *testing.T) {
	t.Helper()

	err := helper.client.Close()
	require.NoError(t, err)

	helper.mockConn.AssertExpectations(t)
}

type connHelper struct {
	*baseHelper
	conn *redis.Conn
}

func newConnHelper(t *testing.T) *connHelper {
	t.Helper()

	helper := newBaseHelper(t)

	return &connHelper{
		baseHelper: helper,
		conn:       helper.client.Conn(),
	}
}

func (helper *connHelper) Close(t *testing.T) {
	t.Helper()

	err := helper.conn.Close()
	require.NoError(t, err)

	helper.baseHelper.Close(t)
}

type transConnHelper struct {
	*baseHelper
	conn *redis.TransConn
}

func newTransConnHelper(t *testing.T) *transConnHelper {
	t.Helper()

	helper := newBaseHelper(t)

	return &transConnHelper{
		baseHelper: helper,
		conn:       helper.client.TransConn(),
	}
}

func (helper *transConnHelper) Close(t *testing.T) {
	t.Helper()

	err := helper.conn.Close()
	require.NoError(t, err)

	helper.baseHelper.Close(t)
}
