package errors

import "errors"

// Is function delegates to standard errors function from go1.13+.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// As function delegates to standard errors function from go1.13+.
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Unwrap function delegates to standard errors function from go1.13+.
func Unwrap(err error) error {
	return errors.Unwrap(err)
}
