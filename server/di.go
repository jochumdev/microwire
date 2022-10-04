// Code generated with jinja2 templates. DO NOT EDIT.

package server

import (
	"fmt"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/transport"
	"time"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct{}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin           = "server"
	cliArgAddress          = "server_address"
	cliArgID               = "server_id"
	cliArgMetadata         = "server_metadata"
	cliArgName             = "server_name"
	cliArgVersion          = "server_version"
	cliArgRegisterTTL      = "server_register_ttl"
	cliArgRegisterInterval = "server_register_interval"
)

func ProvideFlags(
	config *Config,
	cliConfig *cli.Config,
	c cli.Cli,
) (DiFlags, error) {
	if cliConfig.NoFlags {
		// Defined silently ignore that
		return DiFlags{}, nil
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPlugin)),
		cli.Usage("Server for go-micro; rpc"),
		cli.Default(config.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgAddress)),
		cli.Usage("Bind address for the server, eg: 127.0.0.1:8080"),
		cli.Default(config.Address),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgAddress)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgID)),
		cli.Usage("Id of the server. Auto-generated if not specified"),
		cli.Default(config.ID),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgID)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgMetadata)),
		cli.Usage(" A list of key-value pairs defining metadata, e.g.: version=1.0.0"),
		cli.Default([]string{}),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgMetadata)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgName)),
		cli.Usage("Name of the server. go.micro.srv.example"),
		cli.Default(config.Name),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgName)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgVersion)),
		cli.Usage("Version of the server. 1.1.0"),
		cli.Default(config.Version),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgVersion)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgRegisterTTL)),
		cli.Usage("Register TTL in seconds"),
		cli.Default(config.RegisterTTL),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgRegisterTTL)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgRegisterInterval)),
		cli.Usage("Register interval in seconds"),
		cli.Default(config.RegisterInterval),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgRegisterInterval)),
	); err != nil {
		return DiFlags{}, err
	}

	return DiFlags{}, nil
}

func ProvideConfig(
	_ di.DiConfig,
	flags DiFlags,
	config *Config,
	c cli.Cli,
	cliConfig *cli.Config,
	configor config.Config,
) (DiConfig, error) {
	defConfig := NewConfig()
	cfg := sourceConfig{Server: *defConfig}

	if configor != nil {
		if err := configor.Scan(&cfg); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&cfg.Server); err != nil {
		return DiConfig{}, err
	}

	if cliConfig.NoFlags {
		// Dont parse flags if NoFlags has been given
		return DiConfig{}, nil
	}

	defConfig = NewConfig()
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgPlugin)); ok {
		defConfig.Plugin = cli.FlagValue(f, defConfig.Plugin)
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgAddress)); ok {
		defConfig.Address = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgID)); ok {
		defConfig.ID = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgName)); ok {
		defConfig.Name = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgVersion)); ok {
		defConfig.Version = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgRegisterTTL)); ok {
		defConfig.RegisterTTL = cli.FlagValue(f, defConfig.RegisterTTL)
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgRegisterInterval)); ok {
		defConfig.RegisterInterval = cli.FlagValue(f, defConfig.RegisterInterval)
	}
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func ProvideConfigNoFlags(
	config *Config,
	configor config.Config,
) (DiConfig, error) {
	defConfig := NewConfig()
	c := sourceConfig{Server: *defConfig}

	if configor != nil {
		if err := configor.Scan(&c); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&c.Server); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,
	broker broker.Broker,
	log logger.Logger,
	registry registry.Registry,
	transport transport.Transport,
	config *Config,
) (Server, error) {
	if !config.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	pluginFunc, err := Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown plugin server: %s", config.Plugin)
	}

	opts := []Option{WithConfig(config)}
	if len(config.Address) > 0 {
		opts = append(opts, Address(config.Address))
	}
	if len(config.ID) > 0 {
		opts = append(opts, Id(config.ID))
	}
	if len(config.Name) > 0 {
		opts = append(opts, Name(config.Name))
	}
	if len(config.Version) > 0 {
		opts = append(opts, Version(config.Version))
	}

	if config.Logger.Enabled {
		loggerFunc, err := logger.Plugins.Get(config.Logger.Plugin)
		if err != nil {
			return nil, fmt.Errorf("{{Name}} unknown logger: %s", config.Logger.Plugin)
		}
		log = loggerFunc(logger.ConfigToOpts(config.Logger)...)
	}

	opts = append(
		opts,
		RegisterInterval(time.Duration(config.RegisterInterval)*time.Second),
		RegisterTTL(time.Duration(config.RegisterTTL)*time.Second),
		Broker(broker),
		Registry(registry),
		Transport(transport),
		WithLogger(log),
	)

	for _, w := range config.WrapSubscriber {
		opts = append(opts, WrapSubscriber(w))
	}
	for _, w := range config.WrapHandler {
		opts = append(opts, WrapHandler(w))
	}

	return pluginFunc(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
