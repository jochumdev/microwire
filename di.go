package micro

import (
	"github.com/go-micro/microwire/v5/auth"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/cache"
	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/config/configdi"
	"github.com/go-micro/microwire/v5/di"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"
)

type CliArgs []string

func NewService(opts ...Option) (Service, error) {
	options := NewOptions(opts...)

	// Setup cli
	cliConfig := cli.NewConfig()
	cliConfig.Plugin = "urfave"
	cliConfig.Name = options.Name
	cliConfig.Version = options.Version
	cliConfig.Description = options.Description
	cliConfig.Usage = options.Usage
	cliConfig.ArgPrefix = options.ArgPrefix
	cliConfig.NoFlags = options.NoFlags
	cliConfig.Flags = options.Flags
	cliConfig.ConfigFile = options.ConfigFile

	serverConfig := server.NewConfig()
	serverConfig.Address = options.Address
	serverConfig.RegisterTTL = options.RegisterTTL
	serverConfig.RegisterInterval = options.RegisterInterval
	serverConfig.Metadata = options.Metadata
	serverConfig.WrapSubscriber = options.WrapSubscriber
	serverConfig.WrapHandler = options.WrapHandler

	return newService(
		options,
		cliConfig,
		auth.NewConfig(),
		broker.NewConfig(),
		cache.NewConfig(),
		client.NewConfig(),
		registry.NewConfig(),
		serverConfig,
		store.NewConfig(),
		transport.NewConfig(),
	)
}

func ProvideFlags(
	_ auth.DiFlags,
	_ broker.DiFlags,
	_ cache.DiFlags,
	_ client.DiFlags,
	_ registry.DiFlags,
	_ server.DiFlags,
	_ store.DiFlags,
	_ transport.DiFlags,
) (di.DiFlags, error) {
	return di.DiFlags{}, nil
}

func ProvideAllService(
	opts *Options,
	auth auth.Auth,
	broker broker.Broker,
	cache cache.Cache,
	client client.Client,
	registry registry.Registry,
	server server.Server,
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
	if client != nil {
		mOpts = append(mOpts, Client(client))
	}
	if cache != nil {
		mOpts = append(mOpts, Cache(cache))
	}
	if registry != nil {
		mOpts = append(mOpts, Registry(registry))
	}
	if server != nil {
		mOpts = append(mOpts, Server(server))
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
	client.DiSet,
	registry.DiSet,
	server.DiSet,
	store.DiSet,
	transport.DiSet,
)

var DiNoCliSet = wire.NewSet(
	configdi.ProvideConfigor,
	auth.DiNoCliSet,
	broker.DiNoCliSet,
	cache.DiNoCliSet,
	client.DiNoCliSet,
	registry.DiNoCliSet,
	server.DiNoCliSet,
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
