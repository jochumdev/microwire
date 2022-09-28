package broker

import (
	"strings"

	mCli "github.com/go-micro/microwire/cli"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
)

type BrokerOptions struct {
	Name      string
	Addresses string
}

const (
	cliArg        = "broker"
	cliArgAddress = "broker_address"
)

func InjectFlags(opts *mWire.Options, c mCli.CLI) error {
	if err := c.AddString(
		mCli.Name(mCli.PrefixName(opts.ArgPrefix, cliArg)),
		mCli.Usage("Broker for pub/sub. http, nats, rabbitmq"),
		mCli.DefaultValue(opts.Components[mWire.ComponentBroker]),
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

func Inject(opts *mWire.Options, c mWire.InitializedCli) (broker.Broker, error) {
	name := c.String(mCli.PrefixName(opts.ArgPrefix, cliArg))
	addresses := c.String(mCli.PrefixName(opts.ArgPrefix, cliArgAddress))

	b, err := Container.Get(name)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultBrokers[name]; !ok {
			return nil, err
		}
	}

	var result broker.Broker
	if len(addresses) > 0 {
		result = b(broker.Addrs(strings.Split(addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}
