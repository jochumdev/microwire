package transport

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	"github.com/go-micro/microwire/di"
	"github.com/google/wire"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/util/cmd"
)

type DiFlags struct {
	Plugin    string
	Addresses string
}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin    = "transport"
	cliArgAddresses = "transport_address"
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
		mCli.Usage("Transport mechanism used; http"),
		mCli.Default(config.Transport.Plugin),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		mCli.Usage("Comma-separated list of transport addresses"),
		mCli.Default(strings.Join(config.Transport.Addresses, ",")),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgAddresses)),
		mCli.Destination(&result.Addresses),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideConfig(
	flags *DiFlags,
	config *Config,
	configor di.DiConfigor,
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
	defConfig.Transport.Plugin = flags.Plugin
	defConfig.Transport.Addresses = strings.Split(flags.Addresses, ",")
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ di.DiConfig,

	config *Config,
) (transport.Transport, error) {
	if !config.Transport.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Transport.Plugin)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultTransports[config.Transport.Plugin]; !ok {
			return nil, fmt.Errorf("unknown transport: %v", err)
		}
	}

	opts := []transport.Option{}
	if len(config.Transport.Addresses) > 0 {
		opts = append(opts, transport.Addrs(config.Transport.Addresses...))
	}

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
