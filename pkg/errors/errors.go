// Package errors provides basic error primitive with optional stack trace and cause error.
//
// Examples of typical usages:
//      err := errors.New("some message")
//      err := errors.NewFmt("some message with param %s", myParam)
//      err := errors.New("some error with stack trace and cause").WithStackTrace().WithCause(causeError)
//
// In non-trivial cases using subtypes is encouraged, such as:
//      type MyError struct {
//          *errors.Err
//          MyField string
//          MyOtherField int
//      }
package errors

import (
	"fmt"
	"io"
	"runtime"

	"github.com/pkg/errors"
)

// Err structure that implements error interface. Don't create it manually, use one of the functions down below.
type Err struct {
	msg        string
	cause      error
	stackTrace errors.StackTrace
}

// New function creates new error from text message.
func New(msg string) *Err {
	return &Err{
		msg: msg,
	}
}

// NewFmt method creates new error from formatted text message.
func NewFmt(format string, args ...interface{}) *Err {
	return &Err{
		msg: fmt.Sprintf(format, args...),
	}
}

// WithCause method adds cause to error.
func (err *Err) WithCause(cause error) *Err {
	err.cause = cause
	return err
}

// WithStackTrace method adds stacktrace to error.
func (err *Err) WithStackTrace() *Err {
	err.stackTrace = stackTrace()
	return err
}

// Error method retrieves error message.
func (err *Err) Error() string {
	return err.msg
}

// Unwrap method retrieves error cause implementing internal interface inside 'errors' standard package.
func (err *Err) Unwrap() error {
	return err.cause
}

// Format method implements internal interface inside 'fmt' standard package.
func (err *Err) Format(s fmt.State, verb rune) {
	switch verb {
	case 'v':
		_, _ = io.WriteString(s, err.msg)

		if s.Flag('+') {
			err.stackTrace.Format(s, verb)
		}
	case 's':
		_, _ = io.WriteString(s, err.msg)
	case 'q':
		_, _ = fmt.Fprintf(s, "%q", err.msg)
	}
}

// stackTrace function retrieves current stack trace.
func stackTrace() errors.StackTrace {
	const depth = 32
	const skip = 3

	var pcs [depth]uintptr
	n := runtime.Callers(skip, pcs[:])

	f := make([]errors.Frame, n)
	for i := 0; i < len(f); i++ {
		f[i] = errors.Frame(pcs[i])
	}

	return f
}
