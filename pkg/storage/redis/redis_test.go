package redis_test

import (
	"os"
	"testing"

	redigo "github.com/gomodule/redigo/redis"
	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/storage/redis"
)

func TestRedis(t *testing.T) {
	address := os.Getenv("TEST_REDIS_ADDRESS")
	if address == "" {
		t.Log("Provide 'TEST_REDIS_ADDRESS' environment variable to test connection with real redis server")
		return
	}

	client, err := redis.NewClient(redis.WithAddress(address))
	require.NoError(t, err)

	conn := client.Conn()
	reply := conn.Get("test_key")

	require.False(t, reply.IsError())
	require.NoError(t, reply.Error())
	require.True(t, reply.IsNil())

	err = conn.Close()
	require.NoError(t, err)

	err = client.Close()
	require.NoError(t, err)
}

func TestString(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "GET", []interface{}{"test"}).
		Return([]byte("test value"), nil).
		Once()

	reply := helper.conn.Get("test")

	require.False(t, reply.IsError())
	require.False(t, reply.IsNil())
	require.NoError(t, reply.Error())

	valString, err := reply.String()
	require.NoError(t, err)

	require.Equal(t, "test value", valString)

	valBytes, err := reply.Bytes()
	require.NoError(t, err)

	require.Equal(t, "test value", string(valBytes))

	_, err = reply.Float64()
	require.Error(t, err)

	_, err = reply.Int64()
	require.Error(t, err)

	_, err = reply.Int()
	require.Error(t, err)

	_, err = reply.Bool()
	require.Error(t, err)

	_, err = reply.List()
	require.Error(t, err)

	_, err = reply.Map()
	require.Error(t, err)
}

func TestInt(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "GET", []interface{}{"test"}).
		Return([]byte("4"), nil).
		Once()

	reply := helper.conn.Get("test")

	valInt64, err := reply.Int64()
	require.NoError(t, err)

	require.Equal(t, int64(4), valInt64)

	valInt, err := reply.Int()
	require.NoError(t, err)

	require.Equal(t, 4, valInt)
}

func TestFloat(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "GET", []interface{}{"test"}).
		Return([]byte("12.3"), nil).
		Once()

	reply := helper.conn.Get("test")

	valFloat, err := reply.Float64()
	require.NoError(t, err)

	require.Equal(t, 12.3, valFloat)
}

func TestBool(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "GET", []interface{}{"test"}).
		Return([]byte("true"), nil).
		Once()

	reply := helper.conn.Get("test")

	valBool, err := reply.Bool()
	require.NoError(t, err)

	require.Equal(t, true, valBool)
}

func TestList(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "LRANGE", []interface{}{"test", int64(0), int64(-1)}).
		Return([]interface{}{
			[]byte("element1"),
			[]byte("element2"),
		}, nil).
		Once()

	reply := helper.conn.LRange("test", 0, -1)

	valList, err := reply.List()
	require.NoError(t, err)

	require.Equal(t, []string{"element1", "element2"}, valList)

	_, err = reply.String()
	require.Error(t, err)

	_, err = reply.Bytes()
	require.Error(t, err)
}

func TestMap(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HGETALL", []interface{}{"test"}).
		Return([]interface{}{
			[]byte("key1"),
			[]byte("value1"),
			[]byte("key2"),
			[]byte("value2"),
		}, nil)

	reply := helper.conn.HGetAll("test")

	valMap, err := reply.Map()
	require.NoError(t, err)

	require.Equal(t, map[string]string{"key1": "value1", "key2": "value2"}, valMap)

	_, err = reply.String()
	require.Error(t, err)

	_, err = reply.Bytes()
	require.Error(t, err)
}

func TestScan(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HMGET", []interface{}{"test", "key1", "key2"}).
		Return([]interface{}{
			[]byte("value1"),
			[]byte("5"),
		}, nil).
		Twice()

	reply := helper.conn.HMGet("test", "key1", "key2")

	var val1 string
	var val2 int

	err := reply.Scan(&val1, &val2)
	require.NoError(t, err)

	require.Equal(t, "value1", val1)
	require.Equal(t, 5, val2)

	_, err = reply.String()
	require.Error(t, err)

	_, err = reply.Bytes()
	require.Error(t, err)

	var val3 bool

	reply = helper.conn.HMGet("test", "key1", "key2")
	err = reply.Scan(&val1, &val3)
	require.Error(t, err)
}

func TestScanEmpty(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HMGET", []interface{}{"test", "key1", "key2"}).
		Return([]interface{}{}, nil).
		Once()

	var val1, val2 string

	reply := helper.conn.HMGet("test", "key1", "key2")
	err := reply.Scan(&val1, &val2)
	require.Equal(t, redis.ErrNoValues, err)
}

func TestScanError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HMGET", []interface{}{"test", "key1", "key2"}).
		Return([]interface{}{}, errors.New("test error")).
		Once()

	var val1, val2 string

	reply := helper.conn.HMGet("test", "key1", "key2")
	err := reply.Scan(&val1, &val2)
	require.Error(t, err)

	require.Equal(t, "redis error (test error)", err.Error())
}

func TestScanValuesError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HMGET", []interface{}{"test", "key1", "key2"}).
		Return(nil, nil).
		Once()

	var val1, val2 string

	reply := helper.conn.HMGet("test", "key1", "key2")
	err := reply.Scan(&val1, &val2)
	require.Error(t, err)
}

func TestScanStruct(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HGETALL", []interface{}{"test"}).
		Return([]interface{}{
			[]byte("Key1"),
			[]byte("value1"),
			[]byte("Key2"),
			[]byte("5"),
		}, nil).
		Twice()

	reply := helper.conn.HGetAll("test")

	valStruct := struct {
		Key1 string
		Key2 int
	}{}

	err := reply.ScanStruct(&valStruct)
	require.NoError(t, err)

	require.Equal(t, &struct {
		Key1 string
		Key2 int
	}{"value1", 5}, &valStruct)

	_, err = reply.String()
	require.Error(t, err)

	_, err = reply.Bytes()
	require.Error(t, err)

	reply = helper.conn.HGetAll("test")

	valStruct2 := struct {
		Key1 string
		Key2 bool
	}{}

	err = reply.ScanStruct(&valStruct2)
	require.Error(t, err)
}

func TestScanStructEmpty(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HGETALL", []interface{}{"test"}).
		Return([]interface{}{}, nil).
		Once()

	valStruct := struct {
		Key1 string
		Key2 int
	}{}

	reply := helper.conn.HGetAll("test")
	err := reply.ScanStruct(&valStruct)
	require.Equal(t, redis.ErrNoValues, err)
}

func TestScanStructError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HGETALL", []interface{}{"test"}).
		Return([]interface{}{}, errors.New("test error")).
		Once()

	valStruct := struct {
		Key1 string
		Key2 int
	}{}

	reply := helper.conn.HGetAll("test")
	err := reply.ScanStruct(&valStruct)
	require.Error(t, err)

	require.Equal(t, "redis error (test error)", err.Error())
}

func TestScanStructValuesError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "HGETALL", []interface{}{"test"}).
		Return(nil, nil).
		Once()

	valStruct := struct {
		Key1 string
		Key2 int
	}{}

	reply := helper.conn.HGetAll("test")
	err := reply.ScanStruct(&valStruct)
	require.Error(t, err)
}

func TestNil(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "GET", []interface{}{"test"}).
		Return(nil, nil).
		Once()

	reply := helper.conn.Get("test")

	require.False(t, reply.IsError())
	require.True(t, reply.IsNil())
	require.NoError(t, reply.Error())
}

func TestError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "GET", []interface{}{"test"}).
		Return(nil, errors.New("test error")).
		Once()

	reply := helper.conn.Get("test")

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (test error)", reply.Error().Error())
}

func TestBurrow(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "PING", []interface{}(nil)).
		Return(nil, nil).
		Once()

	err := helper.conn.Close()
	require.NoError(t, err)

	helper.conn = helper.client.Conn()
}

func TestBurrowError(t *testing.T) {
	helper := newConnHelper(t)
	defer helper.Close(t)

	helper.mockConn.On("Do", "PING", []interface{}(nil)).
		Return(nil, errors.New("ping error")).
		Once()
	helper.mockConn.On("Close").Return(nil).Once() // One more 'Close' call

	err := helper.conn.Close()
	require.NoError(t, err)

	helper.conn = helper.client.Conn()
}

func TestDealError(t *testing.T) {
	client, err := redis.NewClient(redis.WithDealFunc(func() (redigo.Conn, error) {
		return nil, errors.New("test error")
	}))
	require.NoError(t, err)

	defer client.Close()

	conn := client.Conn()
	defer conn.Close()

	reply := conn.Get("test")

	require.True(t, reply.IsError())
	require.Error(t, reply.Error())
	require.Equal(t, "redis error (test error)", reply.Error().Error())
}
