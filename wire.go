//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package micro

import (
	"github.com/go-micro/microwire/v5/auth"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/cache"
	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
	"github.com/google/wire"
)

func newService(
	options *Options,
	cliConfig *cli.Config,
	authConfig *auth.Config,
	brokerConfig *broker.Config,
	cacheConfig *cache.Config,
	clientConfig *client.Config,
	loggerConfig *logger.Config,
	registryConfig *registry.Config,
	serverConfig *server.Config,
	storeConfig *store.Config,
	transportConfig *transport.Config,
) (Service, error) {
	panic(wire.Build(
		DiCliSet,
		DiSet,
		DiNoDiSet,
	))
}
