package httpservice

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/lightstar/golib/pkg/http/httpservice/context"
)

// HandlerFunc is a function that will handle each request under some route with all used middlewares.
// If it returns an error, it will be written into the log and the internal server response will be sent to the client.
type HandlerFunc func(ctx *context.Context) error

// ServeHTTP method implements http.Handler interface.
func (service *Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	service.router.ServeHTTP(w, r)
}

// GET method applies handler for all GET requests coming to provided path.
// Action is an arbitrary string that identifies this specific handler (used in logs for example).
func (service *Service) GET(path string, action string, handler HandlerFunc) {
	service.router.GET(path, service.handle(action, handler))
}

// POST method applies handler for all POST requests coming to provided path.
// Action is an arbitrary string that identifies this specific handler (used in logs for example).
func (service *Service) POST(path string, action string, handler HandlerFunc) {
	service.router.POST(path, service.handle(action, handler))
}

// PUT method applies handler for all PUT requests coming to provided path.
// Action is an arbitrary string that identifies this specific handler (used in logs for example).
func (service *Service) PUT(path string, action string, handler HandlerFunc) {
	service.router.PUT(path, service.handle(action, handler))
}

// DELETE method applies handler for all DELETE requests coming to provided path.
// Action is an arbitrary string that identifies this specific handler (used in logs for example).
func (service *Service) DELETE(path string, action string, handler HandlerFunc) {
	service.router.DELETE(path, service.handle(action, handler))
}

// OPTIONS method applies handler for all OPTIONS requests coming to provided path.
// Action is an arbitrary string that identifies this specific handler (used in logs for example).
func (service *Service) OPTIONS(path string, action string, handler HandlerFunc) {
	service.router.OPTIONS(path, service.handle(action, handler))
}

func (service *Service) handle(action string, handler HandlerFunc) httprouter.Handle {
	if service.middleware != nil {
		handler = service.middleware(handler)
	}

	enc := service.encoder
	dec := service.decoder

	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		c := service.pool.Get().(*context.Context)
		c.Reset(w, r, params, enc, dec, action)

		if err := handler(c); err != nil {
			service.error(err, c)
		}

		service.pool.Put(c)
	}
}
