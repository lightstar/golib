package redis_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/storage/redis"
)

// nolint: funlen // tests slice is too long to pass that linter
func TestCommandsMultiCall(t *testing.T) {
	type CommandTest struct {
		SubName string
		Command string
		Args    []interface{}
		Func    redis.MultiCallFunc
	}

	tests := []CommandTest{
		{
			Command: "DEL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Del("test1", "test2")
			},
		},
		{
			Command: "EXPIRE",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Expire("test", 1800)
			},
		},
		{
			Command: "EXPIREAT",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ExpireAt("test", 1800)
			},
		},
		{
			Command: "MOVE",
			Args:    []interface{}{"test", int64(1)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Move("test", 1)
			},
		},
		{
			Command: "PERSIST",
			Args:    []interface{}{"test"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Persist("test")
			},
		},
		{
			Command: "PEXPIRE",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.PExpire("test", 1800)
			},
		},
		{
			Command: "PEXPIREAT",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.PExpireAt("test", 1800)
			},
		},
		{
			Command: "RENAME",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Rename("test1", "test2")
			},
		},
		{
			Command: "RENAMENX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.RenameNx("test1", "test2")
			},
		},
		{
			Command: "RESTORE",
			Args:    []interface{}{"test1", int64(1800), "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Restore("test1", 1800, "test2")
			},
		},
		{
			Command: "APPEND",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Append("test1", "test2")
			},
		},
		{
			Command: "BITOP",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.BitOp("test1", "test2")
			},
		},
		{
			Command: "DECR",
			Args:    []interface{}{"test"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Decr("test")
			},
		},
		{
			Command: "DECRBY",
			Args:    []interface{}{"test", int64(5)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.DecrBy("test", 5)
			},
		},
		{
			Command: "INCR",
			Args:    []interface{}{"test"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Incr("test")
			},
		},
		{
			Command: "INCRBY",
			Args:    []interface{}{"test", int64(5)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.IncrBy("test", 5)
			},
		},
		{
			Command: "INCRBYFLOAT",
			Args:    []interface{}{"test", 5.2},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.IncrByFloat("test", 5.2)
			},
		},
		{
			Command: "MSET",
			Args:    []interface{}{"test", "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.MSet(map[string]interface{}{"test": "value"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "MSET",
			Args:    []interface{}{"test1", "value1", "test2", "value2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.MSetFlat("test1", "value1", "test2", "value2")
			},
		},
		{
			Command: "MSETNX",
			Args:    []interface{}{"test", "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.MSetNx(map[string]interface{}{"test": "value"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "MSETNX",
			Args:    []interface{}{"test1", "value1", "test2", "value2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.MSetNxFlat("test1", "value1", "test2", "value2")
			},
		},
		{
			Command: "PSETEX",
			Args:    []interface{}{"test", int64(1800), "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.PSetEx("test", 1800, "value")
			},
		},
		{
			Command: "SET",
			Args:    []interface{}{"test", "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Set("test", "value")
			},
		},
		{
			SubName: "/ADV",
			Command: "SET",
			Args:    []interface{}{"test", "value", "EX", 60},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SetAdvanced("test", "value", "EX", 60)
			},
		},
		{
			Command: "SETBIT",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SetBit("test", 5, "value")
			},
		},
		{
			Command: "SETEX",
			Args:    []interface{}{"test", int64(1800), "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SetEx("test", 1800, "value")
			},
		},
		{
			Command: "SETNX",
			Args:    []interface{}{"test", "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SetNx("test", "value")
			},
		},
		{
			Command: "SETRANGE",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SetRange("test", 5, "value")
			},
		},
		{
			Command: "HDEL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.HDel("test1", "test2")
			},
		},
		{
			Command: "HINCRBY",
			Args:    []interface{}{"test1", "test2", int64(5)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.HIncrBy("test1", "test2", 5)
			},
		},
		{
			Command: "HINCRBYFLOAT",
			Args:    []interface{}{"test1", "test2", 5.2},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.HIncrByFloat("test1", "test2", 5.2)
			},
		},
		{
			Command: "HMSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.HMSet("test", map[string]interface{}{"test1": "value1"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "HMSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.HMSetFlat("test", "test1", "value1")
			},
		},
		{
			Command: "HSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.HSet("test", "test1", "value1")
			},
		},
		{
			Command: "HSETNX",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.HSetNx("test", "test1", "value1")
			},
		},
		{
			SubName: "/BEFORE",
			Command: "LINSERT",
			Args:    []interface{}{"test1", "BEFORE", "test2", "test3"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.LInsertBefore("test1", "test2", "test3")
			},
		},
		{
			SubName: "/AFTER",
			Command: "LINSERT",
			Args:    []interface{}{"test1", "AFTER", "test2", "test3"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.LInsertAfter("test1", "test2", "test3")
			},
		},
		{
			Command: "LPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.LPush("test1", "test2")
			},
		},
		{
			Command: "LPUSHX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.LPushX("test1", "test2")
			},
		},
		{
			Command: "LREM",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.LRem("test", 5, "value")
			},
		},
		{
			Command: "LSET",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.LSet("test", 5, "value")
			},
		},
		{
			Command: "LTRIM",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.LTrim("test", 1, 5)
			},
		},
		{
			Command: "RPOPLPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.RPopLPush("test1", "test2")
			},
		},
		{
			Command: "RPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.RPush("test1", "test2")
			},
		},
		{
			Command: "RPUSHX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.RPushX("test1", "test2")
			},
		},
		{
			Command: "SADD",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SAdd("test1", "test2")
			},
		},
		{
			Command: "SDIFFSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SDiffStore("test1", "test2")
			},
		},
		{
			Command: "SINTERSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SInterStore("test1", "test2")
			},
		},
		{
			Command: "SMOVE",
			Args:    []interface{}{"test1", "test2", "test3"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SMove("test1", "test2", "test3")
			},
		},
		{
			Command: "SREM",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SRem("test1", "test2")
			},
		},
		{
			Command: "SUNIONSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.SUnionStore("test1", "test2")
			},
		},
		{
			Command: "ZADD",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZAdd("test1", "test2")
			},
		},
		{
			Command: "ZINCRBY",
			Args:    []interface{}{"test1", int64(5), "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZIncrBy("test1", 5, "test2")
			},
		},
		{
			SubName: "/FLOAT",
			Command: "ZINCRBY",
			Args:    []interface{}{"test1", 5.2, "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZIncrByFloat("test1", 5.2, "test2")
			},
		},
		{
			Command: "ZINTERSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZInterStore("test1", "test2")
			},
		},
		{
			Command: "ZREM",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZRem("test1", "test2")
			},
		},
		{
			Command: "ZREMRANGEBYRANK",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZRemRangeByRank("test", 1, 5)
			},
		},
		{
			Command: "ZREMRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZRemRangeByScore("test", 1, 5)
			},
		},
		{
			Command: "ZUNIONSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ZUnionStore("test1", "test2")
			},
		},
		{
			Command: "EVAL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.Eval("test1", "test2")
			},
		},
		{
			Command: "EVALSHA",
			Args:    []interface{}{"test1", "test2"},
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.EvalSha("test1", "test2")
			},
		},
		{
			Command: "SCRIPT FLUSH",
			Args:    []interface{}(nil),
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ScriptFlush()
			},
		},
		{
			Command: "SCRIPT KILL",
			Args:    []interface{}(nil),
			Func: func(multiCall redis.MultiCall) error {
				return multiCall.ScriptKill()
			},
		},
	}

	for _, test := range tests {
		func(test CommandTest) {
			t.Run(test.Command+test.SubName, func(t *testing.T) {
				helper := newConnHelper(t)
				defer helper.Close(t)

				helper.mockConn.On("Send", test.Command, test.Args).
					Return(nil, nil).
					Once()
				helper.mockConn.On("Flush").
					Return(nil).
					Once()
				helper.mockConn.On("Receive").
					Return(nil, nil).
					Once()

				reply := helper.conn.MultiCall(test.Func)

				require.False(t, reply.IsError())
				require.NoError(t, reply.Error())
			})
		}(test)
	}
}
