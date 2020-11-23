package redis

import "github.com/gomodule/redigo/redis"

const (
	transStateIdle transState = iota
	transStateActive
)

type transState uint8

// TransConn structure represents connection to redis server retrieved from the pool. It is specially wrapped
// to send all the commands as part of some transaction.
type TransConn struct {
	conn  redis.Conn
	state transState
}

// Close method returns connection to the pool. Don't use connection object after that.
func (conn *TransConn) Close() error {
	err := conn.conn.Close()
	if err != nil {
		return redisError(err)
	}

	return nil
}

// Send method sends specified command with arguments to the redis server without waiting for reply.
// If connection is not in transaction state yet, it will send MULTI command automatically.
func (conn *TransConn) Send(commandName string, args ...interface{}) error {
	if conn.state == transStateIdle {
		err := conn.conn.Send("MULTI")
		if err != nil {
			return redisError(err)
		}

		conn.state = transStateActive
	}

	err := conn.conn.Send(commandName, args...)
	if err != nil {
		return redisError(err)
	}

	return nil
}

// Exec method sends EXEC command to redis server and waits for the reply.
func (conn *TransConn) Exec() Reply {
	if conn.state == transStateIdle {
		return Reply{}
	}

	conn.state = transStateIdle

	data, err := conn.conn.Do("EXEC")
	if err != nil {
		return Reply{err: redisError(err)}
	}

	return Reply{data: data}
}

// Discard method sends DISCARD command to redis server and waits for the reply.
func (conn *TransConn) Discard() Reply {
	if conn.state == transStateIdle {
		return Reply{}
	}

	conn.state = transStateIdle

	data, err := conn.conn.Do("DISCARD")
	if err != nil {
		return Reply{err: redisError(err)}
	}

	return Reply{data: data}
}
