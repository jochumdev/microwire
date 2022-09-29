package microwire

import (
	"fmt"
	"path/filepath"
	"strings"

	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/go-micro/plugins/v4/config/encoder/toml"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/google/wire"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	uJson "go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
	uFile "go-micro.dev/v4/util/file"
)

func ProvideConfigStore() (ConfigStore, error) {
	return NewConfigStore()
}

// ProvideStage1ConfigStore loads the config from opts
func ProvideStage1ConfigStore(
	opts *Options,
	config ConfigStore,
) (mWire.DiStage1ConfigStore, error) {
	defConfig, err := NewConfigStore()
	if err != nil {
		return mWire.DiStage1ConfigStore{}, err
	}

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
	_ mCli.ParsedCli,
	_ mCli.DiConfig,

	cfg ConfigStore,
) (mWire.DiStage2ConfigStore, error) {
	if len(cfg.GetCli().ConfigFile) == 0 {
		return mWire.DiStage2ConfigStore{}, nil
	}

	var configor config.Config
	var err error
	switch strings.ToLower(filepath.Ext(cfg.GetCli().ConfigFile)) {
	case ".toml":
		configor, err = config.NewConfig(
			config.WithSource(file.NewSource(file.WithPath(cfg.GetCli().ConfigFile))),
			config.WithReader(uJson.NewReader(reader.WithEncoder(toml.NewEncoder()))),
		)
	case ".yaml":
		configor, err = config.NewConfig(
			config.WithSource(file.NewSource(file.WithPath(cfg.GetCli().ConfigFile))),
			config.WithReader(uJson.NewReader(reader.WithEncoder(yaml.NewEncoder()))),
		)
	default:
		if ok, err2 := uFile.Exists(fmt.Sprintf("%s.toml", cfg.GetCli().ConfigFile)); ok && err2 == nil {
			configor, err = config.NewConfig(
				config.WithSource(file.NewSource(file.WithPath(fmt.Sprintf("%s.toml", cfg.GetCli().ConfigFile)))),
				config.WithReader(uJson.NewReader(reader.WithEncoder(toml.NewEncoder()))),
			)
		} else if ok, err2 := uFile.Exists(fmt.Sprintf("%s.yaml", cfg.GetCli().ConfigFile)); ok && err2 == nil {
			configor, err = config.NewConfig(
				config.WithSource(file.NewSource(file.WithPath(fmt.Sprintf("%s.yaml", cfg.GetCli().ConfigFile)))),
				config.WithReader(uJson.NewReader(reader.WithEncoder(yaml.NewEncoder()))),
			)
		} else {
			return mWire.DiStage2ConfigStore{}, fmt.Errorf("unknown config file '%s' with extension '%s' given", cfg.GetCli().ConfigFile, filepath.Ext(cfg.GetCli().ConfigFile))
		}
	}
	if err != nil {
		return mWire.DiStage2ConfigStore{}, err
	}
	if err := configor.Load(); err != nil {
		return mWire.DiStage2ConfigStore{}, err
	}

	defConfig, err := NewConfigStore()
	if err != nil {
		return mWire.DiStage2ConfigStore{}, err
	}
	if err := configor.Scan(defConfig); err != nil {
		return mWire.DiStage2ConfigStore{}, err
	}

	if err = cfg.GetBroker().Merge(defConfig.GetBroker()); err != nil {
		return mWire.DiStage2ConfigStore{}, err
	}

	if err = cfg.GetRegistry().Merge(defConfig.GetRegistry()); err != nil {
		return mWire.DiStage2ConfigStore{}, err
	}

	if err = cfg.GetTransport().Merge(defConfig.GetTransport()); err != nil {
		return mWire.DiStage2ConfigStore{}, err
	}

	return mWire.DiStage2ConfigStore{}, nil
}

// ProvideStage3ConfigStore marks that we loaded (env|flags) into the store
func ProvideStage3ConfigStore(
	_ mWire.DiStage2ConfigStore,

	_ mBroker.DiConfig,
	_ mRegistry.DiConfig,
	_ mTransport.DiConfig,

	cfg ConfigStore,
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
