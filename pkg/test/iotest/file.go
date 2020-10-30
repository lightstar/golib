// Package iotest is used as a helper in tests that need files or use some other io.
package iotest

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// WriteFile function writes data into test file creating all needed directories by the way.
// It fails current test if file can't be created.
func WriteFile(t *testing.T, path string, data []byte) {
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, os.ModeDir)
		if err != nil {
			t.Helper()
			t.Fatalf("can't create directory '%s': %v", dir, err)
		}
	}

	err := ioutil.WriteFile(path, data, 0600)
	if err != nil {
		t.Helper()
		t.Fatalf("can't write file '%s': %v", path, err)
	}
}

// RemoveFile function removes test file. It fails current test if file can't be removed.
func RemoveFile(t *testing.T, path string) {
	err := os.Remove(path)
	if err != nil {
		t.Helper()
		t.Fatalf("can't remove file '%s': %v", path, err)
	}
}
