//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package micro

import (
	mBroker "github.com/go-micro/microwire/v5/broker"
	mCli "github.com/go-micro/microwire/v5/cli"
	mRegistry "github.com/go-micro/microwire/v5/registry"
	mStore "github.com/go-micro/microwire/v5/store"
	mTransport "github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"
)

func newService(
	options *Options,
	cliConfig *mCli.Config,
	brokerConfig *mBroker.Config,
	registryConfig *mRegistry.Config,
	storeConfig *mStore.Config,
	transportConfig *mTransport.Config,
) (Service, error) {
	panic(wire.Build(
		DiCliSet,
		DiSet,
		DiNoDiSet,
	))
}
