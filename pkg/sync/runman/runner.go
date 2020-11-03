package runman

import "context"

// Runner interface, implementations of which are managed by Manager structure.
type Runner interface {
	Run(ctx context.Context)
}
