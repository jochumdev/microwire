// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package micro

import (
	"github.com/go-micro/microwire/v5/auth"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/cache"
	"github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/config/configdi"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
)

// Injectors from wire.go:

func newService(options *Options, cliConfig *cli.Config, authConfig *auth.Config, brokerConfig *broker.Config, cacheConfig *cache.Config, registryConfig *registry.Config, serverConfig *server.Config, storeConfig *store.Config, transportConfig *transport.Config) (Service, error) {
	cliCli, err := cli.ProvideCli(cliConfig)
	if err != nil {
		return nil, err
	}
	diFlags, err := auth.ProvideFlags(authConfig, cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	brokerDiFlags, err := broker.ProvideFlags(brokerConfig, cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	cacheDiFlags, err := cache.ProvideFlags(cacheConfig, cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	registryDiFlags, err := registry.ProvideFlags(registryConfig, cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	serverDiFlags, err := server.ProvideFlags(serverConfig, cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	storeDiFlags, err := store.ProvideFlags(storeConfig, cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	transportDiFlags, err := transport.ProvideFlags(transportConfig, cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	diDiFlags, err := ProvideFlags(diFlags, brokerDiFlags, cacheDiFlags, registryDiFlags, serverDiFlags, storeDiFlags, transportDiFlags)
	if err != nil {
		return nil, err
	}
	diParsed, err := cli.ProvideParsed(cliConfig, cliCli)
	if err != nil {
		return nil, err
	}
	diConfig, err := cli.ProvideConfig(diDiFlags, diParsed, cliConfig)
	if err != nil {
		return nil, err
	}
	config, err := configdi.ProvideConfigor(diConfig)
	if err != nil {
		return nil, err
	}
	authDiConfig, err := auth.ProvideConfig(diConfig, diFlags, authConfig, cliConfig, config)
	if err != nil {
		return nil, err
	}
	authAuth, err := auth.Provide(authDiConfig, authConfig)
	if err != nil {
		return nil, err
	}
	brokerDiConfig, err := broker.ProvideConfig(diConfig, brokerDiFlags, brokerConfig, cliConfig, config)
	if err != nil {
		return nil, err
	}
	brokerBroker, err := broker.Provide(brokerDiConfig, brokerConfig)
	if err != nil {
		return nil, err
	}
	cacheDiConfig, err := cache.ProvideConfig(diConfig, cacheDiFlags, cacheConfig, cliConfig, config)
	if err != nil {
		return nil, err
	}
	cacheCache, err := cache.Provide(cacheDiConfig, cacheConfig)
	if err != nil {
		return nil, err
	}
	registryDiConfig, err := registry.ProvideConfig(diConfig, registryDiFlags, registryConfig, cliConfig, config)
	if err != nil {
		return nil, err
	}
	registryRegistry, err := registry.Provide(registryDiConfig, registryConfig)
	if err != nil {
		return nil, err
	}
	serverDiConfig, err := server.ProvideConfig(diConfig, serverDiFlags, serverConfig, cliConfig, config)
	if err != nil {
		return nil, err
	}
	transportDiConfig, err := transport.ProvideConfig(diConfig, transportDiFlags, transportConfig, cliConfig, config)
	if err != nil {
		return nil, err
	}
	transportTransport, err := transport.Provide(transportDiConfig, transportConfig)
	if err != nil {
		return nil, err
	}
	serverServer, err := server.Provide(serverDiConfig, brokerBroker, registryRegistry, transportTransport, serverConfig)
	if err != nil {
		return nil, err
	}
	storeDiConfig, err := store.ProvideConfig(diConfig, storeDiFlags, storeConfig, cliConfig, config)
	if err != nil {
		return nil, err
	}
	storeStore, err := store.Provide(storeDiConfig, storeConfig)
	if err != nil {
		return nil, err
	}
	microService, err := ProvideAllService(options, authAuth, brokerBroker, cacheCache, registryRegistry, serverServer, storeStore, transportTransport)
	if err != nil {
		return nil, err
	}
	return microService, nil
}
