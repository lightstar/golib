package middleware

import (
	"time"

	"github.com/lightstar/golib/pkg/http/httpservice"
	"github.com/lightstar/golib/pkg/http/httpservice/context"
	"github.com/lightstar/golib/pkg/log"
)

// Log middleware writes all requests into the log. Format: [address] uri [status] action: result (time).
func Log(logger log.Logger) httpservice.MiddlewareFunc {
	return func(handler httpservice.HandlerFunc) httpservice.HandlerFunc {
		return func(ctx *context.Context) error {
			beginTime := time.Now()

			err := handler(ctx)

			logger.Debug("[%s] %s [%d] %s: %s (%.2fms)", ctx.RemoteAddr(), ctx.RequestURI(), ctx.Status(),
				ctx.Action(), ctx.Result(), float64(time.Since(beginTime))/float64(time.Millisecond))

			return err
		}
	}
}
