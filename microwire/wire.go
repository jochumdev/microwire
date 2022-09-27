//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package microwire

import (
	"github.com/google/wire"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
)

func DefaultApp(opts ...Option) (*cli.App, error) {
	panic(wire.Build(
		ProvideDefaultServiceInitializer,
		ProvideOptions,
		ProvideBrokerFlags,
		ProvideDefaultFlags,
		ProvideApp,
	))
}

func DefaultService(ctx *cli.Context, opts *Options) (micro.Service, error) {
	panic(wire.Build(
		BrokerServiceSet,
		ProvideDefaultMicroOpts,
		ProvideMicroService,
	))
}
