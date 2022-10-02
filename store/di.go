package store

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct {
	Plugin    string
	Addresses string
	Database  string
	Table     string
}

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
	cliConfig *mCli.Config,
	c mCli.Cli,
) (*DiFlags, error) {
	if cliConfig.Cli.NoFlags {
		// Defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		mCli.Usage("Store for pub/sub. http, nats, rabbitmq"),
		mCli.Default(config.Store.Plugin),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		mCli.Usage("Comma-separated list of store addresses"),
		mCli.Default(strings.Join(config.Store.Addresses, ",")),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		mCli.Destination(&result.Addresses),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgDatabase)),
		mCli.Usage("Database option for the underlying store"),
		mCli.Default(config.Store.Database),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgDatabase)),
		mCli.Destination(&result.Database),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgTable)),
		mCli.Usage("Table option for the underlying store"),
		mCli.Default(config.Store.Table),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgTable)),
		mCli.Destination(&result.Table),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideConfig(
	_ di.DiConfig,
	flags *DiFlags,
	config *Config,
	cliConfig *mCli.Config,
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
	defConfig.Store.Plugin = flags.Plugin
	defConfig.Store.Addresses = strings.Split(flags.Addresses, ",")
	defConfig.Store.Database = flags.Database
	defConfig.Store.Table = flags.Table
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
