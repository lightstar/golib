package encoder

import (
	"net/http"
)

//nolint:gochecknoglobals // this is commonly-used encoder as global variable
var (
	// HTML encoder writes provided data into response as html text.
	HTML = Func(encodeHTML)
)

func encodeHTML(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(status)

	return writeString(w, data)
}
