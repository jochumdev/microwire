package transport

import (
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/util/cmd"
)

type TransportFlags bool

type TransportOptions struct {
	Name      string
	Addresses string
}

const (
	cliArg        = "transport"
	cliArgAddress = "transport_address"
)

func InjectFlags(opts *mWire.Options, c mCli.CLI) error {
	if err := c.AddString(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArg)),
		mCli.Usage("Transport for pub/sub. http, nats, rabbitmq"),
		mCli.DefaultValue(opts.Components[mWire.ComponentTransport]),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArg)),
	); err != nil {
		return err
	}

	if err := c.AddString(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArgAddress)),
		mCli.Usage("Comma-separated list of broker addresses"),
		mCli.EnvVars(mCli.PrefixEnv(opts.ArgPrefix, cliArgAddress)),
	); err != nil {
		return err
	}

	return nil
}

func Inject(opts *mWire.Options, c mWire.InitializedCli) (transport.Transport, error) {
	name := c.String(mCli.PrefixName(opts.ArgPrefix, cliArg))
	addresses := c.String(mCli.PrefixName(opts.ArgPrefix, cliArgAddress))

	b, err := Container.Get(name)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultTransports[name]; !ok {
			return nil, err
		}
	}

	var result transport.Transport
	if len(addresses) > 0 {
		result = b(transport.Addrs(strings.Split(addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}
