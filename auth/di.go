// Code generated with jinja2 templates. DO NOT EDIT.

package auth

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
	cliArgPlugin     = "auth"
	cliArgID         = "auth_id"
	cliArgSecret     = "auth_secret"
	cliArgPublicKey  = "auth_public_key"
	cliArgPrivateKey = "auth_private_key"
	cliArgNamespace  = "auth_namespace"
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
		cli.Usage("Auth for role based access control, e.g. service"),
		cli.Default(config.Auth.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgID)),
		cli.Usage("Account ID used for client authentication"),
		cli.Default(config.Auth.ID),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgID)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgSecret)),
		cli.Usage("Account secret used for client authentication"),
		cli.Default(config.Auth.Secret),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgSecret)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPublicKey)),
		cli.Usage("Public key for JWT auth (base64 encoded PEM)"),
		cli.Default(config.Auth.PublicKey),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPublicKey)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPrivateKey)),
		cli.Usage("Private key for JWT auth (base64 encoded PEM)"),
		cli.Default(config.Auth.PrivateKey),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPrivateKey)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgNamespace)),
		cli.Usage("Namespace for the services auth account"),
		cli.Default(config.Auth.Namespace),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgNamespace)),
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
		defConfig.Auth.Plugin = cli.FlagValue(f, defConfig.Auth.Plugin)
	}
	f, ok := c.Get(cliArgID)
	f2, ok2 := c.Get(cliArgSecret)
	if ok && ok2 {
		if len(cli.FlagValue(f, defConfig.Auth.ID)) > 0 && len(cli.FlagValue(f2, defConfig.Auth.Secret)) > 0 {
			defConfig.Auth.ID = cli.FlagValue(f, "")
			defConfig.Auth.Secret = cli.FlagValue(f2, "")
		}
	}
	f, ok = c.Get(cliArgPublicKey)
	f2, ok2 = c.Get(cliArgPrivateKey)
	if ok && ok2 {
		if len(cli.FlagValue(f, defConfig.Auth.PublicKey)) > 0 && len(cli.FlagValue(f2, defConfig.Auth.PrivateKey)) > 0 {
			defConfig.Auth.PublicKey = cli.FlagValue(f, "")
			defConfig.Auth.PrivateKey = cli.FlagValue(f2, "")
		}
	}
	if f, ok := c.Get(cliArgNamespace); ok {
		defConfig.Auth.Namespace = cli.FlagValue(f, "")
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

	config *Config,
) (Auth, error) {
	if !config.Auth.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(config.Auth.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown auth: %v", err)
	}

	opts := []Option{}
	if len(config.Auth.ID) > 0 && len(config.Auth.Secret) > 0 {
		opts = append(opts, Credentials(
			config.Auth.ID, config.Auth.Secret,
		))
	}
	opts = append(opts, PublicKey(config.Auth.PublicKey))
	opts = append(opts, PrivateKey(config.Auth.PrivateKey))
	opts = append(opts, Namespace(config.Auth.Namespace))

	return b(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
