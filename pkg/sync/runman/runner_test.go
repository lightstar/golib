package runman_test

import (
	"context"
	"time"
)

type runner struct {
	delay time.Duration
	done  bool
}

func (r *runner) Run(ctx context.Context) {
	if r.delay > 0 {
		select {
		case <-time.After(r.delay):
		case <-ctx.Done():
		}
	} else {
		<-ctx.Done()
	}

	r.done = true
}
