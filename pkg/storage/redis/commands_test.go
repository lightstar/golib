package redis_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/storage/redis"
)

//nolint:funlen // tests slice is too long to pass that linter
func TestCommands(t *testing.T) {
	type CommandTest struct {
		SubName string
		Command string
		Args    []interface{}
		Func    func(conn *redis.Conn) redis.Reply
	}

	tests := []CommandTest{
		{
			Command: "DEL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Del("test1", "test2")
			},
		},
		{
			Command: "DUMP",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Dump("test")
			},
		},
		{
			Command: "EXISTS",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Exists("test")
			},
		},
		{
			Command: "EXPIRE",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Expire("test", 1800)
			},
		},
		{
			Command: "EXPIREAT",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ExpireAt("test", 1800)
			},
		},
		{
			Command: "KEYS",
			Args:    []interface{}{"test*"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Keys("test*")
			},
		},
		{
			Command: "MOVE",
			Args:    []interface{}{"test", int64(1)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Move("test", 1)
			},
		},
		{
			Command: "OBJECT",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Object("test1", "test2")
			},
		},
		{
			Command: "PERSIST",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Persist("test")
			},
		},
		{
			Command: "PEXPIRE",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.PExpire("test", 1800)
			},
		},
		{
			Command: "PEXPIREAT",
			Args:    []interface{}{"test", int64(1800)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.PExpireAt("test", 1800)
			},
		},
		{
			Command: "PTTL",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.PTTL("test")
			},
		},
		{
			Command: "RANDOMKEY",
			Args:    []interface{}(nil),
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.RandomKey()
			},
		},
		{
			Command: "RENAME",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Rename("test1", "test2")
			},
		},
		{
			Command: "RENAMENX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.RenameNx("test1", "test2")
			},
		},
		{
			Command: "RESTORE",
			Args:    []interface{}{"test1", int64(1800), "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Restore("test1", 1800, "test2")
			},
		},
		{
			Command: "SORT",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Sort("test1", "test2")
			},
		},
		{
			Command: "TTL",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.TTL("test")
			},
		},
		{
			Command: "TYPE",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Type("test")
			},
		},
		{
			Command: "APPEND",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Append("test1", "test2")
			},
		},
		{
			Command: "BITCOUNT",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.BitCount("test1", "test2")
			},
		},
		{
			Command: "BITOP",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.BitOp("test1", "test2")
			},
		},
		{
			Command: "DECR",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Decr("test")
			},
		},
		{
			Command: "DECRBY",
			Args:    []interface{}{"test", int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.DecrBy("test", 5)
			},
		},
		{
			Command: "GET",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Get("test")
			},
		},
		{
			Command: "GETBIT",
			Args:    []interface{}{"test", int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.GetBit("test", 5)
			},
		},
		{
			Command: "GETRANGE",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.GetRange("test", 1, 5)
			},
		},
		{
			Command: "GETSET",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.GetSet("test1", "test2")
			},
		},
		{
			Command: "INCR",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Incr("test")
			},
		},
		{
			Command: "INCRBY",
			Args:    []interface{}{"test", int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.IncrBy("test", 5)
			},
		},
		{
			Command: "INCRBYFLOAT",
			Args:    []interface{}{"test", 5.2},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.IncrByFloat("test", 5.2)
			},
		},
		{
			Command: "MGET",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.MGet("test1", "test2")
			},
		},
		{
			Command: "MSET",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.MSet(map[string]interface{}{"test": "value"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "MSET",
			Args:    []interface{}{"test1", "value1", "test2", "value2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.MSetFlat("test1", "value1", "test2", "value2")
			},
		},
		{
			Command: "MSETNX",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.MSetNx(map[string]interface{}{"test": "value"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "MSETNX",
			Args:    []interface{}{"test1", "value1", "test2", "value2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.MSetNxFlat("test1", "value1", "test2", "value2")
			},
		},
		{
			Command: "PSETEX",
			Args:    []interface{}{"test", int64(1800), "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.PSetEx("test", 1800, "value")
			},
		},
		{
			Command: "SET",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Set("test", "value")
			},
		},
		{
			SubName: "/ADV",
			Command: "SET",
			Args:    []interface{}{"test", "value", "EX", 60},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SetAdvanced("test", "value", "EX", 60)
			},
		},
		{
			Command: "SETBIT",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SetBit("test", 5, "value")
			},
		},
		{
			Command: "SETEX",
			Args:    []interface{}{"test", int64(1800), "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SetEx("test", 1800, "value")
			},
		},
		{
			Command: "SETNX",
			Args:    []interface{}{"test", "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SetNx("test", "value")
			},
		},
		{
			Command: "SETRANGE",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SetRange("test", 5, "value")
			},
		},
		{
			Command: "STRLEN",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.StrLen("test")
			},
		},
		{
			Command: "HDEL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HDel("test1", "test2")
			},
		},
		{
			Command: "HEXISTS",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HExists("test1", "test2")
			},
		},
		{
			Command: "HGET",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HGet("test1", "test2")
			},
		},
		{
			Command: "HGETALL",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HGetAll("test")
			},
		},
		{
			Command: "HINCRBY",
			Args:    []interface{}{"test1", "test2", int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HIncrBy("test1", "test2", 5)
			},
		},
		{
			Command: "HINCRBYFLOAT",
			Args:    []interface{}{"test1", "test2", 5.2},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HIncrByFloat("test1", "test2", 5.2)
			},
		},
		{
			Command: "HKEYS",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HKeys("test")
			},
		},
		{
			Command: "HLEN",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HLen("test")
			},
		},
		{
			Command: "HMGET",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HMGet("test1", "test2")
			},
		},
		{
			Command: "HMSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HMSet("test", map[string]interface{}{"test1": "value1"})
			},
		},
		{
			SubName: "/FLAT",
			Command: "HMSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HMSetFlat("test", "test1", "value1")
			},
		},
		{
			Command: "HSET",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HSet("test", "test1", "value1")
			},
		},
		{
			Command: "HSETNX",
			Args:    []interface{}{"test", "test1", "value1"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HSetNx("test", "test1", "value1")
			},
		},
		{
			Command: "HVALS",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.HVals("test")
			},
		},
		{
			Command: "BLPOP",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.BLPop("test1", "test2")
			},
		},
		{
			Command: "BRPOP",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.BRPop("test1", "test2")
			},
		},
		{
			Command: "BRPOPLPUSH",
			Args:    []interface{}{"test1", "test2", int64(1800)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.BRPopLPush("test1", "test2", 1800)
			},
		},
		{
			Command: "LINDEX",
			Args:    []interface{}{"test", int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LIndex("test", 5)
			},
		},
		{
			SubName: "/BEFORE",
			Command: "LINSERT",
			Args:    []interface{}{"test1", "BEFORE", "test2", "test3"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LInsertBefore("test1", "test2", "test3")
			},
		},
		{
			SubName: "/AFTER",
			Command: "LINSERT",
			Args:    []interface{}{"test1", "AFTER", "test2", "test3"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LInsertAfter("test1", "test2", "test3")
			},
		},
		{
			Command: "LLEN",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LLen("test")
			},
		},
		{
			Command: "LPOP",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LPop("test")
			},
		},
		{
			Command: "LPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LPush("test1", "test2")
			},
		},
		{
			Command: "LPUSHX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LPushX("test1", "test2")
			},
		},
		{
			Command: "LRANGE",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LRange("test", 1, 5)
			},
		},
		{
			Command: "LREM",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LRem("test", 5, "value")
			},
		},
		{
			Command: "LSET",
			Args:    []interface{}{"test", int64(5), "value"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LSet("test", 5, "value")
			},
		},
		{
			Command: "LTRIM",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.LTrim("test", 1, 5)
			},
		},
		{
			Command: "RPOP",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.RPop("test")
			},
		},
		{
			Command: "RPOPLPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.RPopLPush("test1", "test2")
			},
		},
		{
			Command: "RPUSH",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.RPush("test1", "test2")
			},
		},
		{
			Command: "RPUSHX",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.RPushX("test1", "test2")
			},
		},
		{
			Command: "SADD",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SAdd("test1", "test2")
			},
		},
		{
			Command: "SCARD",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SCard("test")
			},
		},
		{
			Command: "SDIFF",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SDiff("test1", "test2")
			},
		},
		{
			Command: "SDIFFSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SDiffStore("test1", "test2")
			},
		},
		{
			Command: "SINTER",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SInter("test1", "test2")
			},
		},
		{
			Command: "SINTERSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SInterStore("test1", "test2")
			},
		},
		{
			Command: "SISMEMBER",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SIsMember("test1", "test2")
			},
		},
		{
			Command: "SMEMBERS",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SMembers("test")
			},
		},
		{
			Command: "SMOVE",
			Args:    []interface{}{"test1", "test2", "test3"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SMove("test1", "test2", "test3")
			},
		},
		{
			Command: "SPOP",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SPop("test")
			},
		},
		{
			Command: "SRANDMEMBER",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SRandMember("test")
			},
		},
		{
			SubName: "/COUNT",
			Command: "SRANDMEMBER",
			Args:    []interface{}{"test", int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SRandMemberCount("test", 5)
			},
		},
		{
			Command: "SREM",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SRem("test1", "test2")
			},
		},
		{
			Command: "SUNION",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SUnion("test1", "test2")
			},
		},
		{
			Command: "SUNIONSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.SUnionStore("test1", "test2")
			},
		},
		{
			Command: "ZADD",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZAdd("test1", "test2")
			},
		},
		{
			Command: "ZCARD",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZCard("test")
			},
		},
		{
			Command: "ZCOUNT",
			Args:    []interface{}{"test", 1, 5},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZCount("test", 1, 5)
			},
		},
		{
			Command: "ZINCRBY",
			Args:    []interface{}{"test1", int64(5), "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZIncrBy("test1", 5, "test2")
			},
		},
		{
			SubName: "/FLOAT",
			Command: "ZINCRBY",
			Args:    []interface{}{"test1", 5.2, "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZIncrByFloat("test1", 5.2, "test2")
			},
		},
		{
			Command: "ZINTERSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZInterStore("test1", "test2")
			},
		},
		{
			Command: "ZRANGE",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRange("test", 1, 5)
			},
		},
		{
			SubName: "/WITHSCORES",
			Command: "ZRANGE",
			Args:    []interface{}{"test", int64(1), int64(5), "WITHSCORES"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRangeWithScores("test", 1, 5)
			},
		},
		{
			Command: "ZRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRangeByScore("test", 1, 5)
			},
		},
		{
			SubName: "/WITHSCORES",
			Command: "ZRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5, "WITHSCORES"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRangeByScoreWithScores("test", 1, 5)
			},
		},
		{
			SubName: "/WITHLIMIT",
			Command: "ZRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5, "LIMIT", int64(10), int64(15)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRangeByScoreWithLimit("test", 1, 5, 10, 15)
			},
		},
		{
			SubName: "/WITHSCORES/WITHLIMIT",
			Command: "ZRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5, "WITHSCORES", "LIMIT", int64(10), int64(15)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRangeByScoreWithScoresWithLimit("test", 1, 5, 10, 15)
			},
		},
		{
			Command: "ZRANK",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRank("test1", "test2")
			},
		},
		{
			Command: "ZREM",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRem("test1", "test2")
			},
		},
		{
			Command: "ZREMRANGEBYRANK",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRemRangeByRank("test", 1, 5)
			},
		},
		{
			Command: "ZREMRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRemRangeByScore("test", 1, 5)
			},
		},
		{
			Command: "ZREVRANGE",
			Args:    []interface{}{"test", int64(1), int64(5)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRevRange("test", 1, 5)
			},
		},
		{
			SubName: "/WITHSCORES",
			Command: "ZREVRANGE",
			Args:    []interface{}{"test", int64(1), int64(5), "WITHSCORES"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRevRangeWithScores("test", 1, 5)
			},
		},
		{
			Command: "ZREVRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRevRangeByScore("test", 1, 5)
			},
		},
		{
			SubName: "/WITHSCORES",
			Command: "ZREVRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5, "WITHSCORES"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRevRangeByScoreWithScores("test", 1, 5)
			},
		},
		{
			SubName: "/WITHLIMIT",
			Command: "ZREVRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5, "LIMIT", int64(10), int64(15)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRevRangeByScoreWithLimit("test", 1, 5, 10, 15)
			},
		},
		{
			SubName: "/WITHSCORES/WITHLIMIT",
			Command: "ZREVRANGEBYSCORE",
			Args:    []interface{}{"test", 1, 5, "WITHSCORES", "LIMIT", int64(10), int64(15)},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRevRangeByScoreWithScoresWithLimit("test", 1, 5, 10, 15)
			},
		},
		{
			Command: "ZREVRANK",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZRevRank("test1", "test2")
			},
		},
		{
			Command: "ZSCORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZScore("test1", "test2")
			},
		},
		{
			Command: "ZUNIONSTORE",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ZUnionStore("test1", "test2")
			},
		},
		{
			Command: "EVAL",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.Eval("test1", "test2")
			},
		},
		{
			Command: "EVALSHA",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.EvalSha("test1", "test2")
			},
		},
		{
			Command: "SCRIPT EXISTS",
			Args:    []interface{}{"test1", "test2"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ScriptExists("test1", "test2")
			},
		},
		{
			Command: "SCRIPT FLUSH",
			Args:    []interface{}(nil),
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ScriptFlush()
			},
		},
		{
			Command: "SCRIPT KILL",
			Args:    []interface{}(nil),
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ScriptKill()
			},
		},
		{
			Command: "SCRIPT LOAD",
			Args:    []interface{}{"test"},
			Func: func(conn *redis.Conn) redis.Reply {
				return conn.ScriptLoad("test")
			},
		},
	}

	for _, test := range tests {
		func(test CommandTest) {
			t.Run(test.Command+test.SubName, func(t *testing.T) {
				helper := newConnHelper(t)
				defer helper.Close(t)

				helper.mockConn.On("Do", test.Command, test.Args).
					Return(nil, nil).
					Once()

				reply := test.Func(helper.conn)

				require.False(t, reply.IsError())
				require.NoError(t, reply.Error())
			})
		}(test)
	}
}
