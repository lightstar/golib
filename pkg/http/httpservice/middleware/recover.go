package middleware

import (
	"github.com/lightstar/golib/pkg/errors"
	"github.com/lightstar/golib/pkg/http/httpservice"
	"github.com/lightstar/golib/pkg/http/httpservice/context"
)

// Recover middleware converts all panics into errors.
func Recover() httpservice.MiddlewareFunc {
	return func(handler httpservice.HandlerFunc) httpservice.HandlerFunc {
		return func(c *context.Context) (err error) {
			defer func() {
				if rcv := recover(); rcv != nil {
					defer func() {
						_ = recover()
					}()

					err = errors.NewFmt("recovered from panic: %+v", rcv)
				}
			}()

			return handler(c)
		}
	}
}
