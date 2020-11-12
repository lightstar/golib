package httpservice

// MiddlewareFunc is a function that will wrap request handler to provide some shared functionality.
type MiddlewareFunc func(HandlerFunc) HandlerFunc

// UseMiddleware method applies provided middleware function that will wrap all handlers.
// If you call this method several times, each subsequent middleware will be applied after the previous ones.
func (service *Service) UseMiddleware(m MiddlewareFunc) {
	if service.middleware == nil {
		service.middleware = m
	} else {
		mOld := service.middleware

		service.middleware = func(handler HandlerFunc) HandlerFunc {
			return mOld(m(handler))
		}
	}
}

// UseMiddlewareBefore method applies provided middleware function that will wrap all handlers.
// If you call this method several times, each subsequent middleware will be applied before the previous ones.
func (service *Service) UseMiddlewareBefore(m MiddlewareFunc) {
	if service.middleware == nil {
		service.middleware = m
	} else {
		mOld := service.middleware

		service.middleware = func(handler HandlerFunc) HandlerFunc {
			return m(mOld(handler))
		}
	}
}
