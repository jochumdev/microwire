package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	"github.com/google/wire"
)

// DiAllComponentsSet is a set of all things components need, except the components themself.
var DiAllComponentsSet = wire.NewSet(
	mBroker.DiSet,
	mRegistry.DiSet,
	mTransport.DiSet,
)

var DiCliSet = wire.NewSet(
	mCli.ProvideCli,
	mCli.ProvideConfigor,
	mCli.DiSet,
)

var DiNoDiSet = wire.NewSet(
	ProvideFlags,
	ProvideService,
)
