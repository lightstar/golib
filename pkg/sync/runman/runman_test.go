package runman_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/sync/runman"
)

func TestManager(t *testing.T) {
	runner1 := &runner{delay: time.Second}
	runner2 := &runner{}
	runner3 := &runner{}

	manager := runman.New(runner1, runner2, runner3)
	manager.Run(context.Background())

	require.True(t, runner1.done)
	require.True(t, runner2.done)
	require.True(t, runner3.done)
}

func TestManagerWithContext(t *testing.T) {
	runner1 := &runner{}
	runner2 := &runner{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	m := runman.New(runner1, runner2)
	m.Run(ctx)

	require.True(t, runner1.done)
	require.True(t, runner2.done)
}
