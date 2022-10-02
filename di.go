package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	"github.com/go-micro/microwire/di"
	mRegistry "github.com/go-micro/microwire/registry"
	mStore "github.com/go-micro/microwire/store"
	mTransport "github.com/go-micro/microwire/transport"
	"github.com/google/wire"

	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/store"
	"go-micro.dev/v4/transport"
)

type CliArgs []string

func NewService(opts ...Option) (micro.Service, error) {
	options := NewOptions(opts...)

	// Setup cli
	cliConfig := mCli.NewConfig()
	cliConfig.Cli.Plugin = "urfave"
	cliConfig.Cli.Name = options.Name
	cliConfig.Cli.Version = options.Version
	cliConfig.Cli.Description = options.Description
	cliConfig.Cli.Usage = options.Usage
	cliConfig.Cli.NoFlags = options.NoFlags
	cliConfig.Cli.Flags = options.Flags
	cliConfig.Cli.ConfigFile = options.Config

	// Setup Components
	brokerConfig := mBroker.NewConfig()
	registryConfig := mRegistry.NewConfig()
	storeConfig := mStore.NewConfig()
	transportConfig := mTransport.NewConfig()

	return newService(
		options,
		cliConfig,
		brokerConfig,
		registryConfig,
		storeConfig,
		transportConfig,
	)
}

func ProvideFlags(
	_ *mBroker.DiFlags,
	_ *mRegistry.DiFlags,
	_ *mTransport.DiFlags,
) (di.DiFlags, error) {
	return di.DiFlags{}, nil
}

func ProvideAllService(
	opts *Options,
	broker broker.Broker,
	registry registry.Registry,
	store store.Store,
	transport transport.Transport,
) (micro.Service, error) {
	mOpts := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
	}

	if broker != nil {
		mOpts = append(mOpts, micro.Broker(broker))
	}
	if registry != nil {
		mOpts = append(mOpts, micro.Registry(registry))
	}
	if store != nil {
		mOpts = append(mOpts, micro.Store(store))
	}
	if transport != nil {
		mOpts = append(mOpts, micro.Transport(transport))
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

func ProvideConfigFile(
	options *Options,
) (di.DiConfig, error) {
	return di.DiConfig(options.Config), nil
}

// DiAllSet is a set of all things components need, except the components themself.
var DiAllSet = wire.NewSet(
	di.ProvideConfigor,
	mBroker.DiSet,
	mRegistry.DiSet,
	mStore.DiSet,
	mTransport.DiSet,
)

var DiCliSet = wire.NewSet(
	mCli.ProvideCli,
	mCli.ProvideParsed,
	mCli.ProvideConfig,
)

var DiNoDiSet = wire.NewSet(
	ProvideFlags,
	ProvideAllService,
)
