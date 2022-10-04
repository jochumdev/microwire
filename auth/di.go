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
	if cliConfig.NoFlags {
		// Defined silently ignore that
		return DiFlags{}, nil
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPlugin)),
		cli.Usage("Auth for role based access control, e.g. service"),
		cli.Default(config.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPlugin)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgID)),
		cli.Usage("Account ID used for client authentication"),
		cli.Default(config.ID),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgID)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgSecret)),
		cli.Usage("Account secret used for client authentication"),
		cli.Default(config.Secret),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgSecret)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPublicKey)),
		cli.Usage("Public key for JWT auth (base64 encoded PEM)"),
		cli.Default(config.PublicKey),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPublicKey)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgPrivateKey)),
		cli.Usage("Private key for JWT auth (base64 encoded PEM)"),
		cli.Default(config.PrivateKey),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgPrivateKey)),
	); err != nil {
		return DiFlags{}, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.ArgPrefix, cliArgNamespace)),
		cli.Usage("Namespace for the services auth account"),
		cli.Default(config.Namespace),
		cli.EnvVars(cli.PrefixEnv(cliConfig.ArgPrefix, cliArgNamespace)),
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
	cfg := sourceConfig{Auth: *defConfig}

	if configor != nil {
		if err := configor.Scan(&cfg); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&cfg.Auth); err != nil {
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
	f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgID))
	f2, ok2 := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgSecret))
	if ok && ok2 {
		if len(cli.FlagValue(f, defConfig.ID)) > 0 && len(cli.FlagValue(f2, defConfig.Secret)) > 0 {
			defConfig.ID = cli.FlagValue(f, "")
			defConfig.Secret = cli.FlagValue(f2, "")
		}
	}
	f, ok = c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgPublicKey))
	f2, ok2 = c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgPrivateKey))
	if ok && ok2 {
		if len(cli.FlagValue(f, defConfig.PublicKey)) > 0 && len(cli.FlagValue(f2, defConfig.PrivateKey)) > 0 {
			defConfig.PublicKey = cli.FlagValue(f, "")
			defConfig.PrivateKey = cli.FlagValue(f2, "")
		}
	}
	if f, ok := c.Get(cli.PrefixName(cliConfig.ArgPrefix, cliArgNamespace)); ok {
		defConfig.Namespace = cli.FlagValue(f, "")
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
	c := sourceConfig{Auth: *defConfig}

	if configor != nil {
		if err := configor.Scan(&c); err != nil {
			return DiConfig{}, err
		}
	}
	if err := config.Merge(&c.Auth); err != nil {
		return DiConfig{}, err
	}

	return DiConfig{}, nil
}

func Provide(
	// Marker so cli has been merged into Config
	_ DiConfig,

	config *Config,
) (Auth, error) {
	if !config.Enabled {
		// Not enabled silently ignore that
		return nil, nil
	}

	pluginFunc, err := Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown plugin auth: %s", config.Plugin)
	}

	opts := []Option{WithConfig(config)}
	if len(config.ID) > 0 && len(config.Secret) > 0 {
		opts = append(opts, Credentials(
			config.ID, config.Secret,
		))
	}
	opts = append(opts, PublicKey(config.PublicKey))
	opts = append(opts, PrivateKey(config.PrivateKey))
	opts = append(opts, Namespace(config.Namespace))

	return pluginFunc(opts...), nil
}

var DiSet = wire.NewSet(ProvideFlags, ProvideConfig, Provide)
var DiNoCliSet = wire.NewSet(ProvideConfigNoFlags, Provide)
