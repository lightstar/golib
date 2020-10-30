// Package fileutil provides miscellaneous helper functions to work with files.
package fileutil

import (
	"bufio"
	"io/ioutil"
	"os"
)

// ReadAll function reads all data from file with given name.
// It returns retrieved bytes and any io error that occurs.
func ReadAll(name string) ([]byte, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	data, err := ioutil.ReadAll(bufio.NewReader(file))
	if err != nil {
		return nil, err
	}

	return data, nil
}
