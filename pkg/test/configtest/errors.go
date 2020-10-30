package configtest

import "github.com/lightstar/golib/pkg/errors"

var (
	// ErrNoData error is returned when predefined configuration data is nil.
	ErrNoData = errors.New("no data")

	// ErrNoSuchKey error is returned when requested key is not exists in predefined configuration data.
	ErrNoSuchKey = errors.New("no such key")

	// ErrOutputNotPointer error is returned when provided output parameter is not a pointer.
	ErrOutputNotPointer = errors.New("output must be a pointer")
)
