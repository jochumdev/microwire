// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package microwire

import (
	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/cli"
	"github.com/go-micro/microwire/registry"
	"github.com/go-micro/microwire/store"
	"github.com/go-micro/microwire/transport"
	"go-micro.dev/v4"
)

// Injectors from wire.go:

func NewService(opts ...Option) (micro.Service, error) {
	configStore, err := ProvideConfigStore()
	if err != nil {
		return nil, err
	}
	options := NewOptions(opts)
	diStage1ConfigStore, err := ProvideStage1ConfigStore(options, configStore)
	if err != nil {
		return nil, err
	}
	brokerConfigStore, err := ProvideBrokerConfigStore(diStage1ConfigStore, configStore)
	if err != nil {
		return nil, err
	}
	cliConfigStore, err := ProvideCliConfigStore(diStage1ConfigStore, configStore)
	if err != nil {
		return nil, err
	}
	cliCLI, err := ProvideCLI(diStage1ConfigStore, cliConfigStore)
	if err != nil {
		return nil, err
	}
	diFlags, err := broker.ProvideFlags(brokerConfigStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	registryConfigStore, err := ProvideRegistryConfigStore(diStage1ConfigStore, configStore)
	if err != nil {
		return nil, err
	}
	registryDiFlags, err := registry.ProvideFlags(registryConfigStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	transportConfigStore, err := ProvideTransportConfigStore(diStage1ConfigStore, configStore)
	if err != nil {
		return nil, err
	}
	transportDiFlags, err := transport.ProvideFlags(transportConfigStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	cliDiFlags, err := cli.ProvideFlags(cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	storeConfigStore, err := ProvideStoreConfigStore(diStage1ConfigStore, configStore)
	if err != nil {
		return nil, err
	}
	storeDiFlags, err := store.ProvideFlags(storeConfigStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	cliArgs := ProvideCliArgs()
	parsedCli, err := ProvideInitializedCLI(diFlags, registryDiFlags, transportDiFlags, cliDiFlags, storeDiFlags, options, cliCLI, cliArgs)
	if err != nil {
		return nil, err
	}
	diConfig, err := cli.ProvideDiConfig(diStage1ConfigStore, parsedCli, cliDiFlags, cliConfigStore)
	if err != nil {
		return nil, err
	}
	microwireDiConfig, err := ProvideConfigFile(diStage1ConfigStore, parsedCli, diConfig, configStore)
	if err != nil {
		return nil, err
	}
	diStage2ConfigStore, err := ProvideStage2ConfigStore(microwireDiConfig, configStore)
	if err != nil {
		return nil, err
	}
	brokerDiConfig, err := broker.ProvideDiConfig(diStage2ConfigStore, diFlags, cliConfigStore, brokerConfigStore)
	if err != nil {
		return nil, err
	}
	registryDiConfig, err := registry.ProvideDiConfig(diStage2ConfigStore, registryDiFlags, cliConfigStore, registryConfigStore)
	if err != nil {
		return nil, err
	}
	transportDiConfig, err := transport.ProvideDiConfig(diStage2ConfigStore, transportDiFlags, cliConfigStore, transportConfigStore)
	if err != nil {
		return nil, err
	}
	diStage3ConfigStore, err := ProvideStage3ConfigStore(diStage2ConfigStore, brokerDiConfig, registryDiConfig, transportDiConfig, configStore)
	if err != nil {
		return nil, err
	}
	brokerBroker, err := broker.Provide(diStage3ConfigStore, brokerConfigStore, brokerDiConfig)
	if err != nil {
		return nil, err
	}
	registryRegistry, err := registry.Provide(diStage3ConfigStore, registryConfigStore, registryDiConfig)
	if err != nil {
		return nil, err
	}
	storeDiConfig, err := store.ProvideDiConfig(diStage2ConfigStore, storeDiFlags, cliConfigStore, storeConfigStore)
	if err != nil {
		return nil, err
	}
	storeStore, err := store.Provide(diStage3ConfigStore, storeConfigStore, storeDiConfig)
	if err != nil {
		return nil, err
	}
	transportTransport, err := transport.Provide(diStage3ConfigStore, transportConfigStore, transportDiConfig)
	if err != nil {
		return nil, err
	}
	v, err := ProvideMicroOpts(options, parsedCli, diStage3ConfigStore, brokerBroker, registryRegistry, storeStore, transportTransport)
	if err != nil {
		return nil, err
	}
	service, err := ProvideMicroService(configStore, options, v)
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewServiceWithConfigStore(config ConfigStore, opts ...Option) (micro.Service, error) {
	options := NewOptions(opts)
	diStage1ConfigStore, err := ProvideStage1ConfigStore(options, config)
	if err != nil {
		return nil, err
	}
	configStore, err := ProvideBrokerConfigStore(diStage1ConfigStore, config)
	if err != nil {
		return nil, err
	}
	cliConfigStore, err := ProvideCliConfigStore(diStage1ConfigStore, config)
	if err != nil {
		return nil, err
	}
	cliCLI, err := ProvideCLI(diStage1ConfigStore, cliConfigStore)
	if err != nil {
		return nil, err
	}
	diFlags, err := broker.ProvideFlags(configStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	registryConfigStore, err := ProvideRegistryConfigStore(diStage1ConfigStore, config)
	if err != nil {
		return nil, err
	}
	registryDiFlags, err := registry.ProvideFlags(registryConfigStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	transportConfigStore, err := ProvideTransportConfigStore(diStage1ConfigStore, config)
	if err != nil {
		return nil, err
	}
	transportDiFlags, err := transport.ProvideFlags(transportConfigStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	cliDiFlags, err := cli.ProvideFlags(cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	storeConfigStore, err := ProvideStoreConfigStore(diStage1ConfigStore, config)
	if err != nil {
		return nil, err
	}
	storeDiFlags, err := store.ProvideFlags(storeConfigStore, cliConfigStore, cliCLI)
	if err != nil {
		return nil, err
	}
	cliArgs := ProvideCliArgs()
	parsedCli, err := ProvideInitializedCLI(diFlags, registryDiFlags, transportDiFlags, cliDiFlags, storeDiFlags, options, cliCLI, cliArgs)
	if err != nil {
		return nil, err
	}
	diConfig, err := cli.ProvideDiConfig(diStage1ConfigStore, parsedCli, cliDiFlags, cliConfigStore)
	if err != nil {
		return nil, err
	}
	microwireDiConfig, err := ProvideConfigFile(diStage1ConfigStore, parsedCli, diConfig, config)
	if err != nil {
		return nil, err
	}
	diStage2ConfigStore, err := ProvideStage2ConfigStore(microwireDiConfig, config)
	if err != nil {
		return nil, err
	}
	brokerDiConfig, err := broker.ProvideDiConfig(diStage2ConfigStore, diFlags, cliConfigStore, configStore)
	if err != nil {
		return nil, err
	}
	registryDiConfig, err := registry.ProvideDiConfig(diStage2ConfigStore, registryDiFlags, cliConfigStore, registryConfigStore)
	if err != nil {
		return nil, err
	}
	transportDiConfig, err := transport.ProvideDiConfig(diStage2ConfigStore, transportDiFlags, cliConfigStore, transportConfigStore)
	if err != nil {
		return nil, err
	}
	diStage3ConfigStore, err := ProvideStage3ConfigStore(diStage2ConfigStore, brokerDiConfig, registryDiConfig, transportDiConfig, config)
	if err != nil {
		return nil, err
	}
	brokerBroker, err := broker.Provide(diStage3ConfigStore, configStore, brokerDiConfig)
	if err != nil {
		return nil, err
	}
	registryRegistry, err := registry.Provide(diStage3ConfigStore, registryConfigStore, registryDiConfig)
	if err != nil {
		return nil, err
	}
	storeDiConfig, err := store.ProvideDiConfig(diStage2ConfigStore, storeDiFlags, cliConfigStore, storeConfigStore)
	if err != nil {
		return nil, err
	}
	storeStore, err := store.Provide(diStage3ConfigStore, storeConfigStore, storeDiConfig)
	if err != nil {
		return nil, err
	}
	transportTransport, err := transport.Provide(diStage3ConfigStore, transportConfigStore, transportDiConfig)
	if err != nil {
		return nil, err
	}
	v, err := ProvideMicroOpts(options, parsedCli, diStage3ConfigStore, brokerBroker, registryRegistry, storeStore, transportTransport)
	if err != nil {
		return nil, err
	}
	service, err := ProvideMicroService(config, options, v)
	if err != nil {
		return nil, err
	}
	return service, nil
}
