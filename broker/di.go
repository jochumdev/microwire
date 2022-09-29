package broker

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
)

type DiFlags struct {
	Plugin    string
	Addresses string
}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArg        = "broker"
	cliArgAddress = "broker_address"
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
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArg)),
		mCli.Usage("Broker for pub/sub. http, nats, rabbitmq"),
		mCli.Default(config.Plugin),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArg)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(cliConfig.ArgPrefix, cliArgAddress)),
		mCli.Usage("Comma-separated list of broker addresses"),
		mCli.Default(strings.Join(config.Addresses, ",")),
		mCli.EnvVars(mCli.PrefixEnv(cliConfig.ArgPrefix, cliArgAddress)),
		mCli.Destination(&result.Addresses),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideDiConfig(
	// Marker so cli is parsed before coming here
	_ mCli.ParsedCli,

	diFlags *DiFlags,
	cliConfig *mCli.ConfigStore,
	config *ConfigStore,
) (DiConfig, error) {
	if cliConfig.NoFlags {
		// Defined silently ignore that
		return DiConfig{}, nil
	}

	config.Plugin = diFlags.Plugin
	config.Addresses = strings.Split(diFlags.Addresses, ",")
	return DiConfig{}, nil
}

func Provide(
	// We want config at Stage3 (compile->files->flags|env)
	_ mWire.DiStage3ConfigStore,
	config *ConfigStore,

	// Marker so cli has been merged into Config
	_ DiConfig,
) (broker.Broker, error) {
	if len(config.Plugin) == 0 {
		// Not defined silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Plugin)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultBrokers[config.Plugin]; !ok {
			return nil, fmt.Errorf("unknown broker: %v", err)
		}
	}

	var result broker.Broker
	if len(config.Addresses) > 0 {
		result = b(broker.Addrs(config.Addresses...))
	} else {
		result = b()
	}

	return result, nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideDiConfig)
