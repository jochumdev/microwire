package microwire

import (
	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
)

type DiDefaultConfigStore ConfigStore

func ProvideConfigStore() (ConfigStore, error) {
	return NewConfigStore()
}

func ProvideDefaultConfigStore() (DiDefaultConfigStore, error) {
	return NewConfigStore()
}

// ProvideStage1ConfigStore loads the config from opts
func ProvideStage1ConfigStore(
	opts *Options,
	config ConfigStore,
	defConfig DiDefaultConfigStore,
) (mWire.DiStage1ConfigStore, error) {
	if len(opts.ArgPrefix) > 0 && config.GetCli().ArgPrefix == defConfig.GetCli().ArgPrefix {
		config.GetCli().ArgPrefix = opts.ArgPrefix
	}

	config.GetCli().NoFlags = opts.NoFlags

	for n, p := range opts.Components {
		switch n {
		case mCli.ComponentName:
			if config.GetCli().Plugin == defConfig.GetCli().Plugin {
				config.GetCli().Plugin = p
			}
		case mBroker.ComponentName:
			config.GetBroker().Enabled = true
			if config.GetBroker().Plugin == defConfig.GetBroker().Plugin {
				config.GetBroker().Plugin = p
			}
		case mRegistry.ComponentName:
			config.GetRegistry().Enabled = true
			if config.GetRegistry().Plugin == defConfig.GetRegistry().Plugin {
				config.GetRegistry().Plugin = p
			}
		case mTransport.ComponentName:
			config.GetTransport().Enabled = true
			if config.GetTransport().Plugin == defConfig.GetTransport().Plugin {
				config.GetTransport().Plugin = p
			}
		}
	}

	return mWire.DiStage1ConfigStore{}, nil
}

// ProvideStage2ConfigStore loads the config from config sources
func ProvideStage2ConfigStore(
	_ mWire.DiStage1ConfigStore,

	config ConfigStore,
) (mWire.DiStage2ConfigStore, error) {
	return mWire.DiStage2ConfigStore{}, nil
}

// ProvideStage3ConfigStore loads (env|flags) into the store
func ProvideStage3ConfigStore(
	_ mWire.DiStage2ConfigStore,
	_ mCli.ParsedCli,

	config ConfigStore,
) (mWire.DiStage3ConfigStore, error) {
	return mWire.DiStage3ConfigStore{}, nil
}

var DiConfigStagesSet = wire.NewSet(ProvideStage1ConfigStore, ProvideStage2ConfigStore, ProvideStage3ConfigStore)

func ProvideCliConfigStore(
	_ mWire.DiStage1ConfigStore,
	config ConfigStore,
) (*mCli.ConfigStore, error) {
	return config.GetCli(), nil
}

func ProvideBrokerConfigStore(
	_ mWire.DiStage1ConfigStore,
	config ConfigStore,
) (*mBroker.ConfigStore, error) {
	return config.GetBroker(), nil
}

func ProvideRegistryConfigStore(
	_ mWire.DiStage1ConfigStore,
	config ConfigStore,
) (*mRegistry.ConfigStore, error) {
	return config.GetRegistry(), nil
}

func ProvideTransportConfigStore(
	_ mWire.DiStage1ConfigStore,
	config ConfigStore,
) (*mTransport.ConfigStore, error) {
	return config.GetTransport(), nil
}
