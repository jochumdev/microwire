package registry

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	"github.com/google/wire"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/cmd"
)

type DiFlags struct {
	Plugin    string
	Addresses string
}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin    = "registry"
	cliArgAddresses = "registry_address"
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
		mCli.Usage("Registry for discovery. etcd, mdns"),
		mCli.Default(config.Registry.Plugin),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgAddresses)),
		mCli.Usage("Comma-separated list of registry addresses"),
		mCli.Default(strings.Join(config.Registry.Addresses, ",")),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgAddresses)),
		mCli.Destination(&result.Addresses),
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
	defConfig.Registry.Plugin = flags.Plugin
	defConfig.Registry.Addresses = strings.Split(flags.Addresses, ",")
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,

	config *Config,
) (registry.Registry, error) {
	if !config.Registry.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Registry.Plugin)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultRegistries[config.Registry.Plugin]; !ok {
			return nil, fmt.Errorf("unknown registry: %v", err)
		}
	}

	opts := []registry.Option{}
	if len(config.Registry.Addresses) > 0 {
		opts = append(opts, registry.Addrs(config.Registry.Addresses...))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
