// Code generated with jinja2 templates. DO NOT EDIT.

package store

import (
	"fmt"
	"github.com/go-micro/microwire/v5/logger"

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
	if cliConfig.NoFlags {
		// Defined silently ignore that
		return DiFlags{}, nil
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPlugin)),
		cli.Usage("Store for pub/sub. http, nats, rabbitmq"),
		cli.Default(config.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgAddresses)),
		cli.Usage("List of store addresses"),
		cli.Default(config.Addresses),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgAddresses)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgDatabase)),
		cli.Usage("Database option for the underlying store"),
		cli.Default(config.Database),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgDatabase)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgTable)),
		cli.Usage("Table option for the underlying store"),
		cli.Default(config.Table),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgTable)),
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
	cfg := sourceConfig{Store: *defConfig}

	if configor != nil {
		if err := configor.Scan(&cfg); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&cfg.Store); err != nil {
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
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgAddresses)); ok {
		defConfig.Addresses = cli.FlagValue(f, []string{})
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgDatabase)); ok {
		defConfig.Database = cli.FlagValue(f, "")
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgTable)); ok {
		defConfig.Table = cli.FlagValue(f, "")
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
	c := sourceConfig{Store: *defConfig}

	if configor != nil {
		if err := configor.Scan(&c); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&c.Store); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,
	logger logger.Logger,
	config *Config,
) (Store, error) {
	if !config.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	pluginFunc, err := Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown plugin store: %s", config.Plugin)
	}

	opts := []Option{WithConfig(config)}
	if len(config.Addresses) > 0 {
		opts = append(opts, Nodes(config.Addresses...))
	}
	if len(config.Database) > 0 {
		opts = append(opts, Database(config.Database))
	}
	if len(config.Table) > 0 {
		opts = append(opts, Table(config.Table))
	}
	opts = append(opts, WithLogger(logger))

	return pluginFunc(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
