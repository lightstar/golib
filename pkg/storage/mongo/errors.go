package mongo

import (
	"github.com/lightstar/golib/pkg/errors"
)

func mongoError(err error) error {
	return errors.NewFmt("mongo error (%s)", err.Error()).WithCause(err).WithStackTrace()
}
