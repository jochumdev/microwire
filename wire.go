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
	"go-micro.dev/v4"
)

func NewWireService(opts ...mWire.Option) (micro.Service, error) {
	panic(wire.Build(
		ProvideOptions,
		ProvideCLI,
		ProvideCliArgs,
		mBroker.Set,
		mRegistry.Set,
		mTransport.Set,
		AllComponentsSet,
		ProvideInitializedCLI,
		ProvideMicroOpts,
		ProvideMicroService,
	))
}
