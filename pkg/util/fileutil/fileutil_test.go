package fileutil_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/goworld/pkg/test/iotest"
	"github.com/lightstar/goworld/pkg/util/fileutil"
)

const testFilePath = "../../../test/file"

func TestReadAll(t *testing.T) {
	expected := "Hello, this is test file!\r\nBye!\r\n"

	iotest.WriteFile(t, testFilePath, []byte(expected))
	defer iotest.RemoveFile(t, testFilePath)

	data, err := fileutil.ReadAll(testFilePath)
	require.NoError(t, err)

	require.Equal(t, expected, string(data))
}

func TestReadAllError(t *testing.T) {
	_, err := fileutil.ReadAll(testFilePath)
	require.Error(t, err)
}
