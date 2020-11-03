// Package runman provides manager that can run a bunch of Runner interface implementations concurrently.
// When one of the runners finishes, all other ones will be canceled via context.
// You can also cancel them all prematurely by canceling passed context.
//
// Typical usage:
//      manager := runman.New(runner1, runner2, runner3)
//      manager.Run(ctx)
package runman

import (
	"context"
)

// Manager structure used to manage a list of runners.
// Note that it implements Runner interface too, so in theory you can feed one manager instance into another.
type Manager struct {
	runners []Runner
}

// New function creates new manager instance with a list of provided runners.
func New(runners ...Runner) *Manager {
	return &Manager{
		runners: runners,
	}
}

// Run method runs all runners in separate goroutines. When one of them finishes, all other will be canceled.
// When provided context cancels, all runners are canceled too. It returns only when everything finishes.
func (m *Manager) Run(ctx context.Context) {
	countDone := 0
	stateChan := make(chan struct{})
	cancelFuncs := make([]func(), 0, len(m.runners))

	for _, runner := range m.runners {
		runnerCtx, cancelFunc := context.WithCancel(context.Background())

		go func(ctx context.Context, runner Runner, stateChan chan<- struct{}) {
			runner.Run(ctx)
			stateChan <- struct{}{}
		}(runnerCtx, runner, stateChan)

		cancelFuncs = append(cancelFuncs, cancelFunc)
	}

	select {
	case <-stateChan:
		countDone++
	case <-ctx.Done():
	}

	for _, cancelFunc := range cancelFuncs {
		cancelFunc()
	}

	for countDone < len(m.runners) {
		<-stateChan
		countDone++
	}
}
