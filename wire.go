//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package microwire

import (
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"go-micro.dev/v4"
)

func DefaultApp(opts ...mWire.Option) (micro.Service, error) {
	panic(wire.Build(
		mWire.ProvideOptions,
		ProvideCLI,
		ProvideCliArgs,
		ProvideInitializedCLI,
		ProvideMicroOpts,
		ProvideMicroService,
	))
}
