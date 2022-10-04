// Code generated with jinja2 templates. DO NOT EDIT.

package client

import (
	"fmt"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/transport"
	"time"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct{}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin             = "client"
	cliArgPoolSize           = "client_pool_size"
	cliArgPoolTTL            = "client_pool_ttl"
	cliArgPoolRequestTimeout = "client_request_timeout"
	cliArgPoolRetries        = "client_retries"
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
		cli.Usage("Client for go-micro, eg: rpc"),
		cli.Default(config.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolSize)),
		cli.Usage("Sets the client connection pool size"),
		cli.Default(config.PoolSize),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPoolSize)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolTTL)),
		cli.Usage("Sets the client connection pool ttl, e.g: 500ms, 5s, 1m"),
		cli.Default(config.PoolTTL),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPoolTTL)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolRequestTimeout)),
		cli.Usage("Sets the client request timeout, e.g: 500ms, 5s, 1m"),
		cli.Default(config.PoolRequestTimeout),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPoolRequestTimeout)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolRetries)),
		cli.Usage("Sets the client retries"),
		cli.Default(config.PoolRetries),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPoolRetries)),
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
	cfg := sourceConfig{Client: *defConfig}

	if configor != nil {
		if err := configor.Scan(&cfg); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&cfg.Client); err != nil {
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
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolSize)); ok {
		defConfig.PoolSize = cli.FlagValue(f, defConfig.PoolSize)
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolTTL)); ok {
		defConfig.PoolTTL = cli.FlagValue(f, defConfig.PoolTTL)
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolRequestTimeout)); ok {
		defConfig.PoolRequestTimeout = cli.FlagValue(f, defConfig.PoolRequestTimeout)
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgPoolRetries)); ok {
		defConfig.PoolRetries = cli.FlagValue(f, defConfig.PoolRetries)
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
	c := sourceConfig{Client: *defConfig}

	if configor != nil {
		if err := configor.Scan(&c); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&c.Client); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,
	broker broker.Broker,
	log logger.Logger,
	registry registry.Registry,
	transport transport.Transport,
	config *Config,
) (Client, error) {
	if !config.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	pluginFunc, err := Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown plugin client: %s", config.Plugin)
	}

	opts := []Option{WithConfig(config)}
	opts = append(opts, PoolSize(config.PoolSize))
	d, err := time.ParseDuration(config.PoolTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client_pool_ttl: %v", config.PoolTTL)
	}
	opts = append(opts, RequestTimeout(d))
	d, err = time.ParseDuration(config.PoolRequestTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client_request_timeout: %v", config.PoolRequestTimeout)
	}

	if config.Logger.Enabled {
		loggerFunc, err := logger.Plugins.Get(config.Logger.Plugin)
		if err != nil {
			return nil, fmt.Errorf("{{Name}} unknown logger: %s", config.Logger.Plugin)
		}
		log = loggerFunc(logger.ConfigToOpts(config.Logger)...)
	}

	opts = append(
		opts,
		PoolTTL(d),
		Retries(config.PoolRetries),
		Broker(broker),
		Registry(registry),
		Transport(transport),
		WrapCall(config.WrapCall...),
		WithLogger(log),
	)

	return pluginFunc(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
