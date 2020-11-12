// Package encoder provides response encoding functionality for the context package. Feed encoder interface
// implementation into the context object when you reset it, and it will be used to encode response.
//
// Typical usage:
//      enc := encoder.JSON
//      ctx.Reset(..., enc, ...)
//      // ...
//      if err := ctx.Encode(w, http.StatusOK, data); err != nil {
//          ctx.Logger().Error(err.Error())
//      }
package encoder

import (
	"net/http"
)

// Encoder interface that encodes arbitrary data and writes the result into http response.
type Encoder interface {
	Encode(http.ResponseWriter, int, interface{}) error
}

// Func is the function that implements Encoder interface by calling itself on Encode method.
type Func func(http.ResponseWriter, int, interface{}) error

// Encode method implements Encoder interface by calling the function itself.
func (f Func) Encode(w http.ResponseWriter, status int, data interface{}) error {
	return f(w, status, data)
}
