package registry

import (
	"strings"

	uCmd "github.com/go-micro/microwire/util/cmd"
	"github.com/go-micro/microwire/util/generic"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/cmd"
)

type RegistryFlags []cli.Flag

type RegistryOptions struct {
	Name      string
	Addresses string
}

func ProvideRegistryFlags(opts *mWire.Options) RegistryFlags {
	return RegistryFlags{
		&cli.StringFlag{
			Name:    uCmd.PrefixName(opts.ArgPrefix, "registry"),
			Usage:   "Registry for discovery. etcd, mdns",
			Value:   opts.DefaultRegistry,
			EnvVars: []string{uCmd.PrefixEnv(opts.ArgPrefix, "REGISTRY")},
		},
		&cli.StringFlag{
			Name:    uCmd.PrefixName(opts.ArgPrefix, "registry_address"),
			Usage:   "Comma-separated list of registry addresses",
			EnvVars: []string{uCmd.PrefixEnv(opts.ArgPrefix, "REGISTRY_ADDRESS")},
		},
	}
}

func ProvideRegistryOptions(opts *mWire.Options, c *cli.Context) *RegistryOptions {
	return &RegistryOptions{
		Name:      c.String(uCmd.PrefixName(opts.ArgPrefix, "registry")),
		Addresses: c.String(uCmd.PrefixName(opts.ArgPrefix, "registry_addresses")),
	}
}

func Provide(opts *RegistryOptions) (registry.Registry, error) {
	b, err := Container.Get(opts.Name)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultRegistries[opts.Name]; !ok {
			return nil, err
		}
	}

	var result registry.Registry
	if len(opts.Addresses) > 0 {
		result = b(registry.Addrs(strings.Split(opts.Addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}

var Container = generic.NewContainer(func(opts ...registry.Option) registry.Registry { return nil })
var RegistryServiceSet = wire.NewSet(ProvideRegistryOptions, Provide)
