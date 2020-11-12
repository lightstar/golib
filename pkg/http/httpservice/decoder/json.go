package decoder

import (
	"encoding/json"
	"net/http"
)

// nolint: gochecknoglobals // this is commonly-used decoder as global variable
var (
	// JSON decoder treats request body as json.
	JSON = Func(decodeJSON)
)

func decodeJSON(r *http.Request, data interface{}) error {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		return err
	}

	return nil
}
