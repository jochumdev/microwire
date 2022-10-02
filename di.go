package micro

import (
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

func NewService(opts ...Option) (Service, error) {
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
	cliConfig.Cli.ConfigFile = options.ConfigFile

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
	_ *mStore.DiFlags,
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
) (Service, error) {
	mOpts := []Option{
		Name(opts.Name),
		Version(opts.Version),
	}

	if broker != nil {
		mOpts = append(mOpts, Broker(broker))
	}
	if registry != nil {
		mOpts = append(mOpts, Registry(registry))
	}
	if store != nil {
		mOpts = append(mOpts, Store(store))
	}
	if transport != nil {
		mOpts = append(mOpts, Transport(transport))
	}

	for _, fn := range opts.BeforeStart {
		mOpts = append(mOpts, BeforeStart(fn))
	}
	for _, fn := range opts.BeforeStop {
		mOpts = append(mOpts, BeforeStop(fn))
	}
	for _, fn := range opts.AfterStart {
		mOpts = append(mOpts, AfterStart(fn))
	}
	for _, fn := range opts.AfterStop {
		mOpts = append(mOpts, AfterStop(fn))
	}

	service := NewMicroService(
		mOpts...,
	)

	return service, nil
}

func ProvideConfigFile(
	options *Options,
) (di.DiConfig, error) {
	return di.DiConfig(options.ConfigFile), nil
}

// DiSet is a set of all things components need, except the components themself.
var DiSet = wire.NewSet(
	configdi.ProvideConfigor,
	mBroker.DiSet,
	mRegistry.DiSet,
	mStore.DiSet,
	mTransport.DiSet,
)

var DiNoFlagsSet = wire.NewSet(
	configdi.ProvideConfigor,
	mBroker.DiNoFlagsSet,
	mRegistry.DiNoFlagsSet,
	mStore.DiNoFlagsSet,
	mTransport.DiNoFlagsSet,
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
