package log

// Logger interface.
type Logger interface {
	Debug(msg string, v ...interface{})
	Info(msg string, v ...interface{})
	Error(msg string, v ...interface{})
	Fatal(msg string, v ...interface{})
	Sync()
}
