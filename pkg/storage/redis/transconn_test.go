package redis_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/errors"
)

func TestTransConn(t *testing.T) {
	helper := newTransConnHelper(t)
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

	err := helper.conn.Set("test1", "something1")
	require.NoError(t, err)

	err = helper.conn.Set("test2", "something2")
	require.NoError(t, err)

	reply := helper.conn.Exec()

	require.False(t, reply.IsError())
	require.NoError(t, reply.Error())
	require.True(t, reply.IsNil())
}

func TestTransConnNopExec(t *testing.T) {
	helper := newTransConnHelper(t)
	defer helper.Close(t)

	reply := helper.conn.Exec()

	require.False(t, reply.IsError())
	require.NoError(t, reply.Error())
	require.True(t, reply.IsNil())
}

func TestTransConnDiscard(t *testing.T) {
	helper := newTransConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(nil).
		Once()
	helper.mockConn.On("Do", "DISCARD", []interface{}(nil)).
		Return(nil, nil).
		Once()

	err := helper.conn.Set("test", "something")
	require.NoError(t, err)

	reply := helper.conn.Discard()

	require.False(t, reply.IsError())
	require.NoError(t, reply.Error())
	require.True(t, reply.IsNil())
}

func TestTransConnNopDiscard(t *testing.T) {
	helper := newTransConnHelper(t)
	defer helper.Close(t)

	reply := helper.conn.Discard()

	require.False(t, reply.IsError())
	require.NoError(t, reply.Error())
	require.True(t, reply.IsNil())
}

func TestTransConnImplicitDiscard(t *testing.T) {
	helper := newTransConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "DISCARD", []interface{}(nil)).
		Return(nil).
		Once()

	err := helper.conn.Set("test", "something")
	require.NoError(t, err)
}

func TestTransConnMultiError(t *testing.T) {
	helper := newTransConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(errors.New("multi error")).
		Once()
	helper.mockConn.On("Send", "DISCARD", []interface{}(nil)).
		Return(nil).
		Once()

	err := helper.conn.Set("test", "something")

	require.Error(t, err)
	require.Equal(t, "redis error (multi error)", err.Error())
}

func TestTransConnSetError(t *testing.T) {
	helper := newTransConnHelper(t)
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

	err := helper.conn.Set("test", "something")

	require.Error(t, err)
	require.Equal(t, "redis error (set error)", err.Error())
}

//nolint:dupl // this test is very similar to TestTransConnDiscardError, but that's ok
func TestTransConnExecError(t *testing.T) {
	helper := newTransConnHelper(t)
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

	err := helper.conn.Set("test", "something")
	require.NoError(t, err)

	reply := helper.conn.Exec()

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (exec error)", reply.Error().Error())
}

//nolint:dupl // this test is very similar to TestTransConnExecError, but that's ok
func TestTransConnDiscardError(t *testing.T) {
	helper := newTransConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
		Return(nil).
		Once()
	helper.mockConn.On("Send", "SET", []interface{}{"test", "something"}).
		Return(nil).
		Once()
	helper.mockConn.On("Do", "DISCARD", []interface{}(nil)).
		Return(nil, errors.New("discard error")).
		Once()

	err := helper.conn.Set("test", "something")
	require.NoError(t, err)

	reply := helper.conn.Discard()

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (discard error)", reply.Error().Error())
}
