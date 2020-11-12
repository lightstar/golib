// Package decoder provides request decoding functionality for the context. Feed decoder interface
// implementation into the context object when you reset it, and it will be used to decode request data.
//
// Typical usage:
//      dec := decoder.JSON
//      ctx.Reset(..., dec, ...)
//      // ...
//      if err := ctx.Decode(r, &data); err != nil {
//          ctx.BadRequestResponse(nil, err.Error())
//      }
package decoder

import "net/http"

// Decoder interface that decodes http request into arbitrary data.
type Decoder interface {
	Decode(*http.Request, interface{}) error
}

// Func is the function that implements Decoder interface by calling itself on Decode method.
type Func func(*http.Request, interface{}) error

// Decode method implements Decoder interface by calling the function itself.
func (f Func) Decode(r *http.Request, data interface{}) error {
	return f(r, data)
}
