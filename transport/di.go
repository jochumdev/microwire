package transport

import (
	"fmt"
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/util/cmd"
)

type DiFlags struct {
	Plugin    string
	Addresses string
}

type DiOptions DiFlags

const (
	cliArg        = "transport"
	cliArgAddress = "transport_address"
)

func ProvideFlags(opts *mWire.Options, c mCli.CLI) (*DiFlags, error) {
	if _, ok := opts.Components[ComponentName]; !ok {
		// Not defined silently ignore that
		return &DiFlags{}, nil
	}

	result := &DiFlags{}
	if err := c.Add(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArg)),
		mCli.Usage("Transport for pub/sub. http, nats, rabbitmq"),
		mCli.Default(opts.Components[ComponentName]),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArg)),
		mCli.Destination(&result.Plugin),
	); err != nil {
		return nil, err
	}

	if err := c.Add(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArgAddress)),
		mCli.Usage("Comma-separated list of broker addresses"),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArgAddress)),
		mCli.Destination(&result.Addresses),
	); err != nil {
		return nil, err
	}

	return result, nil
}

func ProvideOpts(diOpts *DiFlags, _ mWire.InitializedCli) (*DiOptions, error) {
	return &DiOptions{
		Plugin:    diOpts.Plugin,
		Addresses: diOpts.Addresses,
	}, nil
}

func Provide(opts *mWire.Options, diOpts *DiOptions) (transport.Transport, error) {
	if _, ok := opts.Components[ComponentName]; !ok {
		// Not defined silently ignore that
		return nil, nil
	}

	b, err := Plugins.Get(diOpts.Plugin)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultTransports[diOpts.Plugin]; !ok {
			return nil, fmt.Errorf("unknown transport: %v", err)
		}
	}

	var result transport.Transport
	if len(diOpts.Addresses) > 0 {
		result = b(transport.Addrs(strings.Split(diOpts.Addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}

var Set = wire.NewSet(ProvideFlags, ProvideOpts)
