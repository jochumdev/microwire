package transport

import (
	"strings"

	mCmd "github.com/go-micro/microwire/util/cmd"
	"github.com/go-micro/microwire/util/generic"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/transport"
	"go-micro.dev/v4/util/cmd"

	"github.com/google/wire"
)

type TransportFlags []cli.Flag

type TransportOptions struct {
	Name      string
	Addresses string
}

func ProvideFlags(opts *mWire.Options) TransportFlags {
	return TransportFlags{
		&cli.StringFlag{
			Name:    mCmd.PrefixName(opts.ArgPrefix, "transport"),
			Usage:   "Transport for pub/sub. http, nats, rabbitmq",
			Value:   opts.DefaultTransport,
			EnvVars: []string{mCmd.PrefixEnv(opts.ArgPrefix, "BROKER")},
		},
		&cli.StringFlag{
			Name:    mCmd.PrefixName(opts.ArgPrefix, "transport_address"),
			Usage:   "Comma-separated list of transport addresses",
			EnvVars: []string{mCmd.PrefixEnv(opts.ArgPrefix, "BROKER_ADDRESS")},
		},
	}
}

func ProvideOptions(opts *mWire.Options, c *cli.Context) *TransportOptions {
	return &TransportOptions{
		Name:      c.String(mCmd.PrefixName(opts.ArgPrefix, "transport")),
		Addresses: c.String(mCmd.PrefixName(opts.ArgPrefix, "transport_addresses")),
	}
}

func Provide(opts *TransportOptions) (transport.Transport, error) {
	b, err := Container.Get(opts.Name)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultTransports[opts.Name]; !ok {
			return nil, err
		}
	}

	var result transport.Transport
	if len(opts.Addresses) > 0 {
		result = b(transport.Addrs(strings.Split(opts.Addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}

var TransportServiceSet = wire.NewSet(ProvideOptions, Provide)
var Container = generic.NewContainer(func(opts ...transport.Option) transport.Transport { return nil })
