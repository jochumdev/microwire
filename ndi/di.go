package ndi

import (
	micro "github.com/go-micro/microwire/v5"
	mBroker "github.com/go-micro/microwire/v5/broker"
	mCli "github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config/configdi"
	"github.com/go-micro/microwire/v5/di"
	mRegistry "github.com/go-micro/microwire/v5/registry"
	mStore "github.com/go-micro/microwire/v5/store"
	mTransport "github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"

	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
)

type CliArgs []string

func NewService(opts ...micro.MwOption) (micro.Service, error) {
	options := micro.NewMwOptions(opts...)

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
	opts *micro.MwOptions,
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
	options *micro.MwOptions,
) (di.DiConfig, error) {
	return di.DiConfig(options.Config), nil
}

// DiSet is a set of all things components need, except the components themself.
var DiSet = wire.NewSet(
	configdi.ProvideConfigor,
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
