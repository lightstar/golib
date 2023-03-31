package encoder

import (
	"fmt"
	"io"
)

func writeString(writer io.Writer, data interface{}) error {
	var dataBytes []byte

	switch dataTyped := data.(type) {
	case []byte:
		dataBytes = dataTyped
	case string:
		dataBytes = []byte(dataTyped)
	case fmt.Stringer:
		dataBytes = []byte(dataTyped.String())
	case error:
		dataBytes = []byte(dataTyped.Error())
	default:
		dataBytes = []byte(fmt.Sprintf("%v", data))
	}

	_, err := writer.Write(dataBytes)

	return err
}
