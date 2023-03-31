package redis

import (
	"github.com/gomodule/redigo/redis"
)

// -------------------- Key commands --------------------

// Del method sends DEL command to the redis server and waits for the reply.
func (conn *Conn) Del(args ...interface{}) Reply {
	return conn.Do("DEL", args...)
}

// Dump method sends DUMP command to the redis server and waits for the reply.
func (conn *Conn) Dump(key string) Reply {
	return conn.Do("DUMP", key)
}

// Exists method sends EXISTS command to the redis server and waits for the reply.
func (conn *Conn) Exists(key string) Reply {
	return conn.Do("EXISTS", key)
}

// Expire method sends EXPIRE command to the redis server and waits for the reply.
func (conn *Conn) Expire(key string, seconds int64) Reply {
	return conn.Do("EXPIRE", key, seconds)
}

// ExpireAt method sends EXPIREAT command to the redis server and waits for the reply.
func (conn *Conn) ExpireAt(key string, timestamp int64) Reply {
	return conn.Do("EXPIREAT", key, timestamp)
}

// Keys method sends KEYS command to the redis server and waits for the reply.
func (conn *Conn) Keys(pattern string) Reply {
	return conn.Do("KEYS", pattern)
}

// Move method sends MOVE command to the redis server and waits for the reply.
func (conn *Conn) Move(key string, db int64) Reply {
	return conn.Do("MOVE", key, db)
}

// Object method sends OBJECT command to the redis server and waits for the reply.
func (conn *Conn) Object(args ...interface{}) Reply {
	return conn.Do("OBJECT", args...)
}

// Persist method sends PERSIST command to the redis server and waits for the reply.
func (conn *Conn) Persist(key string) Reply {
	return conn.Do("PERSIST", key)
}

// PExpire method sends PEXPIRE command to the redis server and waits for the reply.
func (conn *Conn) PExpire(key string, milliseconds int64) Reply {
	return conn.Do("PEXPIRE", key, milliseconds)
}

// PExpireAt method sends PEXPIREAT command to the redis server and waits for the reply.
func (conn *Conn) PExpireAt(key string, mtimestamp int64) Reply {
	return conn.Do("PEXPIREAT", key, mtimestamp)
}

// PTTL method sends PTTL command to the redis server and waits for the reply.
func (conn *Conn) PTTL(key string) Reply {
	return conn.Do("PTTL", key)
}

// RandomKey method sends RANDOMKEY command to the redis server and waits for the reply.
func (conn *Conn) RandomKey() Reply {
	return conn.Do("RANDOMKEY")
}

// Rename method sends RENAME command to the redis server and waits for the reply.
func (conn *Conn) Rename(key string, newkey string) Reply {
	return conn.Do("RENAME", key, newkey)
}

// RenameNx method sends RENAMENX command to the redis server and waits for the reply.
func (conn *Conn) RenameNx(key string, newkey string) Reply {
	return conn.Do("RENAMENX", key, newkey)
}

// Restore method sends RESTORE command to the redis server and waits for the reply.
func (conn *Conn) Restore(key string, ttl int64, value interface{}) Reply {
	return conn.Do("RESTORE", key, ttl, value)
}

// Sort method sends SORT command to the redis server and waits for the reply.
func (conn *Conn) Sort(args ...interface{}) Reply {
	return conn.Do("SORT", args...)
}

// TTL method sends TTL command to the redis server and waits for the reply.
func (conn *Conn) TTL(key string) Reply {
	return conn.Do("TTL", key)
}

// Type method sends TYPE command to the redis server and waits for the reply.
func (conn *Conn) Type(key string) Reply {
	return conn.Do("TYPE", key)
}

// -------------------- String commands --------------------

// Append method sends APPEND command to the redis server and waits for the reply.
func (conn *Conn) Append(key string, value interface{}) Reply {
	return conn.Do("APPEND", key, value)
}

// BitCount method sends BITCOUNT command to the redis server and waits for the reply.
func (conn *Conn) BitCount(args ...interface{}) Reply {
	return conn.Do("BITCOUNT", args...)
}

// BitOp method sends BITOP command to the redis server and waits for the reply.
func (conn *Conn) BitOp(args ...interface{}) Reply {
	return conn.Do("BITOP", args...)
}

// Decr method sends DECR command to the redis server and waits for the reply.
func (conn *Conn) Decr(key string) Reply {
	return conn.Do("DECR", key)
}

// DecrBy method sends DECRBY command to the redis server and waits for the reply.
func (conn *Conn) DecrBy(key string, decrement int64) Reply {
	return conn.Do("DECRBY", key, decrement)
}

// Get method sends GET command to the redis server and waits for the reply.
func (conn *Conn) Get(key string) Reply {
	return conn.Do("GET", key)
}

// GetBit method sends GETBIT command to the redis server and waits for the reply.
func (conn *Conn) GetBit(key string, offset int64) Reply {
	return conn.Do("GETBIT", key, offset)
}

// GetRange method sends GETRANGE command to the redis server and waits for the reply.
func (conn *Conn) GetRange(key string, start int64, end int64) Reply {
	return conn.Do("GETRANGE", key, start, end)
}

// GetSet method sends GETSET command to the redis server and waits for the reply.
func (conn *Conn) GetSet(key string, value interface{}) Reply {
	return conn.Do("GETSET", key, value)
}

// Incr method sends INCR command to the redis server and waits for the reply.
func (conn *Conn) Incr(key string) Reply {
	return conn.Do("INCR", key)
}

// IncrBy method sends INCRBY command to the redis server and waits for the reply.
func (conn *Conn) IncrBy(key string, increment int64) Reply {
	return conn.Do("INCRBY", key, increment)
}

// IncrByFloat method sends INCRBYFLOAT command to the redis server and waits for the reply.
func (conn *Conn) IncrByFloat(key string, increment float64) Reply {
	return conn.Do("INCRBYFLOAT", key, increment)
}

// MGet method sends MGET command to the redis server and waits for the reply.
func (conn *Conn) MGet(args ...interface{}) Reply {
	return conn.Do("MGET", args...)
}

// MSet method unfolds provided map, sends MSET command to the redis server and waits for the reply.
func (conn *Conn) MSet(values map[string]interface{}) Reply {
	return conn.Do("MSET", redis.Args{}.AddFlat(values)...)
}

// MSetFlat method sends MSET command to the redis server and waits for the reply.
func (conn *Conn) MSetFlat(args ...interface{}) Reply {
	return conn.Do("MSET", args...)
}

// MSetNx method unfolds provided map, sends MSETNX command to the redis server and waits for the reply.
func (conn *Conn) MSetNx(values map[string]interface{}) Reply {
	return conn.Do("MSETNX", redis.Args{}.AddFlat(values)...)
}

// MSetNxFlat method sends MSETNX command to the redis server and waits for the reply.
func (conn *Conn) MSetNxFlat(args ...interface{}) Reply {
	return conn.Do("MSETNX", args...)
}

// PSetEx method sends PSETEX command to the redis server and waits for the reply.
func (conn *Conn) PSetEx(key string, milliseconds int64, value interface{}) Reply {
	return conn.Do("PSETEX", key, milliseconds, value)
}

// Set method sends SET command to the redis server and waits for the reply.
func (conn *Conn) Set(key string, value interface{}) Reply {
	return conn.Do("SET", key, value)
}

// SetAdvanced method sends SET command with any provided advanced args to the redis server and waits for the reply.
func (conn *Conn) SetAdvanced(args ...interface{}) Reply {
	return conn.Do("SET", args...)
}

// SetBit method sends SETBIT command to the redis server and waits for the reply.
func (conn *Conn) SetBit(key string, offset int64, value interface{}) Reply {
	return conn.Do("SETBIT", key, offset, value)
}

// SetEx method sends SETEX command to the redis server and waits for the reply.
func (conn *Conn) SetEx(key string, seconds int64, value interface{}) Reply {
	return conn.Do("SETEX", key, seconds, value)
}

// SetNx method sends SETNX command to the redis server and waits for the reply.
func (conn *Conn) SetNx(key string, value interface{}) Reply {
	return conn.Do("SETNX", key, value)
}

// SetRange method sends SETRANGE command to the redis server and waits for the reply.
func (conn *Conn) SetRange(key string, offset int64, value interface{}) Reply {
	return conn.Do("SETRANGE", key, offset, value)
}

// StrLen method sends STRLEN command to the redis server and waits for the reply.
func (conn *Conn) StrLen(key string) Reply {
	return conn.Do("STRLEN", key)
}

// -------------------- Hash commands --------------------

// HDel method sends HDEL command to the redis server and waits for the reply.
func (conn *Conn) HDel(args ...interface{}) Reply {
	return conn.Do("HDEL", args...)
}

// HExists method sends HEXISTS command to the redis server and waits for the reply.
func (conn *Conn) HExists(key string, field string) Reply {
	return conn.Do("HEXISTS", key, field)
}

// HGet method sends HGET command to the redis server and waits for the reply.
func (conn *Conn) HGet(key string, field string) Reply {
	return conn.Do("HGET", key, field)
}

// HGetAll method sends HGETALL command to the redis server and waits for the reply.
func (conn *Conn) HGetAll(key string) Reply {
	return conn.Do("HGETALL", key)
}

// HIncrBy method sends HINCRBY command to the redis server and waits for the reply.
func (conn *Conn) HIncrBy(key string, field string, increment int64) Reply {
	return conn.Do("HINCRBY", key, field, increment)
}

// HIncrByFloat method sends HINCRBYFLOAT command to the redis server and waits for the reply.
func (conn *Conn) HIncrByFloat(key string, field string, increment float64) Reply {
	return conn.Do("HINCRBYFLOAT", key, field, increment)
}

// HKeys method sends HKEYS command to the redis server and waits for the reply.
func (conn *Conn) HKeys(key string) Reply {
	return conn.Do("HKEYS", key)
}

// HLen method sends HLEN command to the redis server and waits for the reply.
func (conn *Conn) HLen(key string) Reply {
	return conn.Do("HLEN", key)
}

// HMGet method sends HMGET command to the redis server and waits for the reply.
func (conn *Conn) HMGet(args ...interface{}) Reply {
	return conn.Do("HMGET", args...)
}

// HMSet method unfolds provided map, sends HMSET command to the redis server and waits for the reply.
func (conn *Conn) HMSet(key string, values map[string]interface{}) Reply {
	return conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(values)...)
}

// HMSetFlat method sends HMSET command to the redis server and waits for the reply.
func (conn *Conn) HMSetFlat(args ...interface{}) Reply {
	return conn.Do("HMSET", args...)
}

// HSet method sends HSET command to the redis server and waits for the reply.
func (conn *Conn) HSet(key string, field string, value interface{}) Reply {
	return conn.Do("HSET", key, field, value)
}

// HSetNx method sends HSETNX command to the redis server and waits for the reply.
func (conn *Conn) HSetNx(key string, field string, value interface{}) Reply {
	return conn.Do("HSETNX", key, field, value)
}

// HVals method sends HVALS command to the redis server and waits for the reply.
func (conn *Conn) HVals(key string) Reply {
	return conn.Do("HVALS", key)
}

// -------------------- List commands --------------------

// BLPop method sends BLPOP command to the redis server and waits for the reply.
func (conn *Conn) BLPop(args ...interface{}) Reply {
	return conn.Do("BLPOP", args...)
}

// BRPop method sends BRPOP command to the redis server and waits for the reply.
func (conn *Conn) BRPop(args ...interface{}) Reply {
	return conn.Do("BRPOP", args...)
}

// BRPopLPush method sends BRPOPLPUSH command to the redis server and waits for the reply.
func (conn *Conn) BRPopLPush(source string, destination string, timeout int64) Reply {
	return conn.Do("BRPOPLPUSH", source, destination, timeout)
}

// LIndex method sends LINDEX command to the redis server and waits for the reply.
func (conn *Conn) LIndex(key string, index int64) Reply {
	return conn.Do("LINDEX", key, index)
}

// LInsertBefore method sends LINSERT command with BEFORE modifier to the redis server and waits for the reply.
func (conn *Conn) LInsertBefore(key string, pivot interface{}, value interface{}) Reply {
	return conn.Do("LINSERT", key, "BEFORE", pivot, value)
}

// LInsertAfter method sends LINSERT command with AFTER modifier to the redis server and waits for the reply.
func (conn *Conn) LInsertAfter(key string, pivot interface{}, value interface{}) Reply {
	return conn.Do("LINSERT", key, "AFTER", pivot, value)
}

// LLen method sends LLEN command to the redis server and waits for the reply.
func (conn *Conn) LLen(key string) Reply {
	return conn.Do("LLEN", key)
}

// LPop method sends LPOP command to the redis server and waits for the reply.
func (conn *Conn) LPop(key string) Reply {
	return conn.Do("LPOP", key)
}

// LPush method sends LPUSH command to the redis server and waits for the reply.
func (conn *Conn) LPush(args ...interface{}) Reply {
	return conn.Do("LPUSH", args...)
}

// LPushX method sends LPUSHX command to the redis server and waits for the reply.
func (conn *Conn) LPushX(args ...interface{}) Reply {
	return conn.Do("LPUSHX", args...)
}

// LRange method sends LRANGE command to the redis server and waits for the reply.
func (conn *Conn) LRange(key string, start int64, stop int64) Reply {
	return conn.Do("LRANGE", key, start, stop)
}

// LRem method sends LREM command to the redis server and waits for the reply.
func (conn *Conn) LRem(key string, count int64, value interface{}) Reply {
	return conn.Do("LREM", key, count, value)
}

// LSet method sends LSET command to the redis server and waits for the reply.
func (conn *Conn) LSet(key string, index int64, value interface{}) Reply {
	return conn.Do("LSET", key, index, value)
}

// LTrim method sends LTRIM command to the redis server and waits for the reply.
func (conn *Conn) LTrim(key string, start int64, stop int64) Reply {
	return conn.Do("LTRIM", key, start, stop)
}

// RPop method sends RPOP command to the redis server and waits for the reply.
func (conn *Conn) RPop(key string) Reply {
	return conn.Do("RPOP", key)
}

// RPopLPush method sends RPOPLPUSH command to the redis server and waits for the reply.
func (conn *Conn) RPopLPush(source string, destination string) Reply {
	return conn.Do("RPOPLPUSH", source, destination)
}

// RPush method sends RPUSH command to the redis server and waits for the reply.
func (conn *Conn) RPush(args ...interface{}) Reply {
	return conn.Do("RPUSH", args...)
}

// RPushX method sends RPUSHX command to the redis server and waits for the reply.
func (conn *Conn) RPushX(args ...interface{}) Reply {
	return conn.Do("RPUSHX", args...)
}

// -------------------- Set commands --------------------

// SAdd method sends SADD command to the redis server and waits for the reply.
func (conn *Conn) SAdd(args ...interface{}) Reply {
	return conn.Do("SADD", args...)
}

// SCard method sends SCARD command to the redis server and waits for the reply.
func (conn *Conn) SCard(key string) Reply {
	return conn.Do("SCARD", key)
}

// SDiff method sends SDIFF command to the redis server and waits for the reply.
func (conn *Conn) SDiff(args ...interface{}) Reply {
	return conn.Do("SDIFF", args...)
}

// SDiffStore method sends SDIFFSTORE command to the redis server and waits for the reply.
func (conn *Conn) SDiffStore(args ...interface{}) Reply {
	return conn.Do("SDIFFSTORE", args...)
}

// SInter method sends SINTER command to the redis server and waits for the reply.
func (conn *Conn) SInter(args ...interface{}) Reply {
	return conn.Do("SINTER", args...)
}

// SInterStore method sends SINTERSTORE command to the redis server and waits for the reply.
func (conn *Conn) SInterStore(args ...interface{}) Reply {
	return conn.Do("SINTERSTORE", args...)
}

// SIsMember method sends SISMEMBER command to the redis server and waits for the reply.
func (conn *Conn) SIsMember(key string, member interface{}) Reply {
	return conn.Do("SISMEMBER", key, member)
}

// SMembers method sends SMEMBERS command to the redis server and waits for the reply.
func (conn *Conn) SMembers(key string) Reply {
	return conn.Do("SMEMBERS", key)
}

// SMove method sends SMOVE command to the redis server and waits for the reply.
func (conn *Conn) SMove(source string, destination string, member interface{}) Reply {
	return conn.Do("SMOVE", source, destination, member)
}

// SPop method sends SPOP command to the redis server and waits for the reply.
func (conn *Conn) SPop(key string) Reply {
	return conn.Do("SPOP", key)
}

// SRandMember method sends SRANDMEMBER command to the redis server and waits for the reply.
func (conn *Conn) SRandMember(key string) Reply {
	return conn.Do("SRANDMEMBER", key)
}

// SRandMemberCount method sends SRANDMEMBER command with count param to the redis server and waits for the reply.
func (conn *Conn) SRandMemberCount(key string, count int64) Reply {
	return conn.Do("SRANDMEMBER", key, count)
}

// SRem method sends SREM command to the redis server and waits for the reply.
func (conn *Conn) SRem(args ...interface{}) Reply {
	return conn.Do("SREM", args...)
}

// SUnion method sends SUNION command to the redis server and waits for the reply.
func (conn *Conn) SUnion(args ...interface{}) Reply {
	return conn.Do("SUNION", args...)
}

// SUnionStore method sends SUNIONSTORE command to the redis server and waits for the reply.
func (conn *Conn) SUnionStore(args ...interface{}) Reply {
	return conn.Do("SUNIONSTORE", args...)
}

// -------------------- Sorted set commands --------------------

// ZAdd method sends ZADD command to the redis server and waits for the reply.
func (conn *Conn) ZAdd(args ...interface{}) Reply {
	return conn.Do("ZADD", args...)
}

// ZCard method sends ZCARD command to the redis server and waits for the reply.
func (conn *Conn) ZCard(key string) Reply {
	return conn.Do("ZCARD", key)
}

// ZCount method sends ZCOUNT command to the redis server and waits for the reply.
func (conn *Conn) ZCount(key string, min interface{}, max interface{}) Reply {
	return conn.Do("ZCOUNT", key, min, max)
}

// ZIncrBy method sends ZINCRBY command to the redis server and waits for the reply.
func (conn *Conn) ZIncrBy(key string, increment int64, member interface{}) Reply {
	return conn.Do("ZINCRBY", key, increment, member)
}

// ZIncrByFloat method sends ZINCRBY command with float increment to the redis server and waits for the reply.
func (conn *Conn) ZIncrByFloat(key string, increment float64, member interface{}) Reply {
	return conn.Do("ZINCRBY", key, increment, member)
}

// ZInterStore method sends ZINTERSTORE command to the redis server and waits for the reply.
func (conn *Conn) ZInterStore(args ...interface{}) Reply {
	return conn.Do("ZINTERSTORE", args...)
}

// ZRange method sends ZRANGE command to the redis server and waits for the reply.
func (conn *Conn) ZRange(key string, start int64, stop int64) Reply {
	return conn.Do("ZRANGE", key, start, stop)
}

// ZRangeWithScores method sends ZRANGE command with WITHSCORES modifier to the redis server and waits for the reply.
func (conn *Conn) ZRangeWithScores(key string, start int64, stop int64) Reply {
	return conn.Do("ZRANGE", key, start, stop, "WITHSCORES")
}

// ZRangeByScore method sends ZRANGEBYSCORE command to the redis server and waits for the reply.
func (conn *Conn) ZRangeByScore(key string, min interface{}, max interface{}) Reply {
	return conn.Do("ZRANGEBYSCORE", key, min, max)
}

// ZRangeByScoreWithScores method sends ZRANGEBYSCORE command with WITHSCORES modifier to the redis server and waits
// for the reply.
func (conn *Conn) ZRangeByScoreWithScores(key string, min interface{}, max interface{}) Reply {
	return conn.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES")
}

// ZRangeByScoreWithLimit method sends ZRANGEBYSCORE command with LIMIT modifier to the redis server and waits for the
// reply.
func (conn *Conn) ZRangeByScoreWithLimit(
	key string, min interface{}, max interface{}, offset int64, count int64,
) Reply {
	return conn.Do("ZRANGEBYSCORE", key, min, max, "LIMIT", offset, count)
}

// ZRangeByScoreWithScoresWithLimit method sends ZRANGEBYSCORE command with WITHSCORES and LIMIT modifiers to the redis
// server and waits for the reply.
func (conn *Conn) ZRangeByScoreWithScoresWithLimit(key string, min interface{}, max interface{}, offset int64,
	count int64,
) Reply {
	return conn.Do("ZRANGEBYSCORE", key, min, max, "WITHSCORES", "LIMIT", offset, count)
}

// ZRank method sends ZRANK command to the redis server and waits for the reply.
func (conn *Conn) ZRank(key string, member interface{}) Reply {
	return conn.Do("ZRANK", key, member)
}

// ZRem method sends ZREM command to the redis server and waits for the reply.
func (conn *Conn) ZRem(args ...interface{}) Reply {
	return conn.Do("ZREM", args...)
}

// ZRemRangeByRank method sends ZREMRANGEBYRANK command to the redis server and waits for the reply.
func (conn *Conn) ZRemRangeByRank(key string, start int64, stop int64) Reply {
	return conn.Do("ZREMRANGEBYRANK", key, start, stop)
}

// ZRemRangeByScore method sends ZREMRANGEBYSCORE command to the redis server and waits for the reply.
func (conn *Conn) ZRemRangeByScore(key string, min interface{}, max interface{}) Reply {
	return conn.Do("ZREMRANGEBYSCORE", key, min, max)
}

// ZRevRange method sends ZREVRANGE command to the redis server and waits for the reply.
func (conn *Conn) ZRevRange(key string, start int64, stop int64) Reply {
	return conn.Do("ZREVRANGE", key, start, stop)
}

// ZRevRangeWithScores method sends ZREVRANGE command with WITHSCORES modifier to the redis server and waits for the
// reply.
func (conn *Conn) ZRevRangeWithScores(key string, start int64, stop int64) Reply {
	return conn.Do("ZREVRANGE", key, start, stop, "WITHSCORES")
}

// ZRevRangeByScore method sends ZREVRANGEBYSCORE command to the redis server and waits for the reply.
func (conn *Conn) ZRevRangeByScore(key string, max interface{}, min interface{}) Reply {
	return conn.Do("ZREVRANGEBYSCORE", key, max, min)
}

// ZRevRangeByScoreWithScores method sends ZREVRANGEBYSCORE command with WITHSCORES modifier to the redis server and
// waits for the reply.
func (conn *Conn) ZRevRangeByScoreWithScores(key string, max interface{}, min interface{}) Reply {
	return conn.Do("ZREVRANGEBYSCORE", key, max, min, "WITHSCORES")
}

// ZRevRangeByScoreWithLimit method sends ZREVRANGEBYSCORE command with LIMIT modifier to the redis server and waits
// for the reply.
func (conn *Conn) ZRevRangeByScoreWithLimit(
	key string, max interface{}, min interface{}, offset int64, count int64,
) Reply {
	return conn.Do("ZREVRANGEBYSCORE", key, max, min, "LIMIT", offset, count)
}

// ZRevRangeByScoreWithScoresWithLimit method sends ZREVRANGEBYSCORE command with WITHSCORES and LIMIT modifiers to the
// redis server and waits for the reply.
func (conn *Conn) ZRevRangeByScoreWithScoresWithLimit(
	key string, max interface{}, min interface{}, offset int64, count int64,
) Reply {
	return conn.Do("ZREVRANGEBYSCORE", key, max, min, "WITHSCORES", "LIMIT", offset, count)
}

// ZRevRank method sends ZREVRANK command to the redis server and waits for the reply.
func (conn *Conn) ZRevRank(key string, member interface{}) Reply {
	return conn.Do("ZREVRANK", key, member)
}

// ZScore method sends ZSCORE command to the redis server and waits for the reply.
func (conn *Conn) ZScore(key string, member interface{}) Reply {
	return conn.Do("ZSCORE", key, member)
}

// ZUnionStore method sends ZUNIONSTORE command to the redis server and waits for the reply.
func (conn *Conn) ZUnionStore(args ...interface{}) Reply {
	return conn.Do("ZUNIONSTORE", args...)
}

// -------------------- Script commands --------------------

// Eval method sends EVAL command to the redis server and waits for the reply.
func (conn *Conn) Eval(args ...interface{}) Reply {
	return conn.Do("EVAL", args...)
}

// EvalSha method sends EVALSHA command to the redis server and waits for the reply.
func (conn *Conn) EvalSha(args ...interface{}) Reply {
	return conn.Do("EVALSHA", args...)
}

// ScriptExists method sends SCRIPT EXISTS command to the redis server and waits for the reply.
func (conn *Conn) ScriptExists(args ...interface{}) Reply {
	return conn.Do("SCRIPT EXISTS", args...)
}

// ScriptFlush method sends SCRIPT FLUSH command to the redis server and waits for the reply.
func (conn *Conn) ScriptFlush() Reply {
	return conn.Do("SCRIPT FLUSH")
}

// ScriptKill method sends SCRIPT KILL command to the redis server and waits for the reply.
func (conn *Conn) ScriptKill() Reply {
	return conn.Do("SCRIPT KILL")
}

// ScriptLoad method sends SCRIPT LOAD command to the redis server and waits for the reply.
func (conn *Conn) ScriptLoad(script string) Reply {
	return conn.Do("SCRIPT LOAD", script)
}
