package daemon_test

import (
	"context"
	"syscall"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/lightstar/golib/pkg/daemon"
	"github.com/lightstar/golib/pkg/log"
	"github.com/lightstar/golib/pkg/test/iotest"
)

func TestDaemon(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-daemon"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	testProcessor := &processor{}

	dmn, err := daemon.New(
		daemon.WithName("test-daemon"),
		daemon.WithDelay(2000),
		daemon.WithProcessor(testProcessor),
		daemon.WithLogger(logger),
	)
	require.NoError(t, err)

	require.Equal(t, "test-daemon", dmn.Name())
	require.Equal(t, 2*time.Second, dmn.Delay())

	go func() {
		<-time.After(5 * time.Second)

		select {
		case dmn.SigChan() <- syscall.SIGTERM:
		default:
		}
	}()

	dmn.Run(context.Background())

	require.Equal(t, 2, testProcessor.processCalled)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) got signal 'terminated'\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}

func TestProcessFunc(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-daemon"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	processCalled := 0

	dmn, err := daemon.New(
		daemon.WithName("test-daemon"),
		daemon.WithDelay(2000),
		daemon.WithProcessFunc(func() {
			processCalled++
		}),
		daemon.WithLogger(logger),
	)
	require.NoError(t, err)

	require.Equal(t, "test-daemon", dmn.Name())
	require.Equal(t, 2*time.Second, dmn.Delay())

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dmn.Run(ctx)

	require.Equal(t, 1, processCalled)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) context deadline exceeded\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}

func TestNopDaemon(t *testing.T) {
	stdout := iotest.NewBuffer()
	stderr := iotest.NewBuffer()
	logger := log.MustNew(
		log.WithName("test-daemon"),
		log.WithStdout(stdout),
		log.WithStderr(stderr),
	)

	dmn, err := daemon.New(
		daemon.WithName("test-daemon"),
		daemon.WithLogger(logger),
	)
	require.NoError(t, err)

	require.Equal(t, "test-daemon", dmn.Name())

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	dmn.Run(ctx)

	require.Regexp(t, `^\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) started\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) context deadline exceeded\n`+
		`\[\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}\] \(test-daemon\) stopped\n$`, stdout.String())
	require.Empty(t, stderr.String())
}
