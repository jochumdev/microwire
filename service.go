package micro

import (
	"os"
	"os/signal"
	rtime "runtime"

	"github.com/go-micro/microwire/v5/client"
	log "github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/server"
	signalutil "github.com/go-micro/microwire/v5/util/signal"
)

type service struct {
	opts Options
}

func newMicroService(opts ...Option) Service {
	options := NewOptions(opts...)
	return &service{
		opts: *options,
	}
}

func (s *service) Name() string {
	return s.opts.Server.Options().Name
}

// Init initializes options.
func (s *service) Init(opts ...Option) {
	// process options
	for _, o := range opts {
		o(&s.opts)
	}
}

func (s *service) Options() Options {
	return s.opts
}

func (s *service) Client() client.Client {
	return s.opts.Client
}

func (s *service) Server() server.Server {
	return s.opts.Server
}

func (s *service) String() string {
	return "micro"
}

func (s *service) Start() error {
	for _, fn := range s.opts.BeforeStart {
		if err := fn(s); err != nil {
			return err
		}
	}

	if err := s.opts.Server.Start(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStart {
		if err := fn(s); err != nil {
			return err
		}
	}

	return nil
}

func (s *service) Stop() error {
	var err error

	for _, fn := range s.opts.BeforeStop {
		if err = fn(s); err != nil {
			return err
		}
	}

	if err = s.opts.Server.Stop(); err != nil {
		return err
	}

	for _, fn := range s.opts.AfterStop {
		if err = fn(s); err != nil {
			return err
		}
	}

	return err
}

func (s *service) Run() (err error) {
	logger := s.opts.Logger

	// exit when help flag is provided
	for _, v := range os.Args[1:] {
		if v == "-h" || v == "--help" {
			os.Exit(0)
		}
	}

	// start the profiler
	if s.opts.Profile != nil {
		// to view mutex contention
		rtime.SetMutexProfileFraction(5)
		// to view blocking profile
		rtime.SetBlockProfileRate(1)

		if err = s.opts.Profile.Start(); err != nil {
			return err
		}
		defer func() {
			err = s.opts.Profile.Stop()
			if err != nil {
				logger.Log(log.ErrorLevel, err)
			}
		}()
	}

	logger.Logf(log.InfoLevel, "Starting [service] %s", s.Name())

	if err = s.Start(); err != nil {
		return err
	}

	ch := make(chan os.Signal, 1)
	if s.opts.Signal {
		signal.Notify(ch, signalutil.Shutdown()...)
	}

	select {
	// wait on kill signal
	case <-ch:
	// wait on context cancel
	case <-s.opts.Context.Done():
	}

	return s.Stop()
}
