// Code generated with jinja2 templates. DO NOT EDIT.

package server

import (
	"fmt"
	"github.com/go-micro/microwire/v5/broker"
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
	if cliConfig.Cli.NoFlags {
		// Defined silently ignore that
		return DiFlags{}, nil
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		cli.Usage("Server for go-micro; rpc"),
		cli.Default(config.Server.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddress)),
		cli.Usage("Bind address for the server, eg: 127.0.0.1:8080"),
		cli.Default(config.Server.Address),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgAddress)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgID)),
		cli.Usage("Id of the server. Auto-generated if not specified"),
		cli.Default(config.Server.ID),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgID)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgMetadata)),
		cli.Usage(" A list of key-value pairs defining metadata, e.g.: version=1.0.0"),
		cli.Default([]string{}),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgMetadata)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgName)),
		cli.Usage("Name of the server. go.micro.srv.example"),
		cli.Default(config.Server.Name),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgName)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgVersion)),
		cli.Usage("Version of the server. 1.1.0"),
		cli.Default(config.Server.Version),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgVersion)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgRegisterTTL)),
		cli.Usage("Register TTL in seconds"),
		cli.Default(config.Server.RegisterTTL),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgRegisterTTL)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgRegisterInterval)),
		cli.Usage("Register interval in seconds"),
		cli.Default(config.Server.RegisterInterval),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgRegisterInterval)),
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

	if configor != nil {
		if err := configor.Scan(defConfig); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	if cliConfig.Cli.NoFlags {
		// Dont parse flags if NoFlags has been given
		return DiConfig{}, nil
	}

	defConfig = NewConfig()
	if f, ok := c.Get(cliArgPlugin); ok {
		defConfig.Server.Plugin = cli.FlagValue(f, defConfig.Server.Plugin)
	}
	if f, ok := c.Get(cliArgAddress); ok {
		defConfig.Server.Address = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cliArgID); ok {
		defConfig.Server.ID = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cliArgName); ok {
		defConfig.Server.Name = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cliArgVersion); ok {
		defConfig.Server.Version = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cliArgRegisterTTL); ok {
		defConfig.Server.RegisterTTL = cli.FlagValue(f, defConfig.Server.RegisterTTL)
	}
	if f, ok := c.Get(cliArgRegisterInterval); ok {
		defConfig.Server.RegisterInterval = cli.FlagValue(f, defConfig.Server.RegisterInterval)
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

	if configor != nil {
		if err := configor.Scan(defConfig); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,
	broker broker.Broker,
	registry registry.Registry,
	transport transport.Transport,
	config *Config,
) (Server, error) {
	if !config.Server.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Server.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown server: %v", err)
	}

	opts := []Option{}
	if len(config.Server.Address) > 0 {
		opts = append(opts, Address(config.Server.Address))
	}
	if len(config.Server.ID) > 0 {
		opts = append(opts, Id(config.Server.ID))
	}
	if len(config.Server.Name) > 0 {
		opts = append(opts, Name(config.Server.Name))
	}
	if len(config.Server.Version) > 0 {
		opts = append(opts, Version(config.Server.Version))
	}

	opts = append(
		opts,
		RegisterInterval(time.Duration(config.Server.RegisterInterval)*time.Second),
		RegisterTTL(time.Duration(config.Server.RegisterTTL)*time.Second),
		Broker(broker),
		Registry(registry),
		Transport(transport),
	)

	for _, w := range config.Server.WrapSubscriber {
		opts = append(opts, WrapSubscriber(w))
	}
	for _, w := range config.Server.WrapHandler {
		opts = append(opts, WrapHandler(w))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
