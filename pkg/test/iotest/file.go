// Package iotest is used as a helper in tests that need files or use some other io.
package iotest

import (
	"os"
	"path/filepath"
	"testing"
)

const (
	testDirPerm  = 0o700
	testFilePerm = 0o600
)

// WriteFile function writes data into test file creating all needed directories by the way.
// It fails current test if file can't be created.
func WriteFile(t *testing.T, path string, data []byte) {
	t.Helper()

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, testDirPerm)
		if err != nil {
			t.Fatalf("can't create directory '%s': %v", dir, err)
		}
	}

	err := os.WriteFile(path, data, testFilePerm)
	if err != nil {
		t.Fatalf("can't write file '%s': %v", path, err)
	}
}

// RemoveFile function removes test file. It fails current test if file can't be removed.
func RemoveFile(t *testing.T, path string) {
	t.Helper()

	err := os.Remove(path)
	if err != nil {
		t.Fatalf("can't remove file '%s': %v", path, err)
	}
}
