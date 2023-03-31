package encoder

import (
	"net/http"
)

//nolint:gochecknoglobals // this is commonly-used encoder as global variable
var (
	// Plain encoder writes provided data into response as plain text.
	Plain = Func(encodePlain)
)

func encodePlain(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(status)

	return writeString(w, data)
}
