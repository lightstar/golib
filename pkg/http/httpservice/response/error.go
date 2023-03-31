package response

//nolint:gochecknoglobals // these are some commonly-used error responses
var (
	// Unauthorized error response is supposed to be sent to the client on unauthorized error.
	Unauthorized = Error{Error: "unauthorized"}
	// BadRequest error response is supposed to be sent to the client on bad request error.
	BadRequest = Error{Error: "bad request"}
	// InternalError error response is supposed to be sent to the client on internal error.
	InternalError = Error{Error: "internal error"}
)

// Error structure provides the most basic error payload sent to the client.
type Error struct {
	Error string `json:"error"`
}

// String method implements Stringer interface.
func (response Error) String() string {
	return response.Error
}
