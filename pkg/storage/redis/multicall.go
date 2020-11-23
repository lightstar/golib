package redis

import "github.com/gomodule/redigo/redis"

// MultiCall structure is a bulk or transaction processing context. Do not create it manually.
// Use it inside MultiCallFunc to send commands that are part of that bulk or transaction processing.
type MultiCall struct {
	conn        redis.Conn
	transaction bool
}

// MultiCallFunc is a function that will be wrapped in bulk and transaction processing. Inside it should call all
// necessary commands to redis server using provided MultiCall context.
type MultiCallFunc func(MultiCall) error

// multiCall method creates new MultiCall structure used for bulk or transaction processing.
func multiCall(transaction bool, conn Conn) MultiCall {
	return MultiCall{
		transaction: transaction,
		conn:        conn.conn,
	}
}

// Send method sends provided command and arguments to redis server as part of bulk or transaction processing.
// Generally you will not call this method directly, but use helper command methods instead.
func (multiCall MultiCall) Send(commandName string, args ...interface{}) error {
	err := multiCall.conn.Send(commandName, args...)
	if err != nil {
		return redisError(err)
	}

	return nil
}

// process method wraps provided function for bulk or transaction processing.
func (multiCall MultiCall) process(f MultiCallFunc) Reply {
	var err error

	if multiCall.transaction {
		err = multiCall.conn.Send("MULTI")
		if err != nil {
			return Reply{err: redisError(err)}
		}
	}

	err = f(multiCall)
	if err != nil {
		return Reply{err: err}
	}

	var data interface{}
	if multiCall.transaction {
		data, err = multiCall.conn.Do("EXEC")
		if err != nil {
			return Reply{err: redisError(err)}
		}
	} else {
		err = multiCall.conn.Flush()
		if err != nil {
			return Reply{err: redisError(err)}
		}

		data, err = multiCall.conn.Receive()
		if err != nil {
			return Reply{err: redisError(err)}
		}
	}

	return Reply{data: data}
}
