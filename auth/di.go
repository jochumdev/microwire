// Code generated with jinja2 templates. DO NOT EDIT.

package auth

import (
	"fmt"

	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/di"
	"github.com/google/wire"
)

type DiFlags struct {
	Plugin     string
	ID         string
	Secret     string
	PublicKey  string
	PrivateKey string
	Namespace  string
}

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
) (*DiFlags, error) {
	if cliConfig.Cli.NoFlags {
		// Defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		cli.Usage("Auth for role based access control, e.g. service"),
		cli.Default(config.Auth.Plugin),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPlugin)),
		cli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgID)),
		cli.Usage("Account ID used for client authentication"),
		cli.Default(config.Auth.ID),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgID)),
		cli.Destination(&result.ID),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgSecret)),
		cli.Usage("Account secret used for client authentication"),
		cli.Default(config.Auth.Secret),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgSecret)),
		cli.Destination(&result.Secret),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPublicKey)),
		cli.Usage("Public key for JWT auth (base64 encoded PEM)"),
		cli.Default(config.Auth.PublicKey),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPublicKey)),
		cli.Destination(&result.PublicKey),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgPrivateKey)),
		cli.Usage("Private key for JWT auth (base64 encoded PEM)"),
		cli.Default(config.Auth.PrivateKey),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgPrivateKey)),
		cli.Destination(&result.PrivateKey),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		cli.Name(cli.PrefixName(cliConfig.Cli.ArgPrefix, cliArgNamespace)),
		cli.Usage("Namespace for the services auth account"),
		cli.Default(config.Auth.Namespace),
		cli.EnvVars(cli.PrefixEnv(cliConfig.Cli.ArgPrefix, cliArgNamespace)),
		cli.Destination(&result.Namespace),
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
	defConfig.Auth.Plugin = flags.Plugin

	if len(flags.ID) > 0 && len(flags.Secret) > 0 {
		defConfig.Auth.ID = flags.ID
		defConfig.Auth.Secret = flags.Secret
	}
	if len(flags.PublicKey) > 0 && len(flags.PrivateKey) > 0 {
		defConfig.Auth.PublicKey = flags.PublicKey
		defConfig.Auth.PrivateKey = flags.PrivateKey
	}
	defConfig.Auth.Namespace = flags.Namespace
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
