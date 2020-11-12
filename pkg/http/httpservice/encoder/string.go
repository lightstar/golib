package encoder

import (
	"fmt"
	"io"
)

func writeString(w io.Writer, data interface{}) error {
	var b []byte

	switch dataTyped := data.(type) {
	case []byte:
		b = dataTyped
	case string:
		b = []byte(dataTyped)
	case fmt.Stringer:
		b = []byte(dataTyped.String())
	case error:
		b = []byte(dataTyped.Error())
	default:
		b = []byte(fmt.Sprintf("%v", data))
	}

	_, err := w.Write(b)

	return err
}
