package cli

import (
	"fmt"
	"os"

	"github.com/go-micro/microwire/v5/di"
)

type DiParsed struct{}

func ProvideCli(
	config *Config,
) (Cli, error) {
	c, err := Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown cli given: %v", err)
	}

	return c(), nil
}

func ProvideParsed(
	config *Config,
	c Cli,
) (DiParsed, error) {
	result := DiParsed{}

	// User flags
	for _, f := range config.Flags {
		if err := c.Add(f.AsOptions()...); err != nil {
			return result, err
		}
	}

	if config.NoFlags {
		if err := c.Add(
			Name(PrefixName(config.ArgPrefix, "config")),
			Usage("Config file"),
			Default(config.ConfigFile),
			EnvVars(PrefixEnv(config.ArgPrefix, "config")),
		); err != nil {
			return result, err
		}
	}

	// Initialize the CLI / parse flags
	if err := c.Parse(
		os.Args,
		CliName(config.Name),
		CliVersion(config.Version),
		CliDescription(config.Description),
		CliUsage(config.Usage),
	); err != nil {
		return result, err
	}

	return result, nil
}

func ProvideConfig(
	_ di.DiFlags,
	diFlags DiParsed,
	c Cli,
	cfg *Config,
) (di.DiConfig, error) {
	if cfg.NoFlags {
		// Defined silently ignore that
		defConfig := NewConfig()
		if f, ok := c.Get("config"); ok {
			defConfig.ConfigFile = FlagValue(f, "")
		}
		if err := cfg.Merge(defConfig); err != nil {
			return "", err
		}
	}

	return di.DiConfig(cfg.ConfigFile), nil
}
