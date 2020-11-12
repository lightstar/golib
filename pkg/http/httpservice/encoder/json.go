package encoder

import (
	"encoding/json"
	"net/http"
)

// nolint: gochecknoglobals // this is commonly-used encoder as global variable
var (
	// JSON encoder writes provided data into response as json payload.
	JSON = Func(encodeJSON)
)

func encodeJSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		return err
	}

	return nil
}
