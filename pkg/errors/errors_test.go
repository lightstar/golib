package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/errors"
)

func TestError(t *testing.T) {
	err := errors.New("test")
	require.Equal(t, "test", err.Error())
}

func TestNewFmt(t *testing.T) {
	err := errors.NewFmt("this is test error with argument %d", 5)
	require.Equal(t, "this is test error with argument 5", err.Error())
}

func TestFormat(t *testing.T) {
	err := errors.New("test").WithStackTrace()

	require.Equal(t, fmt.Sprintf("%v", err), "test")
	require.Equal(t, fmt.Sprintf("%s", err), "test")
	require.Equal(t, fmt.Sprintf("%q", err), "\"test\"")

	require.Regexp(t, `^test\ngithub.com/lightstar/golib/pkg/errors_test.TestFormat\n\t`,
		fmt.Sprintf("%+v", err))
}

func TestCause(t *testing.T) {
	cause := errors.New("cause")
	err := errors.New("test").WithCause(cause)

	require.Same(t, cause, err.Unwrap())
}

func TestStd(t *testing.T) {
	cause := errors.New("cause")
	err := errors.New("test").WithCause(cause)

	require.Same(t, cause, errors.Unwrap(err))
	require.True(t, errors.Is(err, err))
	require.True(t, errors.Is(err, cause))
	require.False(t, errors.Is(err, errors.New("other")))
}

func TestStdAs(t *testing.T) {
	type customErrorType struct {
		*errors.Err
	}

	cause := &customErrorType{Err: errors.New("custom")}
	err := errors.New("test").WithCause(cause)
	var target *customErrorType

	require.True(t, errors.As(err, &target))
	require.Equal(t, "custom", target.Error())

	require.False(t, errors.As(errors.New("other"), &target))
}
