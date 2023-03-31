package redis_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/storage/redis"
)

func TestMultiCall(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "SET", []interface{}{"test1", "something1"}).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test2", "something2"}).
		Return(nil).
		Once()
	helper.mockConn.On("Flush").
		Return(nil).
		Once()
	helper.mockConn.On("Receive").
		Return(nil, nil).
		Once()

	reply := helper.conn.MultiCall(func(multiCall redis.MultiCall) error {
		if err := multiCall.Set("test1", "something1"); err != nil {
			return err
		}

		return multiCall.Set("test2", "something2")
	})

	require.False(t, reply.IsError())
	require.NoError(t, reply.Error())
	require.True(t, reply.IsNil())
}

func TestMultiCallTransaction(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test1", "something1"}).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test2", "something2"}).
		Return(nil).
		Once()
	helper.mockConn.On("Do", "EXEC", []interface{}(nil)).
		Return(nil, nil).
		Once()

	reply := helper.conn.Transaction(func(multiCall redis.MultiCall) error {
		if err := multiCall.Set("test1", "something1"); err != nil {
			return err
		}

		return multiCall.Set("test2", "something2")
	})

	require.False(t, reply.IsError())
	require.NoError(t, reply.Error())
	require.True(t, reply.IsNil())
}

func TestMultiCallSetError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(errors.New("set error")).
		Once()

	reply := helper.conn.MultiCall(func(multiCall redis.MultiCall) error {
		return multiCall.Set("test", "something")
	})

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (set error)", reply.Error().Error())
}

func TestMultiCallFlushError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(nil).
		Once()
	helper.mockConn.On("Flush").
		Return(errors.New("flush error")).
		Once()

	reply := helper.conn.MultiCall(func(multiCall redis.MultiCall) error {
		return multiCall.Set("test", "something")
	})

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (flush error)", reply.Error().Error())
}

func TestMultiCallReceiveError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(nil).
		Once()
	helper.mockConn.On("Flush").
		Return(nil).
		Once()
	helper.mockConn.On("Receive").
		Return(nil, errors.New("receive error")).
		Once()

	reply := helper.conn.MultiCall(func(multiCall redis.MultiCall) error {
		return multiCall.Set("test", "something")
	})

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (receive error)", reply.Error().Error())
}

func TestMultiCallTransactionMultiError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(errors.New("multi error")).
		Once()
	helper.mockConn.On("Send", "DISCARD", []interface{}(nil)).
		Return(nil).
		Once()

	reply := helper.conn.Transaction(func(multiCall redis.MultiCall) error {
		return multiCall.Set("test", "something")
	})

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (multi error)", reply.Error().Error())
}

func TestMultiCallTransactionSetError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(errors.New("set error")).
		Once()
	helper.mockConn.On("Send", "DISCARD", []interface{}(nil)).
		Return(nil).
		Once()

	reply := helper.conn.Transaction(func(multiCall redis.MultiCall) error {
		return multiCall.Set("test", "something")
	})

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (set error)", reply.Error().Error())
}

func TestMultiCallTransactionExecError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(nil).
		Once()
	helper.mockConn.On("Do", "EXEC", []interface{}(nil)).
		Return(nil, errors.New("exec error")).
		Once()

	reply := helper.conn.Transaction(func(multiCall redis.MultiCall) error {
		return multiCall.Set("test", "something")
	})

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (exec error)", reply.Error().Error())
}
