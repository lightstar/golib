// Package response provides wrapper around http.ResponseWriter used inside context object, also it provides
// some commonly-used response payloads.
package response

import "net/http"

// Writer structure is a wrapper around http.ResponseWriter. It can be used to intercept and retrieve written
// http status. It is used inside context object and you won't probably use it yourself.
type Writer struct {
	http.ResponseWriter
	status        int
	headerWritten bool
}

// Status method retrieves written http status code.
func (w *Writer) Status() int {
	return w.status
}

// HeaderWritten method retrieves flag - is header already written or not.
func (w *Writer) HeaderWritten() bool {
	return w.headerWritten
}

// WriteHeader method writes header into the response and intercepts status code.
func (w *Writer) WriteHeader(status int) {
	if w.headerWritten {
		return
	}

	w.status = status
	w.ResponseWriter.WriteHeader(status)
	w.headerWritten = true
}

// Write method writes data into the response and calls WriteHeader method if header is not already written.
func (w *Writer) Write(data []byte) (int, error) {
	if !w.headerWritten {
		w.WriteHeader(http.StatusOK)
	}

	return w.ResponseWriter.Write(data)
}

// Reset method prepares structure for the next request.
func (w *Writer) Reset(responseWriter http.ResponseWriter) {
	w.ResponseWriter = responseWriter
	w.status = 0
	w.headerWritten = false
}
