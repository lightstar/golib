package mongo_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/storage/mongo"
)

type sessionHelper struct {
	client  *mongo.Client
	session *mongo.Session
}

func newSessionHelper(t *testing.T) *sessionHelper {
	t.Helper()

	testAddress := os.Getenv("TEST_MONGO_ADDRESS")
	if testAddress == "" {
		t.Skip("Provide 'TEST_MONGO_ADDRESS' environment variable to test mongo")
	}

	client, err := mongo.NewClient(mongo.WithAddress(testAddress))
	require.NoError(t, err)

	session := client.Session(context.Background(), "test", "test")

	return &sessionHelper{
		client:  client,
		session: session,
	}
}

func (helper *sessionHelper) Close(t *testing.T) {
	t.Helper()

	err := helper.session.Drop()
	require.NoError(t, err)

	err = helper.client.Close()
	require.NoError(t, err)
}

func (helper *sessionHelper) requireFindOne(t *testing.T, id interface{}, requiredData *sampleDataType) {
	t.Helper()

	data := sampleDataType{}
	count, err := helper.session.FindOne(mongo.Data{{Key: "_id", Value: id}}, &data)

	require.NoError(t, err)

	if requiredData == nil {
		require.Equal(t, int64(0), count)
	} else {
		require.Equal(t, int64(1), count)
		require.Equal(t, requiredData, &data)
	}
}
