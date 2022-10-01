package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	"github.com/go-micro/microwire/di"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"

	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
)

type CliArgs []string

func NewService(opts ...Option) (micro.Service, error) {
	options := NewOptions(opts)

	// Setup cli
	cliConfig := mCli.NewConfig()
	cliConfig.Name = options.Name
	cliConfig.Version = options.Version
	cliConfig.Description = options.Description
	cliConfig.Usage = options.Usage
	cliConfig.Flags = options.Flags
	cliConfig.ConfigFile = options.Config

	// Setup Components
	brokerConfig := mBroker.NewConfig()
	registryConfig := mRegistry.NewConfig()
	transportConfig := mTransport.NewConfig()

	return newService(options, cliConfig, brokerConfig, registryConfig, transportConfig)
}

func ProvideFlags(
	_ *mBroker.DiFlags,
) (di.DiFlags, error) {
	return di.DiFlags{}, nil
}

func ProvideService(
	opts *Options,
	broker broker.Broker,
) (micro.Service, error) {
	mOpts := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
	}

	if broker != nil {
		mOpts = append(mOpts, micro.Broker(broker))
	}

	for _, fn := range opts.BeforeStart {
		mOpts = append(mOpts, micro.BeforeStart(fn))
	}
	for _, fn := range opts.BeforeStop {
		mOpts = append(mOpts, micro.BeforeStop(fn))
	}
	for _, fn := range opts.AfterStart {
		mOpts = append(mOpts, micro.AfterStart(fn))
	}
	for _, fn := range opts.AfterStop {
		mOpts = append(mOpts, micro.AfterStop(fn))
	}

	service := micro.NewService(
		mOpts...,
	)

	for _, fn := range opts.Actions {
		if err := fn(service); err != nil {
			return nil, err
		}
	}

	return service, nil
}
