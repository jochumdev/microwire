//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	"github.com/google/wire"
	"go-micro.dev/v4"
)

func newService(
	options *Options,
	cliConfig *mCli.Config,
	brokerConfig *mBroker.Config,
	registryConfig *mRegistry.Config,
	transportConfig *mTransport.Config,
) (micro.Service, error) {
	panic(wire.Build(
		DiCliSet,
		DiAllSet,
		DiNoDiSet,
	))
}
