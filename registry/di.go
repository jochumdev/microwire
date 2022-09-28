package registry

import (
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/cmd"
)

type RegistryFlags bool

type RegistryOptions struct {
	Name      string
	Addresses string
}

const (
	cliArg        = "registry"
	cliArgAddress = "registry_address"
)

func InjectFlags(opts *mWire.Options, c mCli.CLI) error {
	if err := c.AddString(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArg)),
		mCli.Usage("Registry for discovery. etcd, mdns"),
		mCli.DefaultValue(opts.Components[mWire.ComponentRegistry]),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArg)),
	); err != nil {
		return err
	}

	if err := c.AddString(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArgAddress)),
		mCli.Usage("Comma-separated list of registry addresses"),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArgAddress)),
	); err != nil {
		return err
	}

	return nil
}

func Inject(opts *mWire.Options, c mWire.InitializedCli) (registry.Registry, error) {
	name := c.String(mCli.PrefixName(opts.ArgPrefix, cliArg))
	addresses := c.String(mCli.PrefixName(opts.ArgPrefix, cliArgAddress))

	b, err := Container.Get(name)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultRegistries[name]; !ok {
			return nil, err
		}
	}

	var result registry.Registry
	if len(addresses) > 0 {
		result = b(registry.Addrs(strings.Split(addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}
