package registry

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
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
	config *ConfigStore,
	cliConfig *mCli.ConfigStore,
	c mCli.CLI,
) (*DiFlags, error) {
	if cliConfig.NoFlags {
		// Defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgPlugin)),
		mCli.Usage("Registry for discovery. etcd, mdns"),
		mCli.Default(config.Plugin),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgAddresses)),
		mCli.Usage("Comma-separated list of registry addresses"),
		mCli.Default(strings.Join(config.Addresses, ",")),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgAddresses)),
		mCli.Destination(&result.Addresses),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideDiConfig(
	// Stage2Config must have been populated before
	_ mWire.DiStage2ConfigStore,

	diFlags *DiFlags,
	cliConfig *mCli.ConfigStore,
	config *ConfigStore,
) (DiConfig, error) {
	if cliConfig.NoFlags {
		// Defined silently ignore that
		return DiConfig{}, nil
	}

	defCfg := NewConfigStore()
	defCfg.Plugin = diFlags.Plugin
	defCfg.Addresses = strings.Split(diFlags.Addresses, ",")
	if err := config.Merge(&defCfg); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// We want config at Stage3 (compile->files->flags|env)
	_ mWire.DiStage3ConfigStore,

	config *ConfigStore,

	// Marker so cli has been merged into Config
	_ DiConfig,
) (registry.Registry, error) {
	if !config.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Plugin)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultRegistries[config.Plugin]; !ok {
			return nil, fmt.Errorf("unknown registry: %v", err)
		}
	}

	opts := []registry.Option{}
	if len(config.Addresses) > 0 {
		opts = append(opts, registry.Addrs(config.Addresses...))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideDiConfig, Provide)
