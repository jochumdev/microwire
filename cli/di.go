package cli

import (
	"fmt"
	"os"

	"github.com/go-micro/microwire/v5/di"
)

type DiParsed struct {
	ConfigFile string
}

func ProvideCli(
	config *Config,
) (Cli, error) {
	c, err := Plugins.Get(config.Cli.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown cli given: %v", err)
	}

	return c(), nil
}

func ProvideParsed(
	config *Config,
	c Cli,
) (*DiParsed, error) {
	result := &DiParsed{}

	// User flags
	for _, f := range config.Cli.Flags {
		if err := c.Add(f.AsOptions()...); err != nil {
			return result, err
		}
	}

	if config.Cli.NoFlags {
		if err := c.Add(
			Name(PrefixName(config.Cli.ArgPrefix, "config")),
			Usage("Config file"),
			Default(config.Cli.ConfigFile),
			EnvVars(PrefixEnv(config.Cli.ArgPrefix, "config")),
			Destination(&result.ConfigFile),
		); err != nil {
			return result, err
		}
	}

	// Initialize the CLI / parse flags
	if err := c.Parse(
		os.Args,
		CliName(config.Cli.Name),
		CliVersion(config.Cli.Version),
		CliDescription(config.Cli.Description),
		CliUsage(config.Cli.Usage),
	); err != nil {
		return result, err
	}

	return result, nil
}

func ProvideConfig(
	_ di.DiFlags,
	diFlags *DiParsed,
	cfg *Config,
) (di.DiConfig, error) {
	if cfg.Cli.NoFlags {
		// Defined silently ignore that
		defConfig := NewConfig()
		defConfig.Cli.ConfigFile = diFlags.ConfigFile
		if err := cfg.Merge(defConfig); err != nil {
			return "", err
		}
	}

	return di.DiConfig(cfg.Cli.ConfigFile), nil
}
