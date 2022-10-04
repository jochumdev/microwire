// Code generated with jinja2 templates. DO NOT EDIT.

package broker

import (
	"fmt"
	"github.com/go-micro/microwire/v5/registry"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct{}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin    = "broker"
	cliArgAddresses = "broker_address"
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
		cli.Usage("Broker for pub/sub. http, nats, rabbitmq"),
		cli.Default(config.Broker.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		cli.Usage("List of broker addresses"),
		cli.Default(config.Broker.Addresses),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
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
		defConfig.Broker.Plugin = cli.FlagValue(f, defConfig.Broker.Plugin)
	}
	if f, ok := c.Get(cliArgAddresses); ok {
		defConfig.Broker.Addresses = cli.FlagValue(f, []string{})
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
	registry registry.Registry,
	config *Config,
) (Broker, error) {
	if !config.Broker.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Broker.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown broker: %v", err)
	}

	opts := []Option{}
	if len(config.Broker.Addresses) > 0 {
		opts = append(opts, Addrs(config.Broker.Addresses...))
	}

	opts = append(opts, Registry(registry))

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
