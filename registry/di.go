// Code generated with jinja2 templates. DO NOT EDIT.

package registry

import (
	"fmt"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct{}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin    = "registry"
	cliArgAddresses = "registry_address"
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
		cli.Usage("Registry for discovery. etcd, mdns"),
		cli.Default(config.Registry.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		cli.Usage("List of registry addresses"),
		cli.Default(config.Registry.Addresses),
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
		defConfig.Registry.Plugin = cli.FlagValue(f, defConfig.Registry.Plugin)
	}
	if f, ok := c.Get(cliArgAddresses); ok {
		defConfig.Registry.Addresses = cli.FlagValue(f, []string{})
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

	config *Config,
) (Registry, error) {
	if !config.Registry.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Registry.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown registry: %v", err)
	}

	opts := []Option{}
	if len(config.Registry.Addresses) > 0 {
		opts = append(opts, Addrs(config.Registry.Addresses...))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
