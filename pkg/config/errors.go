package config

import (
	"github.com/lightstar/golib/pkg/errors"
)

var (
	// ErrNoSuchKey error is returned when source configuration data doesn't have the wanted key.
	ErrNoSuchKey = errors.New("no such key")

	// ErrNotMap error is returned when retrieved data is not a map, but it has to be.
	ErrNotMap = errors.New("data by key is not a map")
)
