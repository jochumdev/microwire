//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
)

func DefaultApp(opts ...mWire.Option) (*cli.App, error) {
	panic(wire.Build(
		ProvideDefaultServiceInitializer,
		mWire.ProvideOptions,
		mBroker.ProvideFlags,
		mRegistry.ProvideFlags,
		mTransport.ProvideFlags,
		ProvideDefaultFlags,
		mWire.ProvideApp,
	))
}

func DefaultService(ctx *cli.Context, opts *mWire.Options) (micro.Service, error) {
	panic(wire.Build(
		mBroker.BrokerServiceSet,
		mRegistry.RegistryServiceSet,
		mTransport.TransportServiceSet,
		ProvideDefaultMicroOpts,
		mWire.ProvideMicroService,
	))
}
