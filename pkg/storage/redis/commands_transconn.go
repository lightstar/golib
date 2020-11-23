package redis

import "github.com/gomodule/redigo/redis"

// -------------------- Key commands --------------------

// Del method sends DEL command to the redis server.
func (conn *TransConn) Del(args ...interface{}) error {
	return conn.Send("DEL", args...)
}

// Expire method sends EXPIRE command to the redis server.
func (conn *TransConn) Expire(key string, seconds int64) error {
	return conn.Send("EXPIRE", key, seconds)
}

// ExpireAt method sends EXPIREAT command to the redis server.
func (conn *TransConn) ExpireAt(key string, timestamp int64) error {
	return conn.Send("EXPIREAT", key, timestamp)
}

// Move method sends MOVE command to the redis server.
func (conn *TransConn) Move(key string, db int64) error {
	return conn.Send("MOVE", key, db)
}

// Persist method sends PERSIST command to the redis server.
func (conn *TransConn) Persist(key string) error {
	return conn.Send("PERSIST", key)
}

// PExpire method sends PEXPIRE command to the redis server.
func (conn *TransConn) PExpire(key string, milliseconds int64) error {
	return conn.Send("PEXPIRE", key, milliseconds)
}

// PExpireAt method sends PEXPIREAT command to the redis server.
func (conn *TransConn) PExpireAt(key string, mtimestamp int64) error {
	return conn.Send("PEXPIREAT", key, mtimestamp)
}

// Rename method sends RENAME command to the redis server.
func (conn *TransConn) Rename(key string, newkey string) error {
	return conn.Send("RENAME", key, newkey)
}

// RenameNx method sends RENAMENX command to the redis server.
func (conn *TransConn) RenameNx(key string, newkey string) error {
	return conn.Send("RENAMENX", key, newkey)
}

// Restore method sends RESTORE command to the redis server.
func (conn *TransConn) Restore(key string, ttl int64, value interface{}) error {
	return conn.Send("RESTORE", key, ttl, value)
}

// -------------------- String commands --------------------

// Append method sends APPEND command to the redis server.
func (conn *TransConn) Append(key string, value interface{}) error {
	return conn.Send("APPEND", key, value)
}

// BitOp method sends BITOP command to the redis server.
func (conn *TransConn) BitOp(args ...interface{}) error {
	return conn.Send("BITOP", args...)
}

// Decr method sends DECR command to the redis server.
func (conn *TransConn) Decr(key string) error {
	return conn.Send("DECR", key)
}

// DecrBy method sends DECRBY command to the redis server.
func (conn *TransConn) DecrBy(key string, decrement int64) error {
	return conn.Send("DECRBY", key, decrement)
}

// Incr method sends INCR command to the redis server.
func (conn *TransConn) Incr(key string) error {
	return conn.Send("INCR", key)
}

// IncrBy method sends INCRBY command to the redis server.
func (conn *TransConn) IncrBy(key string, increment int64) error {
	return conn.Send("INCRBY", key, increment)
}

// IncrByFloat method sends INCRBYFLOAT command to the redis server.
func (conn *TransConn) IncrByFloat(key string, increment float64) error {
	return conn.Send("INCRBYFLOAT", key, increment)
}

// MSet method unfolds provided map and sends MSET command to the redis server.
func (conn *TransConn) MSet(values map[string]interface{}) error {
	return conn.Send("MSET", redis.Args{}.AddFlat(values)...)
}

// MSetFlat method sends MSET command to the redis server.
func (conn *TransConn) MSetFlat(args ...interface{}) error {
	return conn.Send("MSET", args...)
}

// MSetNx method unfolds provided map and sends MSETNX command to the redis server.
func (conn *TransConn) MSetNx(values map[string]interface{}) error {
	return conn.Send("MSETNX", redis.Args{}.AddFlat(values)...)
}

// MSetNxFlat method sends MSETNX command to the redis server.
func (conn *TransConn) MSetNxFlat(args ...interface{}) error {
	return conn.Send("MSETNX", args...)
}

// PSetEx method sends PSETEX command to the redis server.
func (conn *TransConn) PSetEx(key string, milliseconds int64, value interface{}) error {
	return conn.Send("PSETEX", key, milliseconds, value)
}

// Set method sends SET command to the redis server.
func (conn *TransConn) Set(key string, value interface{}) error {
	return conn.Send("SET", key, value)
}

// SetAdvanced method sends SET command with any provided advanced args to the redis server.
func (conn *TransConn) SetAdvanced(args ...interface{}) error {
	return conn.Send("SET", args...)
}

// SetBit method sends SETBIT command to the redis server.
func (conn *TransConn) SetBit(key string, offset int64, value interface{}) error {
	return conn.Send("SETBIT", key, offset, value)
}

// SetEx method sends SETEX command to the redis server.
func (conn *TransConn) SetEx(key string, seconds int64, value interface{}) error {
	return conn.Send("SETEX", key, seconds, value)
}

// SetNx method sends SETNX command to the redis server.
func (conn *TransConn) SetNx(key string, value interface{}) error {
	return conn.Send("SETNX", key, value)
}

// SetRange method sends SETRANGE command to the redis server.
func (conn *TransConn) SetRange(key string, offset int64, value interface{}) error {
	return conn.Send("SETRANGE", key, offset, value)
}

// -------------------- Hash commands --------------------

// HDel method sends HDEL command to the redis server.
func (conn *TransConn) HDel(args ...interface{}) error {
	return conn.Send("HDEL", args...)
}

// HIncrBy method sends HINCRBY command to the redis server.
func (conn *TransConn) HIncrBy(key string, field string, increment int64) error {
	return conn.Send("HINCRBY", key, field, increment)
}

// HIncrByFloat method sends HINCRBYFLOAT command to the redis server.
func (conn *TransConn) HIncrByFloat(key string, field string, increment float64) error {
	return conn.Send("HINCRBYFLOAT", key, field, increment)
}

// HMSet method unfolds provided map and sends HMSET command to the redis server.
func (conn *TransConn) HMSet(key string, values map[string]interface{}) error {
	return conn.Send("HMSET", redis.Args{}.Add(key).AddFlat(values)...)
}

// HMSetFlat method sends HMSET command to the redis server.
func (conn *TransConn) HMSetFlat(args ...interface{}) error {
	return conn.Send("HMSET", args...)
}

// HSet method sends HSET command to the redis server.
func (conn *TransConn) HSet(key string, field string, value interface{}) error {
	return conn.Send("HSET", key, field, value)
}

// HSetNx method sends HSETNX command to the redis server.
func (conn *TransConn) HSetNx(key string, field string, value interface{}) error {
	return conn.Send("HSETNX", key, field, value)
}

// -------------------- List commands --------------------

// LInsertBefore method sends LINSERT command with BEFORE modifier to the redis server.
func (conn *TransConn) LInsertBefore(key string, pivot interface{}, value interface{}) error {
	return conn.Send("LINSERT", key, "BEFORE", pivot, value)
}

// LInsertAfter method sends LINSERT command with AFTER modifier to the redis server.
func (conn *TransConn) LInsertAfter(key string, pivot interface{}, value interface{}) error {
	return conn.Send("LINSERT", key, "AFTER", pivot, value)
}

// LPush method sends LPUSH command to the redis server.
func (conn *TransConn) LPush(args ...interface{}) error {
	return conn.Send("LPUSH", args...)
}

// LPushX method sends LPUSHX command to the redis server.
func (conn *TransConn) LPushX(key string, value interface{}) error {
	return conn.Send("LPUSHX", key, value)
}

// LRem method sends LREM command to the redis server.
func (conn *TransConn) LRem(key string, count int64, value interface{}) error {
	return conn.Send("LREM", key, count, value)
}

// LSet method sends LSET command to the redis server.
func (conn *TransConn) LSet(key string, index int64, value interface{}) error {
	return conn.Send("LSET", key, index, value)
}

// LTrim method sends LTRIM command to the redis server.
func (conn *TransConn) LTrim(key string, start int64, stop int64) error {
	return conn.Send("LTRIM", key, start, stop)
}

// RPopLPush method sends RPOPLPUSH command to the redis server.
func (conn *TransConn) RPopLPush(source string, destination string) error {
	return conn.Send("RPOPLPUSH", source, destination)
}

// RPush method sends RPUSH command to the redis server.
func (conn *TransConn) RPush(args ...interface{}) error {
	return conn.Send("RPUSH", args...)
}

// RPushX method sends RPUSHX command to the redis server.
func (conn *TransConn) RPushX(key string, value interface{}) error {
	return conn.Send("RPUSHX", key, value)
}

// -------------------- Set commands --------------------

// SAdd method sends SADD command to the redis server.
func (conn *TransConn) SAdd(args ...interface{}) error {
	return conn.Send("SADD", args...)
}

// SDiffStore method sends SDIFFSTORE command to the redis server.
func (conn *TransConn) SDiffStore(args ...interface{}) error {
	return conn.Send("SDIFFSTORE", args...)
}

// SInterStore method sends SINTERSTORE command to the redis server.
func (conn *TransConn) SInterStore(args ...interface{}) error {
	return conn.Send("SINTERSTORE", args...)
}

// SMove method sends SMOVE command to the redis server.
func (conn *TransConn) SMove(source string, destination string, member interface{}) error {
	return conn.Send("SMOVE", source, destination, member)
}

// SRem method sends SREM command to the redis server.
func (conn *TransConn) SRem(args ...interface{}) error {
	return conn.Send("SREM", args...)
}

// SUnionStore method sends SUNIONSTORE command to the redis server.
func (conn *TransConn) SUnionStore(args ...interface{}) error {
	return conn.Send("SUNIONSTORE", args...)
}

// -------------------- Sorted set commands --------------------

// ZAdd method sends ZADD command to the redis server.
func (conn *TransConn) ZAdd(args ...interface{}) error {
	return conn.Send("ZADD", args...)
}

// ZIncrBy method sends ZINCRBY command to the redis server.
func (conn *TransConn) ZIncrBy(key string, increment int64, member interface{}) error {
	return conn.Send("ZINCRBY", key, increment, member)
}

// ZIncrByFloat method sends ZINCRBY command with float increment to the redis server.
func (conn *TransConn) ZIncrByFloat(key string, increment float64, member interface{}) error {
	return conn.Send("ZINCRBY", key, increment, member)
}

// ZInterStore method sends ZINTERSTORE command to the redis server.
func (conn *TransConn) ZInterStore(args ...interface{}) error {
	return conn.Send("ZINTERSTORE", args...)
}

// ZRem method sends ZREM command to the redis server.
func (conn *TransConn) ZRem(args ...interface{}) error {
	return conn.Send("ZREM", args...)
}

// ZRemRangeByRank method sends ZREMRANGEBYRANK command to the redis server.
func (conn *TransConn) ZRemRangeByRank(key string, start int64, stop int64) error {
	return conn.Send("ZREMRANGEBYRANK", key, start, stop)
}

// ZRemRangeByScore method sends ZREMRANGEBYSCORE command to the redis server.
func (conn *TransConn) ZRemRangeByScore(key string, min interface{}, max interface{}) error {
	return conn.Send("ZREMRANGEBYSCORE", key, min, max)
}

// ZUnionStore method sends ZUNIONSTORE command to the redis server.
func (conn *TransConn) ZUnionStore(args ...interface{}) error {
	return conn.Send("ZUNIONSTORE", args...)
}

// -------------------- Script commands --------------------

// Eval method sends EVAL command to the redis server.
func (conn *TransConn) Eval(args ...interface{}) error {
	return conn.Send("EVAL", args...)
}

// EvalSha method sends EVALSHA command to the redis server.
func (conn *TransConn) EvalSha(args ...interface{}) error {
	return conn.Send("EVALSHA", args...)
}

// ScriptFlush method sends SCRIPT FLUSH command to the redis server.
func (conn *TransConn) ScriptFlush() error {
	return conn.Send("SCRIPT FLUSH")
}

// ScriptKill method sends SCRIPT KILL command to the redis server.
func (conn *TransConn) ScriptKill() error {
	return conn.Send("SCRIPT KILL")
}
