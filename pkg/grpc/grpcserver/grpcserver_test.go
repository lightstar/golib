package grpcserver_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/lightstar/golib/api/testproto"
	"github.com/lightstar/golib/internal/test/iotest"
	"github.com/lightstar/golib/pkg/grpc/grpcserver"
	"github.com/lightstar/golib/pkg/log"
)

type testService struct {
	testproto.UnimplementedTestServer
	handler func(input *testproto.Input) (*testproto.Output, error)
}

func (service *testService) GetData(_ context.Context, inp *testproto.Input) (*testproto.Output, error) {
	return service.handler(inp)
}

func TestServer(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-server"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	var gotA int32

	service := &testService{
		handler: func(input *testproto.Input) (*testproto.Output, error) {
			gotA = input.A

			return &testproto.Output{B: 42}, nil
		},
	}

	server := grpcserver.MustNew(
		grpcserver.WithName("test-server"),
		grpcserver.WithAddress("127.0.0.1:5050"),
		grpcserver.WithLogger(logger),
		grpcserver.WithRegisterFn(func(s *grpc.Server) {
			testproto.RegisterTestServer(s, service)
		}),
	)

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "127.0.0.1:5050", server.Address())

	ctx, cancel := context.WithCancel(context.Background())
	stopChan := make(chan struct{})

	go func() {
		server.Run(ctx)
		close(stopChan)
	}()

	data, err := getRemoteData(&testproto.Input{A: 5})
	require.NoError(t, err)

	require.Equal(t, gotA, int32(5))
	require.Equal(t, data.B, int32(42))

	cancel()
	<-stopChan

	_, err = getRemoteData(&testproto.Input{A: 5})
	require.Error(t, err)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}

func TestServerWait(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-server"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	handlerEnteredChan := make(chan struct{})
	handlerFinished := false

	service := &testService{
		handler: func(input *testproto.Input) (*testproto.Output, error) {
			close(handlerEnteredChan)
			time.Sleep(time.Second)

			handlerFinished = true

			return &testproto.Output{}, nil
		},
	}

	server := grpcserver.MustNew(
		grpcserver.WithName("test-server"),
		grpcserver.WithAddress("127.0.0.1:5050"),
		grpcserver.WithLogger(logger),
		grpcserver.WithRegisterFn(func(s *grpc.Server) {
			testproto.RegisterTestServer(s, service)
		}),
	)

	require.Equal(t, "test-server", server.Name())
	require.Equal(t, "127.0.0.1:5050", server.Address())

	ctx, cancel := context.WithCancel(context.Background())
	stopChan := make(chan struct{})

	go func() {
		server.Run(ctx)
		close(stopChan)
	}()

	time.Sleep(2 * time.Second)

	go func() {
		_, _ = getRemoteData(&testproto.Input{})
	}()

	<-handlerEnteredChan
	cancel()
	<-stopChan

	require.True(t, handlerFinished)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-server\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}

func getRemoteData(input *testproto.Input) (*testproto.Output, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, "127.0.0.1:5050",
		grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}

	client := testproto.NewTestClient(conn)

	return client.GetData(ctx, input)
}
