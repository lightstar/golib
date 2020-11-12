package response_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/response"
)

func TestMessage(t *testing.T) {
	e := &response.Message{Message: "some message"}
	require.Equal(t, "some message", e.String())
}
