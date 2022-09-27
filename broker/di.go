package broker

import (
	"strings"

	uCmd "github.com/go-micro/microwire/util/cmd"
	"github.com/go-micro/microwire/util/generic"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
)

type BrokerFlags []cli.Flag

type BrokerOptions struct {
	Name      string
	Addresses string
}

func ProvideBrokerFlags(opts *mWire.Options) BrokerFlags {
	return BrokerFlags{
		&cli.StringFlag{
			Name:    uCmd.PrefixName(opts.ArgPrefix, "broker"),
			Usage:   "Broker for pub/sub. http, nats, rabbitmq",
			Value:   opts.DefaultBroker,
			EnvVars: []string{uCmd.PrefixEnv(opts.ArgPrefix, "BROKER")},
		},
		&cli.StringFlag{
			Name:    uCmd.PrefixName(opts.ArgPrefix, "broker_address"),
			Usage:   "Comma-separated list of broker addresses",
			EnvVars: []string{uCmd.PrefixEnv(opts.ArgPrefix, "BROKER_ADDRESS")},
		},
	}
}

func ProvideBrokerOptions(opts *mWire.Options, c *cli.Context) *BrokerOptions {
	return &BrokerOptions{
		Name:      c.String(uCmd.PrefixName(opts.ArgPrefix, "broker")),
		Addresses: c.String(uCmd.PrefixName(opts.ArgPrefix, "broker_addresses")),
	}
}

func Provide(opts *BrokerOptions) (broker.Broker, error) {
	b, err := Container.Get(opts.Name)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultBrokers[opts.Name]; !ok {
			return nil, err
		}
	}

	var result broker.Broker
	if len(opts.Addresses) > 0 {
		result = b(broker.Addrs(strings.Split(opts.Addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}

var BrokerServiceSet = wire.NewSet(ProvideBrokerOptions, Provide)
var Container = generic.NewContainer(func(opts ...broker.Option) broker.Broker { return nil })
