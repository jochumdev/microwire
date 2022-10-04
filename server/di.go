// Code generated with jinja2 templates. DO NOT EDIT.

package server

import (
	"fmt"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/transport"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct {
	Plugin  string
	Address string
	ID      string
	Name    string
	Version string
}

// DiConfig is marker that DiFlags has been parsed into Config
type DiConfig struct{}

const (
	cliArgPlugin  = "server"
	cliArgAddress = "server_address"
	cliArgID      = "server_id"
	cliArgName    = "server_name"
	cliArgVersion = "server_version"
)

func ProvideFlags(
	config *Config,
	cliConfig *cli.Config,
	c cli.Cli,
) (*DiFlags, error) {
	if cliConfig.Cli.NoFlags {
		// Defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		cli.Usage("Server for go-micro; rpc"),
		cli.Default(config.Server.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		cli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgAddress)),
		cli.Usage("Bind address for the server, eg: 127.0.0.1:8080"),
		cli.Default(config.Server.Address),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgAddress)),
		cli.Destination(&result.Address),
	); err != nil {
		return nil, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgID)),
		cli.Usage("Id of the server. Auto-generated if not specified"),
		cli.Default(config.Server.ID),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgID)),
		cli.Destination(&result.ID),
	); err != nil {
		return nil, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgName)),
		cli.Usage("Name of the server. go.micro.srv.example"),
		cli.Default(config.Server.Name),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgName)),
		cli.Destination(&result.Name),
	); err != nil {
		return nil, err
	}
	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgVersion)),
		cli.Usage("Version of the server. 1.1.0"),
		cli.Default(config.Server.Version),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgVersion)),
		cli.Destination(&result.Version),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideConfig(
	_ di.DiConfig,
	flags *DiFlags,
	config *Config,
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
	defConfig.Server.Plugin = flags.Plugin

	defConfig.Server.Address = flags.Address
	defConfig.Server.ID = flags.ID
	defConfig.Server.Name = flags.Name
	defConfig.Server.Version = flags.Version
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
) (Server, error) {
	if !config.Server.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Server.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown server: %v", err)
	}

	opts := []Option{}
	if len(config.Server.Address) > 0 {
		opts = append(opts, Address(config.Server.Address))
	}
	if len(config.Server.ID) > 0 {
		opts = append(opts, Id(config.Server.ID))
	}
	if len(config.Server.Name) > 0 {
		opts = append(opts, Name(config.Server.Name))
	}
	if len(config.Server.Version) > 0 {
		opts = append(opts, Version(config.Server.Version))
	}

	opts = append(opts, Broker(broker), Registry(registry), Transport(transport))

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
