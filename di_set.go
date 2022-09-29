package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	"github.com/google/wire"
)

var DiBrokerSet = wire.NewSet(
	ProvideBrokerConfigStore,
	mBroker.DiSet,
)

var DiRegistrySet = wire.NewSet(
	ProvideRegistryConfigStore,
	mRegistry.DiSet,
)

var DiTransportSet = wire.NewSet(
	ProvideTransportConfigStore,
	mTransport.DiSet,
)

// DiAllComponentsSuperSet is a set of all things components need, except the components themself.
// Components have been excluded so users of go-micro have another way to filter them out.
var DiAllComponentsSuperSet = wire.NewSet(
	DiBrokerSet,
	DiRegistrySet,
	DiTransportSet,
)

var DiAllComponentProvidersSet = wire.NewSet(
	mBroker.Provide,
	mRegistry.Provide,
	mTransport.Provide,
)

var DiCliSet = wire.NewSet(
	ProvideCliConfigStore,
	ProvideCLI,
	ProvideCliArgs,
	ProvideInitializedCLI,
	mCli.DiSet,
)

var DiMicroServiceSet = wire.NewSet(
	ProvideMicroOpts,
	ProvideMicroService,
)
