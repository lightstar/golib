package context

import (
	"net/http"

	"github.com/lightstar/golib/pkg/http/httpservice/response"
)

// DoResponse method sends data to the client and sets result.
// Error is unlikely, but if any, it will be immediately put into the log for simplicity, not returned.
func (ctx *Context) DoResponse(data interface{}, status int, result string) {
	ctx.SetResult(result)

	err := ctx.Encode(status, data)
	if err != nil {
		ctx.logger.Error(err.Error())
	}
}

// InternalErrorResponse method sends internal error response to the client and sets result.
// If result parameter is empty string, default result will be used which is "internal error".
func (ctx *Context) InternalErrorResponse(result string) {
	if result == "" {
		result = "internal error"
	}

	ctx.DoResponse(response.InternalError, http.StatusInternalServerError, result)
}

// UnauthorizedResponse method sends unauthorized response to the client and sets result.
// If result parameter is empty string, default result will be used which is "unauthorized".
func (ctx *Context) UnauthorizedResponse(result string) {
	if result == "" {
		result = "unauthorized"
	}

	ctx.DoResponse(response.Unauthorized, http.StatusUnauthorized, result)
}

// BadRequestResponse method sends bad request response to the client and sets result.
// If data parameter is nil, default response will be used which is response.BadRequest.
// Similarly, if result parameter is empty string, default result will be used which is "bad request".
func (ctx *Context) BadRequestResponse(data interface{}, result string) {
	if data == nil {
		data = response.BadRequest
	}

	if result == "" {
		result = "bad request"
	}

	ctx.DoResponse(data, http.StatusBadRequest, result)
}

// OKResponse method sends ok response to the client and sets result.
// If data parameter is nil, default response will be used which is response.Ok.
// Similarly, if result parameter is empty string, default result will be used which is "ok".
func (ctx *Context) OKResponse(data interface{}, result string) {
	if data == nil {
		data = response.Ok
	}

	if result == "" {
		result = "ok"
	}

	ctx.DoResponse(data, http.StatusOK, result)
}
