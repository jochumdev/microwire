// Code generated with jinja2 templates. DO NOT EDIT.

package client

import (
	"fmt"
	"github.com/go-micro/microwire/v5/broker"
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
	if cliConfig.Cli.NoFlags {
		// Defined silently ignore that
		return DiFlags{}, nil
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		cli.Usage("Client for go-micro, eg: rpc"),
		cli.Default(config.Client.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPoolSize)),
		cli.Usage("Sets the client connection pool size"),
		cli.Default(config.Client.PoolSize),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPoolSize)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPoolTTL)),
		cli.Usage("Sets the client connection pool ttl, e.g: 500ms, 5s, 1m"),
		cli.Default(config.Client.PoolTTL),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPoolTTL)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPoolRequestTimeout)),
		cli.Usage("Sets the client request timeout, e.g: 500ms, 5s, 1m"),
		cli.Default(config.Client.PoolRequestTimeout),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPoolRequestTimeout)),
	); err != nil {
		return DiFlags{}, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPoolRetries)),
		cli.Usage("Sets the client retries"),
		cli.Default(config.Client.PoolRetries),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPoolRetries)),
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

	if configor != nil {
		if err := configor.Scan(defConfig); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	if cliConfig.Cli.NoFlags {
		// Dont parse flags if NoFlags has been given
		return DiConfig{}, nil
	}

	defConfig = NewConfig()
	if f, ok := c.Get(cliArgPlugin); ok {
		defConfig.Client.Plugin = cli.FlagValue(f, defConfig.Client.Plugin)
	}
	if f, ok := c.Get(cliArgPoolSize); ok {
		defConfig.Client.PoolSize = cli.FlagValue(f, defConfig.Client.PoolSize)
	}
	if f, ok := c.Get(cliArgPoolTTL); ok {
		defConfig.Client.PoolTTL = cli.FlagValue(f, defConfig.Client.PoolTTL)
	}
	if f, ok := c.Get(cliArgPoolRequestTimeout); ok {
		defConfig.Client.PoolRequestTimeout = cli.FlagValue(f, defConfig.Client.PoolRequestTimeout)
	}
	if f, ok := c.Get(cliArgPoolRetries); ok {
		defConfig.Client.PoolRetries = cli.FlagValue(f, defConfig.Client.PoolRetries)
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

	if configor != nil {
		if err := configor.Scan(defConfig); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(defConfig); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,
	broker broker.Broker,
	registry registry.Registry,
	transport transport.Transport,
	config *Config,
) (Client, error) {
	if !config.Client.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Client.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown client: %v", err)
	}

	opts := []Option{}
	opts = append(opts, PoolSize(config.Client.PoolSize))
	d, err := time.ParseDuration(config.Client.PoolTTL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client_pool_ttl: %v", config.Client.PoolTTL)
	}
	opts = append(opts, RequestTimeout(d))
	d, err = time.ParseDuration(config.Client.PoolRequestTimeout)
	if err != nil {
		return nil, fmt.Errorf("failed to parse client_request_timeout: %v", config.Client.PoolRequestTimeout)
	}

	opts = append(
		opts,
		PoolTTL(d),
		Retries(config.Client.PoolRetries),
		Broker(broker),
		Registry(registry),
		Transport(transport),
		WrapCall(config.Client.WrapCall...),
	)

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
