package config

import (
	"github.com/lightstar/goworld/pkg/errors"
)

var (
	// ErrNoSource error is returned when environment variables don't define appropriate source.
	ErrNoSource = errors.New("no source defined")

	// ErrUnknownEncoder error is returned when encoder defined in environment variable CONFIG_ENCODER is unsupported.
	ErrUnknownEncoder = errors.New("unknown encoder")

	// ErrNoSuchKey error is returned when source configuration data doesn't have the wanted key.
	ErrNoSuchKey = errors.New("no such key")

	// ErrNotMap error is returned when retrieved data is not a map but it has to be.
	ErrNotMap = errors.New("data by key is not a map")

	// ErrNoData error is returned when source doesn't have any configuration data.
	ErrNoData = errors.New("no data")
)
