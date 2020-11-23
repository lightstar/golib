package redis

import "github.com/gomodule/redigo/redis"

// Conn structure represents connection to redis server retrieved from the pool.
type Conn struct {
	conn redis.Conn
}

// Close method returns connection to the pool. Don't use connection object after that.
func (conn *Conn) Close() error {
	err := conn.conn.Close()
	if err != nil {
		return redisError(err)
	}

	return nil
}

// Do method sends specified command with arguments to the redis server and waits for the reply.
func (conn *Conn) Do(commandName string, args ...interface{}) Reply {
	data, err := conn.conn.Do(commandName, args...)
	if err != nil {
		return Reply{err: redisError(err)}
	}

	return Reply{data: data}
}

// MultiCall method wraps calling of provided function for bulk processing. It will return unified reply of all
// commands that were sent in this function.
func (conn Conn) MultiCall(f MultiCallFunc) Reply {
	return multiCall(false, conn).process(f)
}

// Transaction method wraps calling of provided function for transaction processing, i.e. it will send MULTI command
// before and EXEC command after. Then it will return unified reply of all commands that were sent.
func (conn Conn) Transaction(f MultiCallFunc) Reply {
	return multiCall(true, conn).process(f)
}
