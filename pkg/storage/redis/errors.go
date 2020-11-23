package redis

import (
	"github.com/lightstar/golib/pkg/errors"
)

// ErrNoValues error is returned by Scan and ScanStruct reply methods when there are no values returned by redis server,
// i.e. reply was empty.
var ErrNoValues = errors.New("no values")

func redisError(err error) error {
	return errors.NewFmt("redis error (%s)", err.Error()).WithCause(err).WithStackTrace()
}
