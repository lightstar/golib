package daemon

// Processor interface. You can assign it's implementation to daemon and it's Process method will be called at regular
// intervals according to delay setting.
type Processor interface {
	Process()
}

// ProcessFunc is a function implementing Processor interface.
type ProcessFunc func()

// Process method here just calls the underlying function itself.
func (processFunc ProcessFunc) Process() {
	processFunc()
}
