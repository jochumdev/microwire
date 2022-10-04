// Code generated with jinja2 templates. DO NOT EDIT.

package store

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
	cliArgPlugin    = "store"
	cliArgAddresses = "store_address"
	cliArgDatabase  = "store_database"
	cliArgTable     = "store_table"
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
		cli.Usage("Store for pub/sub. http, nats, rabbitmq"),
		cli.Default(config.Store.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		cli.Usage("List of store addresses"),
		cli.Default(config.Store.Addresses),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgDatabase)),
		cli.Usage("Database option for the underlying store"),
		cli.Default(config.Store.Database),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgDatabase)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgTable)),
		cli.Usage("Table option for the underlying store"),
		cli.Default(config.Store.Table),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgTable)),
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
		defConfig.Store.Plugin = cli.FlagValue(f, defConfig.Store.Plugin)
	}
	if f, ok := c.Get(cliArgAddresses); ok {
		defConfig.Store.Addresses = cli.FlagValue(f, []string{})
	}
	if f, ok := c.Get(cliArgDatabase); ok {
		defConfig.Store.Database = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cliArgTable); ok {
		defConfig.Store.Table = cli.FlagValue(f, "")
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
) (Store, error) {
	if !config.Store.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Store.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown store: %v", err)
	}

	opts := []Option{}
	if len(config.Store.Addresses) > 0 {
		opts = append(opts, Nodes(config.Store.Addresses...))
	}
	if len(config.Store.Database) > 0 {
		opts = append(opts, Database(config.Store.Database))
	}
	if len(config.Store.Table) > 0 {
		opts = append(opts, Table(config.Store.Table))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
