package redis

import "github.com/gomodule/redigo/redis"

// -------------------- Key commands --------------------

// Del method sends DEL command to the redis server.
func (multiCall MultiCall) Del(args ...interface{}) error {
	return multiCall.Send("DEL", args...)
}

// Expire method sends EXPIRE command to the redis server.
func (multiCall MultiCall) Expire(key string, seconds int64) error {
	return multiCall.Send("EXPIRE", key, seconds)
}

// ExpireAt method sends EXPIREAT command to the redis server.
func (multiCall MultiCall) ExpireAt(key string, timestamp int64) error {
	return multiCall.Send("EXPIREAT", key, timestamp)
}

// Move method sends MOVE command to the redis server.
func (multiCall MultiCall) Move(key string, db int64) error {
	return multiCall.Send("MOVE", key, db)
}

// Persist method sends PERSIST command to the redis server.
func (multiCall MultiCall) Persist(key string) error {
	return multiCall.Send("PERSIST", key)
}

// PExpire method sends PEXPIRE command to the redis server.
func (multiCall MultiCall) PExpire(key string, milliseconds int64) error {
	return multiCall.Send("PEXPIRE", key, milliseconds)
}

// PExpireAt method sends PEXPIREAT command to the redis server.
func (multiCall MultiCall) PExpireAt(key string, mtimestamp int64) error {
	return multiCall.Send("PEXPIREAT", key, mtimestamp)
}

// Rename method sends RENAME command to the redis server.
func (multiCall MultiCall) Rename(key string, newkey string) error {
	return multiCall.Send("RENAME", key, newkey)
}

// RenameNx method sends RENAMENX command to the redis server.
func (multiCall MultiCall) RenameNx(key string, newkey string) error {
	return multiCall.Send("RENAMENX", key, newkey)
}

// Restore method sends RESTORE command to the redis server.
func (multiCall MultiCall) Restore(key string, ttl int64, value interface{}) error {
	return multiCall.Send("RESTORE", key, ttl, value)
}

// -------------------- String commands --------------------

// Append method sends APPEND command to the redis server.
func (multiCall MultiCall) Append(key string, value interface{}) error {
	return multiCall.Send("APPEND", key, value)
}

// BitOp method sends BITOP command to the redis server.
func (multiCall MultiCall) BitOp(args ...interface{}) error {
	return multiCall.Send("BITOP", args...)
}

// Decr method sends DECR command to the redis server.
func (multiCall MultiCall) Decr(key string) error {
	return multiCall.Send("DECR", key)
}

// DecrBy method sends DECRBY command to the redis server.
func (multiCall MultiCall) DecrBy(key string, decrement int64) error {
	return multiCall.Send("DECRBY", key, decrement)
}

// Incr method sends INCR command to the redis server.
func (multiCall MultiCall) Incr(key string) error {
	return multiCall.Send("INCR", key)
}

// IncrBy method sends INCRBY command to the redis server.
func (multiCall MultiCall) IncrBy(key string, increment int64) error {
	return multiCall.Send("INCRBY", key, increment)
}

// IncrByFloat method sends INCRBYFLOAT command to the redis server.
func (multiCall MultiCall) IncrByFloat(key string, increment float64) error {
	return multiCall.Send("INCRBYFLOAT", key, increment)
}

// MSet method unfolds provided map and sends MSET command to the redis server.
func (multiCall MultiCall) MSet(values map[string]interface{}) error {
	return multiCall.Send("MSET", redis.Args{}.AddFlat(values)...)
}

// MSetFlat method sends MSET command to the redis server.
func (multiCall MultiCall) MSetFlat(args ...interface{}) error {
	return multiCall.Send("MSET", args...)
}

// MSetNx method unfolds provided map and sends MSETNX command to the redis server.
func (multiCall MultiCall) MSetNx(values map[string]interface{}) error {
	return multiCall.Send("MSETNX", redis.Args{}.AddFlat(values)...)
}

// MSetNxFlat method sends MSETNX command to the redis server.
func (multiCall MultiCall) MSetNxFlat(args ...interface{}) error {
	return multiCall.Send("MSETNX", args...)
}

// PSetEx method sends PSETEX command to the redis server.
func (multiCall MultiCall) PSetEx(key string, milliseconds int64, value interface{}) error {
	return multiCall.Send("PSETEX", key, milliseconds, value)
}

// Set method sends SET command to the redis server.
func (multiCall MultiCall) Set(key string, value interface{}) error {
	return multiCall.Send("SET", key, value)
}

// SetAdvanced method sends SET command with any provided advanced args to the redis server.
func (multiCall MultiCall) SetAdvanced(args ...interface{}) error {
	return multiCall.Send("SET", args...)
}

// SetBit method sends SETBIT command to the redis server.
func (multiCall MultiCall) SetBit(key string, offset int64, value interface{}) error {
	return multiCall.Send("SETBIT", key, offset, value)
}

// SetEx method sends SETEX command to the redis server.
func (multiCall MultiCall) SetEx(key string, seconds int64, value interface{}) error {
	return multiCall.Send("SETEX", key, seconds, value)
}

// SetNx method sends SETNX command to the redis server.
func (multiCall MultiCall) SetNx(key string, value interface{}) error {
	return multiCall.Send("SETNX", key, value)
}

// SetRange method sends SETRANGE command to the redis server.
func (multiCall MultiCall) SetRange(key string, offset int64, value interface{}) error {
	return multiCall.Send("SETRANGE", key, offset, value)
}

// -------------------- Hash commands --------------------

// HDel method sends HDEL command to the redis server.
func (multiCall MultiCall) HDel(args ...interface{}) error {
	return multiCall.Send("HDEL", args...)
}

// HIncrBy method sends HINCRBY command to the redis server.
func (multiCall MultiCall) HIncrBy(key string, field string, increment int64) error {
	return multiCall.Send("HINCRBY", key, field, increment)
}

// HIncrByFloat method sends HINCRBYFLOAT command to the redis server.
func (multiCall MultiCall) HIncrByFloat(key string, field string, increment float64) error {
	return multiCall.Send("HINCRBYFLOAT", key, field, increment)
}

// HMSet method unfolds provided map and sends HMSET command to the redis server.
func (multiCall MultiCall) HMSet(key string, values map[string]interface{}) error {
	return multiCall.Send("HMSET", redis.Args{}.Add(key).AddFlat(values)...)
}

// HMSetFlat method sends HMSET command to the redis server.
func (multiCall MultiCall) HMSetFlat(args ...interface{}) error {
	return multiCall.Send("HMSET", args...)
}

// HSet method sends HSET command to the redis server.
func (multiCall MultiCall) HSet(key string, field string, value interface{}) error {
	return multiCall.Send("HSET", key, field, value)
}

// HSetNx method sends HSETNX command to the redis server.
func (multiCall MultiCall) HSetNx(key string, field string, value interface{}) error {
	return multiCall.Send("HSETNX", key, field, value)
}

// -------------------- List commands --------------------

// LInsertBefore method sends LINSERT command with BEFORE modifier to the redis server.
func (multiCall MultiCall) LInsertBefore(key string, pivot interface{}, value interface{}) error {
	return multiCall.Send("LINSERT", key, "BEFORE", pivot, value)
}

// LInsertAfter method sends LINSERT command with AFTER modifier to the redis server.
func (multiCall MultiCall) LInsertAfter(key string, pivot interface{}, value interface{}) error {
	return multiCall.Send("LINSERT", key, "AFTER", pivot, value)
}

// LPush method sends LPUSH command to the redis server.
func (multiCall MultiCall) LPush(args ...interface{}) error {
	return multiCall.Send("LPUSH", args...)
}

// LPushX method sends LPUSHX command to the redis server.
func (multiCall MultiCall) LPushX(key string, value interface{}) error {
	return multiCall.Send("LPUSHX", key, value)
}

// LRem method sends LREM command to the redis server.
func (multiCall MultiCall) LRem(key string, count int64, value interface{}) error {
	return multiCall.Send("LREM", key, count, value)
}

// LSet method sends LSET command to the redis server.
func (multiCall MultiCall) LSet(key string, index int64, value interface{}) error {
	return multiCall.Send("LSET", key, index, value)
}

// LTrim method sends LTRIM command to the redis server.
func (multiCall MultiCall) LTrim(key string, start int64, stop int64) error {
	return multiCall.Send("LTRIM", key, start, stop)
}

// RPopLPush method sends RPOPLPUSH command to the redis server.
func (multiCall MultiCall) RPopLPush(source string, destination string) error {
	return multiCall.Send("RPOPLPUSH", source, destination)
}

// RPush method sends RPUSH command to the redis server.
func (multiCall MultiCall) RPush(args ...interface{}) error {
	return multiCall.Send("RPUSH", args...)
}

// RPushX method sends RPUSHX command to the redis server.
func (multiCall MultiCall) RPushX(key string, value interface{}) error {
	return multiCall.Send("RPUSHX", key, value)
}

// -------------------- Set commands --------------------

// SAdd method sends SADD command to the redis server.
func (multiCall MultiCall) SAdd(args ...interface{}) error {
	return multiCall.Send("SADD", args...)
}

// SDiffStore method sends SDIFFSTORE command to the redis server.
func (multiCall MultiCall) SDiffStore(args ...interface{}) error {
	return multiCall.Send("SDIFFSTORE", args...)
}

// SInterStore method sends SINTERSTORE command to the redis server.
func (multiCall MultiCall) SInterStore(args ...interface{}) error {
	return multiCall.Send("SINTERSTORE", args...)
}

// SMove method sends SMOVE command to the redis server.
func (multiCall MultiCall) SMove(source string, destination string, member interface{}) error {
	return multiCall.Send("SMOVE", source, destination, member)
}

// SRem method sends SREM command to the redis server.
func (multiCall MultiCall) SRem(args ...interface{}) error {
	return multiCall.Send("SREM", args...)
}

// SUnionStore method sends SUNIONSTORE command to the redis server.
func (multiCall MultiCall) SUnionStore(args ...interface{}) error {
	return multiCall.Send("SUNIONSTORE", args...)
}

// -------------------- Sorted set commands --------------------

// ZAdd method sends ZADD command to the redis server.
func (multiCall MultiCall) ZAdd(args ...interface{}) error {
	return multiCall.Send("ZADD", args...)
}

// ZIncrBy method sends ZINCRBY command to the redis server.
func (multiCall MultiCall) ZIncrBy(key string, increment int64, member interface{}) error {
	return multiCall.Send("ZINCRBY", key, increment, member)
}

// ZIncrByFloat method sends ZINCRBY command with float increment to the redis server.
func (multiCall MultiCall) ZIncrByFloat(key string, increment float64, member interface{}) error {
	return multiCall.Send("ZINCRBY", key, increment, member)
}

// ZInterStore method sends ZINTERSTORE command to the redis server.
func (multiCall MultiCall) ZInterStore(args ...interface{}) error {
	return multiCall.Send("ZINTERSTORE", args...)
}

// ZRem method sends ZREM command to the redis server.
func (multiCall MultiCall) ZRem(args ...interface{}) error {
	return multiCall.Send("ZREM", args...)
}

// ZRemRangeByRank method sends ZREMRANGEBYRANK command to the redis server.
func (multiCall MultiCall) ZRemRangeByRank(key string, start int64, stop int64) error {
	return multiCall.Send("ZREMRANGEBYRANK", key, start, stop)
}

// ZRemRangeByScore method sends ZREMRANGEBYSCORE command to the redis server.
func (multiCall MultiCall) ZRemRangeByScore(key string, min interface{}, max interface{}) error {
	return multiCall.Send("ZREMRANGEBYSCORE", key, min, max)
}

// ZUnionStore method sends ZUNIONSTORE command to the redis server.
func (multiCall MultiCall) ZUnionStore(args ...interface{}) error {
	return multiCall.Send("ZUNIONSTORE", args...)
}

// -------------------- Script commands --------------------

// Eval method sends EVAL command to the redis server.
func (multiCall MultiCall) Eval(args ...interface{}) error {
	return multiCall.Send("EVAL", args...)
}

// EvalSha method sends EVALSHA command to the redis server.
func (multiCall MultiCall) EvalSha(args ...interface{}) error {
	return multiCall.Send("EVALSHA", args...)
}

// ScriptFlush method sends SCRIPT FLUSH command to the redis server.
func (multiCall MultiCall) ScriptFlush() error {
	return multiCall.Send("SCRIPT FLUSH")
}

// ScriptKill method sends SCRIPT KILL command to the redis server.
func (multiCall MultiCall) ScriptKill() error {
	return multiCall.Send("SCRIPT KILL")
}
