package response

// nolint: gochecknoglobals // these are some commonly-used message responses
var (
	// Ok response is supposed to be sent to the client on success.
	Ok = Message{Message: "OK"}
)

// Message structure provides the most basic payload with message sent to the client.
type Message struct {
	Message string `json:"message"`
}

// String method implements Stringer interface.
func (response Message) String() string {
	return response.Message
}
