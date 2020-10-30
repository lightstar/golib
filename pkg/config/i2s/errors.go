package i2s

import "github.com/lightstar/golib/pkg/errors"

var (
	// ErrMismatchedTypes error is returned when types of input and output fields doesn't match.
	ErrMismatchedTypes = errors.New("mismatched types")

	// ErrMapKeyNotString error is returned when some map has keys with type different from string.
	ErrMapKeyNotString = errors.New("map keys must be strings")

	// ErrOutputNotPointer error is returned when provided output parameter is not a pointer.
	ErrOutputNotPointer = errors.New("output must be a pointer")

	// ErrUnknownField error is returned when output structure doesn't have an appropriate field for input data.
	ErrUnknownField = errors.New("unknown field")

	// ErrUnsupportedType error is returned when source contains data of type unsupported by this package.
	ErrUnsupportedType = errors.New("unsupported type")
)
