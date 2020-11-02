// Package daemon provides service that runs blocking loop.
// It will be stopped after receiving external system signal TERM or INT.
// You can also send TERM signal manually by calling Terminate method.
//
// Optionally you can set some custom processor that will be called continuously with provided delay.
//
// Typical usage:
//      // ... Setup any services you need and run them as goroutines
//      daemon.New(
//          daemon.WithName("my-daemon"),
//          daemon.WithDelay(1000),
//          daemon.WithProcessor(myProcessor),
//      ).Run()
//      // ... Shutdown your services and free resources
package daemon

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/lightstar/golib/pkg/log"
)

// Daemon structure that provides daemon service. Don't create manually, use the functions down below instead.
type Daemon struct {
	name      string
	delay     time.Duration
	logger    log.Logger
	processor Processor
	sigChan   chan os.Signal
}

// New function creates new daemon with provided options.
func New(opts ...Option) (*Daemon, error) {
	config, err := buildConfig(opts)
	if err != nil {
		return nil, err
	}

	logger := config.logger
	if logger == nil {
		var err error

		logger, err = log.New(log.WithName(config.name))
		if err != nil {
			return nil, err
		}
	}

	sigChan := make(chan os.Signal)

	return &Daemon{
		name:      config.name,
		delay:     config.delay,
		logger:    logger,
		processor: config.processor,
		sigChan:   sigChan,
	}, nil
}

// MustNew function creates new daemon with provided options and panics on any error.
func MustNew(opts ...Option) *Daemon {
	daemon, err := New(opts...)
	if err != nil {
		panic(err)
	}

	return daemon
}

// Name method gets daemon's name.
func (daemon *Daemon) Name() string {
	return daemon.name
}

// Delay method gets daemon's delay with which processor is called (if it is set).
func (daemon *Daemon) Delay() time.Duration {
	return daemon.delay
}

// SigChan method gets daemon's signal channel. You can send any value there to simulate incoming system signal.
func (daemon *Daemon) SigChan() chan<- os.Signal {
	return daemon.sigChan
}

// Run method runs daemon blocking loop. It will be stopped after receiving TERM or INT system signal.
// You can also pass cancellable context to stop daemon prematurely.
// If processor was set, it will be called at regular intervals according to delay setting.
func (daemon *Daemon) Run(ctx context.Context) {
	stopChan := make(chan struct{})

	signal.Notify(daemon.sigChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		select {
		case sig := <-daemon.sigChan:
			daemon.logger.Info("got signal '%s'", sig.String())
		case <-ctx.Done():
			daemon.logger.Info(ctx.Err().Error())
		}

		close(stopChan)
	}()

	daemon.logger.Info("started")

	if daemon.processor != nil {
	LOOP:
		for {
			select {
			case <-time.After(daemon.delay):
				daemon.processor.Process()
			case <-stopChan:
				break LOOP
			}
		}
	} else {
		<-stopChan
	}

	daemon.logger.Info("stopped")
	daemon.logger.Sync()
}
