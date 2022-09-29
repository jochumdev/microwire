//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package microwire

import (
	"github.com/google/wire"
	"go-micro.dev/v4"
)

func NewService(opts ...Option) (micro.Service, error) {
	panic(wire.Build(
		ProvideOptions,
		ProvideConfigStore,
		ProvideDefaultConfigStore,
		DiCliSet,
		DiConfigStagesSet,
		DiAllComponentsSuperSet,
		DiAllComponentProvidersSet,
		DiMicroServiceSet,
	))
}

func NewServiceWithConfigStore(
	config ConfigStore,
	opts ...Option,
) (micro.Service, error) {
	panic(wire.Build(
		ProvideOptions,
		ProvideDefaultConfigStore,
		DiCliSet,
		DiConfigStagesSet,
		DiAllComponentsSuperSet,
		DiAllComponentProvidersSet,
		DiMicroServiceSet,
	))
}
