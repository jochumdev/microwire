package cli

import (
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
)

type DiFlags struct {
	ConfigFile string
}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

// ParsedCli is a marker to let di know that CLI has been Parsed
type ParsedCli CLI

func ProvideFlags(
	config *ConfigStore,
	c CLI,
) (*DiFlags, error) {
	if config.NoFlags {
		// Defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}

	if err := c.Add(
		Name(PrefixName(config.ArgPrefix, "config")),
		Usage("Config file"),
		Default(config.ConfigFile),
		EnvVars(PrefixEnv(config.ArgPrefix, "config")),
		Destination(&result.ConfigFile),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideDiConfig(
	_ mWire.DiStage1ConfigStore,
	_ ParsedCli,

	diFlags *DiFlags,
	config *ConfigStore,
) (DiConfig, error) {
	if config.NoFlags {
		// Defined silently ignore that
		return DiConfig{}, nil
	}

	defConfig := NewConfigStore()
	defConfig.ConfigFile = diFlags.ConfigFile
	if err := config.Merge(&defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideDiConfig)
