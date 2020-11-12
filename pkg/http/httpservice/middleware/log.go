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
		return func(c *context.Context) error {
			beginTime := time.Now()

			err := handler(c)

			logger.Debug("[%s] %s [%d] %s: %s (%.2fms)", c.RemoteAddr(), c.RequestURI(), c.Status(),
				c.Action(), c.Result(), float64(time.Since(beginTime))/float64(time.Millisecond))

			return err
		}
	}
}
