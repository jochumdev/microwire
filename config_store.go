package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
)

type ConfigStore interface {
	GetBroker() *mBroker.ConfigStore
	GetCli() *mCli.ConfigStore
	GetRegistry() *mRegistry.ConfigStore
	GetTransport() *mTransport.ConfigStore
}

type ConfigStoreImpl struct {
	Broker    mBroker.ConfigStore    `json:"broker" yaml:"Broker"`
	Global    mCli.ConfigStore       `json:"global" yaml:"Global"`
	Registry  mRegistry.ConfigStore  `json:"registry" yaml:"Registry"`
	Transport mTransport.ConfigStore `json:"transport" yaml:"Transport"`
}

func (s *ConfigStoreImpl) GetBroker() *mBroker.ConfigStore {
	return &s.Broker
}

func (s *ConfigStoreImpl) GetCli() *mCli.ConfigStore {
	return &s.Global
}

func (s *ConfigStoreImpl) GetRegistry() *mRegistry.ConfigStore {
	return &s.Registry
}

func (s *ConfigStoreImpl) GetTransport() *mTransport.ConfigStore {
	return &s.Transport
}

func NewConfigStore() (ConfigStore, error) {
	return &ConfigStoreImpl{
		Broker:    mBroker.DefaultConfigStore(),
		Global:    mCli.DefaultConfigStore(),
		Registry:  mRegistry.DefaultConfigStore(),
		Transport: mTransport.DefaultConfigStore(),
	}, nil
}
