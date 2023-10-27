package mongo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSessionWithContext(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	ctx := context.Background()
	helper.session.WithContext(ctx)
	require.Equal(t, ctx, helper.session.Context())
}

func TestSessionWithDatabase(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	helper.session.WithDatabase("test2", "testCollection2")

	require.Equal(t, "test2", helper.session.Database())
	require.Equal(t, "testCollection2", helper.session.Collection())
}

func TestSessionWithCollection(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	helper.session.WithCollection("testCollection2")

	require.Equal(t, "test", helper.session.Database())
	require.Equal(t, "testCollection2", helper.session.Collection())
}
