package microwire

import (
	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/util/cmd"
	"github.com/urfave/cli/v2"

	"github.com/google/wire"
)

type BrokerFlags []cli.Flag

func ProvideBrokerFlags(opts *Options) BrokerFlags {
	return BrokerFlags{
		&cli.StringFlag{
			Name:    cmd.PrefixName(opts.ArgPrefix, "broker"),
			Usage:   "Broker for pub/sub. http, nats, rabbitmq",
			Value:   "http",
			EnvVars: []string{cmd.PrefixEnv(opts.ArgPrefix, "BROKER")},
		},
		&cli.StringFlag{
			Name:    cmd.PrefixName(opts.ArgPrefix, "broker_address"),
			Usage:   "Comma-separated list of broker addresses",
			EnvVars: []string{cmd.PrefixEnv(opts.ArgPrefix, "BROKER_ADDRESS")},
		},
	}
}

func ProvideBrokerOptions(opts *Options, c *cli.Context) *broker.BrokerOptions {
	return &broker.BrokerOptions{
		Name:      c.String(cmd.PrefixName(opts.ArgPrefix, "broker")),
		Addresses: c.String(cmd.PrefixName(opts.ArgPrefix, "broker_addresses")),
	}
}

var BrokerServiceSet = wire.NewSet(ProvideBrokerOptions, broker.ProvideBroker)
