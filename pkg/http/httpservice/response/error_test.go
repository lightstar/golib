package response_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/http/httpservice/response"
)

func TestError(t *testing.T) {
	e := &response.Error{Error: "some error"}
	require.Equal(t, "some error", e.String())
}
