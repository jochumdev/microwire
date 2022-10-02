package micro

import (
	"github.com/go-micro/microwire/v5/auth"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/cache"
	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config/configdi"
	"github.com/go-micro/microwire/v5/di"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"
)

type CliArgs []string

func NewService(opts ...Option) (Service, error) {
	options := NewOptions(opts...)

	// Setup cli
	cliConfig := cli.NewConfig()
	cliConfig.Cli.Plugin = "urfave"
	cliConfig.Cli.Name = options.Name
	cliConfig.Cli.Version = options.Version
	cliConfig.Cli.Description = options.Description
	cliConfig.Cli.Usage = options.Usage
	cliConfig.Cli.NoFlags = options.NoFlags
	cliConfig.Cli.Flags = options.Flags
	cliConfig.Cli.ConfigFile = options.ConfigFile

	return newService(
		options,
		cliConfig,
		auth.NewConfig(),
		broker.NewConfig(),
		cache.NewConfig(),
		registry.NewConfig(),
		store.NewConfig(),
		transport.NewConfig(),
	)
}

func ProvideFlags(
	_ *auth.DiFlags,
	_ *broker.DiFlags,
	_ *cache.DiFlags,
	_ *registry.DiFlags,
	_ *store.DiFlags,
	_ *transport.DiFlags,
) (di.DiFlags, error) {
	return di.DiFlags{}, nil
}

func ProvideAllService(
	opts *Options,
	auth auth.Auth,
	broker broker.Broker,
	cache cache.Cache,
	registry registry.Registry,
	store store.Store,
	transport transport.Transport,
) (Service, error) {
	mOpts := []Option{
		Name(opts.Name),
		Version(opts.Version),
	}

	if auth != nil {
		mOpts = append(mOpts, Auth(auth))
	}
	if broker != nil {
		mOpts = append(mOpts, Broker(broker))
	}
	if cache != nil {
		mOpts = append(mOpts, Cache(cache))
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
	auth.DiSet,
	broker.DiSet,
	cache.DiSet,
	registry.DiSet,
	store.DiSet,
	transport.DiSet,
)

var DiNoCliSet = wire.NewSet(
	configdi.ProvideConfigor,
	auth.DiNoCliSet,
	broker.DiNoCliSet,
	cache.DiNoCliSet,
	registry.DiNoCliSet,
	store.DiNoCliSet,
	transport.DiNoCliSet,
)

var DiCliSet = wire.NewSet(
	cli.ProvideCli,
	cli.ProvideParsed,
	cli.ProvideConfig,
)

var DiNoDiSet = wire.NewSet(
	ProvideFlags,
	ProvideAllService,
)
