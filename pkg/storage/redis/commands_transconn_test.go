package redis_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/storage/redis"
)

//nolint:funlen // tests slice is too long to pass that linter
func TestCommandsTransConn(t *testing.T) {
	type CommandTest struct {
		SubName string
		Command string
		Args    []interface{}
		Func    func(conn *redis.TransConn) error
	}

	tests := []CommandTest{
		{
			Command: "DEL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.Del("test1", "test2")
			},
		},
		{
			Command: "EXPIRE",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.TransConn) error {
				return conn.Expire("test", 1800)
			},
		},
		{
			Command: "EXPIREAT",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.TransConn) error {
				return conn.ExpireAt("test", 1800)
			},
		},
		{
			Command: "MOVE",
			Args:    []interface{}{"test", int64(1)},
			Func: func(conn *redis.TransConn) error {
				return conn.Move("test", 1)
			},
		},
		{
			Command: "PERSIST",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.TransConn) error {
				return conn.Persist("test")
			},
		},
		{
			Command: "PEXPIRE",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.TransConn) error {
				return conn.PExpire("test", 1800)
			},
		},
		{
			Command: "PEXPIREAT",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.TransConn) error {
				return conn.PExpireAt("test", 1800)
			},
		},
		{
			Command: "RENAME",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.Rename("test1", "test2")
			},
		},
		{
			Command: "RENAMENX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.RenameNx("test1", "test2")
			},
		},
		{
			Command: "RESTORE",
			Args:    []interface{}{"test1", int64(1800), "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.Restore("test1", 1800, "test2")
			},
		},
		{
			Command: "APPEND",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.Append("test1", "test2")
			},
		},
		{
			Command: "BITOP",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.BitOp("test1", "test2")
			},
		},
		{
			Command: "DECR",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.TransConn) error {
				return conn.Decr("test")
			},
		},
		{
			Command: "DECRBY",
			Args:    []interface{}{"test", int64(5)},
			Func: func(conn *redis.TransConn) error {
				return conn.DecrBy("test", 5)
			},
		},
		{
			Command: "INCR",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.TransConn) error {
				return conn.Incr("test")
			},
		},
		{
			Command: "INCRBY",
			Args:    []interface{}{"test", int64(5)},
			Func: func(conn *redis.TransConn) error {
				return conn.IncrBy("test", 5)
			},
		},
		{
			Command: "INCRBYFLOAT",
			Args:    []interface{}{"test", 5.2},
			Func: func(conn *redis.TransConn) error {
				return conn.IncrByFloat("test", 5.2)
			},
		},
		{
			Command: "MSET",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.MSet(map[string]interface{}{"test": "value"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "MSET",
			Args:    []interface{}{"test1", "value1", "test2", "value2"},
			Func: func(conn *redis.TransConn) error {
				return conn.MSetFlat("test1", "value1", "test2", "value2")
			},
		},
		{
			Command: "MSETNX",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.MSetNx(map[string]interface{}{"test": "value"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "MSETNX",
			Args:    []interface{}{"test1", "value1", "test2", "value2"},
			Func: func(conn *redis.TransConn) error {
				return conn.MSetNxFlat("test1", "value1", "test2", "value2")
			},
		},
		{
			Command: "PSETEX",
			Args:    []interface{}{"test", int64(1800), "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.PSetEx("test", 1800, "value")
			},
		},
		{
			Command: "SET",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.Set("test", "value")
			},
		},
		{
			SubName: "/ADV",
			Command: "SET",
			Args:    []interface{}{"test", "value", "EX", 60},
			Func: func(conn *redis.TransConn) error {
				return conn.SetAdvanced("test", "value", "EX", 60)
			},
		},
		{
			Command: "SETBIT",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.SetBit("test", 5, "value")
			},
		},
		{
			Command: "SETEX",
			Args:    []interface{}{"test", int64(1800), "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.SetEx("test", 1800, "value")
			},
		},
		{
			Command: "SETNX",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.SetNx("test", "value")
			},
		},
		{
			Command: "SETRANGE",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.SetRange("test", 5, "value")
			},
		},
		{
			Command: "HDEL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.HDel("test1", "test2")
			},
		},
		{
			Command: "HINCRBY",
			Args:    []interface{}{"test1", "test2", int64(5)},
			Func: func(conn *redis.TransConn) error {
				return conn.HIncrBy("test1", "test2", 5)
			},
		},
		{
			Command: "HINCRBYFLOAT",
			Args:    []interface{}{"test1", "test2", 5.2},
			Func: func(conn *redis.TransConn) error {
				return conn.HIncrByFloat("test1", "test2", 5.2)
			},
		},
		{
			Command: "HMSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.TransConn) error {
				return conn.HMSet("test", map[string]interface{}{"test1": "value1"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "HMSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.TransConn) error {
				return conn.HMSetFlat("test", "test1", "value1")
			},
		},
		{
			Command: "HSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.TransConn) error {
				return conn.HSet("test", "test1", "value1")
			},
		},
		{
			Command: "HSETNX",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.TransConn) error {
				return conn.HSetNx("test", "test1", "value1")
			},
		},
		{
			SubName: "/BEFORE",
			Command: "LINSERT",
			Args:    []interface{}{"test1", "BEFORE", "test2", "test3"},
			Func: func(conn *redis.TransConn) error {
				return conn.LInsertBefore("test1", "test2", "test3")
			},
		},
		{
			SubName: "/AFTER",
			Command: "LINSERT",
			Args:    []interface{}{"test1", "AFTER", "test2", "test3"},
			Func: func(conn *redis.TransConn) error {
				return conn.LInsertAfter("test1", "test2", "test3")
			},
		},
		{
			Command: "LPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.LPush("test1", "test2")
			},
		},
		{
			Command: "LPUSHX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.LPushX("test1", "test2")
			},
		},
		{
			Command: "LREM",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.LRem("test", 5, "value")
			},
		},
		{
			Command: "LSET",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.TransConn) error {
				return conn.LSet("test", 5, "value")
			},
		},
		{
			Command: "LTRIM",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.TransConn) error {
				return conn.LTrim("test", 1, 5)
			},
		},
		{
			Command: "RPOPLPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.RPopLPush("test1", "test2")
			},
		},
		{
			Command: "RPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.RPush("test1", "test2")
			},
		},
		{
			Command: "RPUSHX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.RPushX("test1", "test2")
			},
		},
		{
			Command: "SADD",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.SAdd("test1", "test2")
			},
		},
		{
			Command: "SDIFFSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.SDiffStore("test1", "test2")
			},
		},
		{
			Command: "SINTERSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.SInterStore("test1", "test2")
			},
		},
		{
			Command: "SMOVE",
			Args:    []interface{}{"test1", "test2", "test3"},
			Func: func(conn *redis.TransConn) error {
				return conn.SMove("test1", "test2", "test3")
			},
		},
		{
			Command: "SREM",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.SRem("test1", "test2")
			},
		},
		{
			Command: "SUNIONSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.SUnionStore("test1", "test2")
			},
		},
		{
			Command: "ZADD",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.ZAdd("test1", "test2")
			},
		},
		{
			Command: "ZINCRBY",
			Args:    []interface{}{"test1", int64(5), "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.ZIncrBy("test1", 5, "test2")
			},
		},
		{
			SubName: "/FLOAT",
			Command: "ZINCRBY",
			Args:    []interface{}{"test1", 5.2, "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.ZIncrByFloat("test1", 5.2, "test2")
			},
		},
		{
			Command: "ZINTERSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.ZInterStore("test1", "test2")
			},
		},
		{
			Command: "ZREM",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.ZRem("test1", "test2")
			},
		},
		{
			Command: "ZREMRANGEBYRANK",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.TransConn) error {
				return conn.ZRemRangeByRank("test", 1, 5)
			},
		},
		{
			Command: "ZREMRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5},
			Func: func(conn *redis.TransConn) error {
				return conn.ZRemRangeByScore("test", 1, 5)
			},
		},
		{
			Command: "ZUNIONSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.ZUnionStore("test1", "test2")
			},
		},
		{
			Command: "EVAL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.Eval("test1", "test2")
			},
		},
		{
			Command: "EVALSHA",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.TransConn) error {
				return conn.EvalSha("test1", "test2")
			},
		},
		{
			Command: "SCRIPT FLUSH",
			Args:    []interface{}(nil),
			Func: func(conn *redis.TransConn) error {
				return conn.ScriptFlush()
			},
		},
		{
			Command: "SCRIPT KILL",
			Args:    []interface{}(nil),
			Func: func(conn *redis.TransConn) error {
				return conn.ScriptKill()
			},
		},
	}

	for _, test := range tests {
		func(test CommandTest) {
			t.Run(test.Command+test.SubName, func(t *testing.T) {
				helper := newTransConnHelper(t)
				defer helper.Close(t)

				helper.mockConn.On("Send", "MULTI", []interface{}(nil)).
					Return(nil).
					Once()
				helper.mockConn.On("Send", test.Command, test.Args).
					Return(nil, nil).
					Once()
				helper.mockConn.On("Do", "EXEC", []interface{}(nil)).
					Return(nil, nil).
					Once()

				err := test.Func(helper.conn)
				require.NoError(t, err)

				reply := helper.conn.Exec()

				require.False(t, reply.IsError())
				require.NoError(t, reply.Error())
				require.True(t, reply.IsNil())
			})
		}(test)
	}
}
