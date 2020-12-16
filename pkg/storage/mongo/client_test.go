package mongo_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/storage/mongo"
)

func TestPing(t *testing.T) {
	helper := newSessionHelper(t)
	defer helper.Close(t)

	err := helper.client.Ping()
	require.NoError(t, err)
}

func TestPingError(t *testing.T) {
	client, err := mongo.NewClient(
		mongo.WithAddress("unknown_address"),
		mongo.WithConnectTimeout(1),
		mongo.WithSocketTimeout(1),
	)
	require.NoError(t, err)

	err = client.Ping()
	require.Error(t, err)
}
