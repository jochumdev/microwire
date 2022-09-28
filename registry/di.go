package registry

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/cmd"
)

type DiFlags struct{}

type DiOptions struct {
	Plugin    string
	Addresses string
}

const (
	cliArg        = "registry"
	cliArgAddress = "registry_address"
)

func ProvideFlags(opts *mWire.Options, c mCli.CLI) (*DiFlags, error) {
	if _, ok := opts.Components[ComponentName]; !ok {
		// Not defined silently ignore that
		return nil, nil
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArg)),
		mCli.Usage("Registry for discovery. etcd, mdns"),
		mCli.Default(opts.Components[ComponentName]),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArg)),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArgAddress)),
		mCli.Usage("Comma-separated list of registry addresses"),
		mCli.Default(""),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArgAddress)),
	); err != nil {
		return nil, err
	}

	return &DiFlags{}, nil
}

func ProvideOpts(opts *mWire.Options, c mWire.InitializedCli) (*DiOptions, error) {
	if _, ok := opts.Components[ComponentName]; !ok {
		// Not defined silently ignore that
		return nil, nil
	}

	return &DiOptions{
		Plugin:    c.StringValue(mCli.PrefixName(opts.ArgPrefix, cliArg)),
		Addresses: c.StringValue(mCli.PrefixName(opts.ArgPrefix, cliArgAddress)),
	}, nil
}

func Provide(opts *mWire.Options, diOpts *DiOptions) (registry.Registry, error) {
	if _, ok := opts.Components[ComponentName]; !ok {
		// Not defined silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(diOpts.Plugin)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultRegistries[diOpts.Plugin]; !ok {
			return nil, fmt.Errorf("unknown registry: %v", err)
		}
	}

	var result registry.Registry
	if len(diOpts.Addresses) > 0 {
		result = b(registry.Addrs(strings.Split(diOpts.Addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}

// Provide is not in here as we always need to call it
var Set = wire.NewSet(ProvideFlags, ProvideOpts)
