package runman_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/sync/runman"
)

func TestManager(t *testing.T) {
	r1 := &runner{delay: time.Second}
	r2 := &runner{}
	r3 := &runner{}

	manager := runman.New(r1, r2, r3)
	manager.Run(context.Background())

	require.True(t, r1.done)
	require.True(t, r2.done)
	require.True(t, r3.done)
}

func TestManagerWithContext(t *testing.T) {
	r1 := &runner{}
	r2 := &runner{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	m := runman.New(r1, r2)
	m.Run(ctx)

	require.True(t, r1.done)
	require.True(t, r2.done)
}
