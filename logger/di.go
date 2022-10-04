// Code generated with jinja2 templates. DO NOT EDIT.

package logger

import (
	"fmt"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct{}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin      = "logger"
	cliArgLoggerLevel = "logger_level"
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
		cli.Usage("Logger to use"),
		cli.Default(config.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgLoggerLevel)),
		cli.Usage("Log level, e.g.: trace, debug, info, warn, error, fatal"),
		cli.Default(config.Level),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgLoggerLevel)),
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
	cfg := sourceConfig{Logger: *defConfig}

	if configor != nil {
		if err := configor.Scan(&cfg); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&cfg.Logger); err != nil {
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
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgLoggerLevel)); ok {
		defConfig.Level = cli.FlagValue(f, config.Level)
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
	c := sourceConfig{Logger: *defConfig}

	if configor != nil {
		if err := configor.Scan(&c); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&c.Logger); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,

	config *Config,
) (Logger, error) {
	if !config.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	pluginFunc, err := Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown plugin logger: %s", config.Plugin)
	}

	opts := []Option{WithConfig(config)}
	opts = ConfigToOpts(config)

	return pluginFunc(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
