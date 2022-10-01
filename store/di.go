package store

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	"github.com/google/wire"
	"go-micro.dev/v4/store"
	"go-micro.dev/v4/util/cmd"
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
	if cliConfig.NoFlags {
		// Defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgPlugin)),
		mCli.Usage("Store for pub/sub. http, nats, rabbitmq"),
		mCli.Default(config.Store.Plugin),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgAddresses)),
		mCli.Usage("Comma-separated list of store addresses"),
		mCli.Default(strings.Join(config.Store.Addresses, ",")),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgAddresses)),
		mCli.Destination(&result.Addresses),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgDatabase)),
		mCli.Usage("Database option for the underlying store"),
		mCli.Default(config.Store.Database),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgDatabase)),
		mCli.Destination(&result.Database),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgTable)),
		mCli.Usage("Table option for the underlying store"),
		mCli.Default(config.Store.Table),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgTable)),
		mCli.Destination(&result.Table),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideConfig(
	flags *DiFlags,
	config *Config,
	configor mCli.DiConfigor,
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

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,

	config *Config,
) (store.Store, error) {
	if !config.Store.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Store.Plugin)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultStores[config.Store.Plugin]; !ok {
			return nil, fmt.Errorf("unknown store: %v", err)
		}
	}

	opts := []store.Option{}
	if len(config.Store.Addresses) > 0 {
		opts = append(opts, store.Nodes(config.Store.Addresses...))
	}
	if len(config.Store.Database) > 0 {
		opts = append(opts, store.Database(config.Store.Database))
	}
	if len(config.Store.Table) > 0 {
		opts = append(opts, store.Table(config.Store.Table))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
