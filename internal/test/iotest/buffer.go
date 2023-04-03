package iotest

import "bytes"

// Buffer structure that implements ReadWriteSyncer interface used in tests.
type Buffer struct {
	*bytes.Buffer
}

// NewBuffer function creates new buffer.
func NewBuffer() *Buffer {
	return &Buffer{new(bytes.Buffer)}
}

// Sync method that does nothing but needed to satisfy interface.
func (buf *Buffer) Sync() error {
	return nil
}
