package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/transport"
)

func ProvideDefaultServiceInitializer() mWire.InitializeServiceFunc {
	return DefaultService
}

func ProvideDefaultFlags(opts *mWire.Options, brokerFlags mBroker.BrokerFlags, registryFlags mRegistry.RegistryFlags, transportFlags mTransport.TransportFlags) mWire.InternalFlags {
	result := []cli.Flag{}
	result = append(result, []cli.Flag(brokerFlags)...)
	result = append(result, []cli.Flag(registryFlags)...)
	result = append(result, []cli.Flag(transportFlags)...)
	result = append(result, opts.Flags...)
	return mWire.InternalFlags(result)
}

func ProvideDefaultMicroOpts(opts *mWire.Options, broker broker.Broker, registry registry.Registry, transport transport.Transport) []micro.Option {
	result := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
		micro.Registry(registry),
		micro.Broker(broker),
		micro.Transport(transport),
	}

	for _, fn := range opts.BeforeStart {
		result = append(result, micro.BeforeStart(fn))
	}
	for _, fn := range opts.BeforeStop {
		result = append(result, micro.BeforeStop(fn))
	}
	for _, fn := range opts.AfterStart {
		result = append(result, micro.AfterStart(fn))
	}
	for _, fn := range opts.AfterStop {
		result = append(result, micro.AfterStop(fn))
	}

	return result
}
