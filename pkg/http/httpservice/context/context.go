// Package context provides context object for data that belongs to an individual http request.
// After creating context with New function you must reset it with Reset method, and then use Reset method for each
// new request.
//
// Typical usage:
//      ctx := context.New(logger)
//
//      // ... Request comes
//      ctx.Reset(w, r, enc, dec, action)
//      if err := handler(ctx); err != nil {
//          // Handle error
//      }
//
//      // ... Another request comes
//      ctx.Reset(w, r, enc, dec, action)
//      if err := handler(ctx); err != nil {
//          // Handle error
//      }
package context

import (
	"net/http"

	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/http/httpservice/decoder"
	"github.com/lightstar/golib/pkg/http/httpservice/encoder"
	"github.com/lightstar/golib/pkg/http/httpservice/response"
	"github.com/lightstar/golib/pkg/log"
)

// Context structure that provides context functionality.
type Context struct {
	action   string
	result   string
	response *response.Writer
	request  *http.Request
	encoder  encoder.Encoder
	decoder  decoder.Decoder
	logger   log.Logger
}

// New function creates new context object.
func New(logger log.Logger) *Context {
	return &Context{
		response: new(response.Writer),
		logger:   logger,
	}
}

// Logger method retrieves logger object.
func (ctx *Context) Logger() log.Logger {
	return ctx.logger
}

// SetResult method sets arbitrary request result that can be used later (in middleware.Log for example).
func (ctx *Context) SetResult(result string) {
	ctx.result = result
}

// Result method retrieves request result that was set previously. Initially (after Reset) it equals to "unknown".
func (ctx *Context) Result() string {
	return ctx.result
}

// Status method retrieves status code that was set in response. Initially (after Reset) it equals to 0.
func (ctx *Context) Status() int {
	return ctx.response.Status()
}

// Action method retrieves action name that was set in Reset method.
func (ctx *Context) Action() string {
	return ctx.action
}

// Response method retrieves inner response object.
func (ctx *Context) Response() *response.Writer {
	return ctx.response
}

// Request method retrieves inner request object.
func (ctx *Context) Request() *http.Request {
	return ctx.request
}

// RemoteAddr method retrieves client's remote address.
func (ctx *Context) RemoteAddr() string {
	return ctx.request.RemoteAddr
}

// RequestURI method retrieves request's URI.
func (ctx *Context) RequestURI() string {
	return ctx.request.RequestURI
}

// Encode method encodes provided data using encoder (that was set in Reset method) and writes it into response.
func (ctx *Context) Encode(status int, data interface{}) error {
	err := ctx.encoder.Encode(ctx.response, status, data)
	if err != nil {
		return errors.NewFmt("can't encode data '%v' (%s)", data, err.Error()).WithCause(err)
	}

	return nil
}

// Decode method decodes data from request using decoder (that was set in Reset method) and puts it into data parameter.
func (ctx *Context) Decode(data interface{}) error {
	err := ctx.decoder.Decode(ctx.request, data)
	if err != nil {
		return errors.NewFmt("can't decode data (%s)", err.Error()).WithCause(err)
	}

	return nil
}

// Reset methods prepares context to handling new request.
func (ctx *Context) Reset(
	w http.ResponseWriter,
	r *http.Request,
	enc encoder.Encoder,
	dec decoder.Decoder,
	action string,
) {
	ctx.response.Reset(w)

	ctx.action = action
	ctx.result = "unknown"
	ctx.request = r
	ctx.encoder = enc
	ctx.decoder = dec
}
