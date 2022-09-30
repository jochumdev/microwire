package microwire

import (
	"fmt"
	"path/filepath"
	"strings"

	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mStore "github.com/go-micro/microwire/store"
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

type DiConfig config.Config

func ProvideConfigStore() (ConfigStore, error) {
	return NewConfigStore()
}

// ProvideStage1ConfigStore loads the config from opts
func ProvideStage1ConfigStore(
	opts *Options,
	config ConfigStore,
) (mWire.DiStage1ConfigStore, error) {
	defConfig := mCli.NewConfigStore()

	defConfig.ArgPrefix = opts.ArgPrefix
	defConfig.NoFlags = opts.NoFlags
	if err := config.GetCli().Merge(&defConfig); err != nil {
		return mWire.DiStage1ConfigStore{}, err
	}

	return mWire.DiStage1ConfigStore{}, nil
}

func ProvideConfigFile(
	_ mWire.DiStage1ConfigStore,
	_ mCli.ParsedCli,
	_ mCli.DiConfig,

	cfg ConfigStore,
) (DiConfig, error) {
	// Guess the file extension
	strFp := strings.ToLower(cfg.GetCli().ConfigFile)
	if ok, err2 := uFile.Exists(fmt.Sprintf("%s.toml", strFp)); ok && err2 == nil {
		strFp = fmt.Sprintf("%s.toml", strFp)
	} else if ok, err2 := uFile.Exists(fmt.Sprintf("%s.yaml", strFp)); ok && err2 == nil {
		strFp = fmt.Sprintf("%s.yaml", strFp)
	} else if ok, err2 := uFile.Exists(fmt.Sprintf("%s.yml", strFp)); ok && err2 == nil {
		strFp = fmt.Sprintf("%s.yml", strFp)
	} else if ok, err2 := uFile.Exists(fmt.Sprintf("%s.yml", strFp)); !ok || err2 != nil {
		return nil, fmt.Errorf("unknown config file '%s' with extension '%s' given", strFp, filepath.Ext(strFp))
	}

	// Provide config.Config based on the file extension
	switch filepath.Ext(strFp) {
	case ".toml":
		configor, err := config.NewConfig(
			config.WithSource(file.NewSource(file.WithPath(strFp))),
			config.WithReader(uJson.NewReader(reader.WithEncoder(toml.NewEncoder()))),
		)
		if err != nil {
			return nil, err
		}
		return configor, nil
	case ".yaml":
		configor, err := config.NewConfig(
			config.WithSource(file.NewSource(file.WithPath(strFp))),
			config.WithReader(uJson.NewReader(reader.WithEncoder(yaml.NewEncoder()))),
		)
		if err != nil {
			return nil, err
		}
		return configor, nil
	case ".yml":
		configor, err := config.NewConfig(
			config.WithSource(file.NewSource(file.WithPath(strFp))),
			config.WithReader(uJson.NewReader(reader.WithEncoder(yaml.NewEncoder()))),
		)
		if err != nil {
			return nil, err
		}
		return configor, nil
	default:
		return nil, fmt.Errorf("unknown file extension '%s'", filepath.Ext(strFp))
	}
}

// ProvideStage2ConfigStore loads the config from config sources
func ProvideStage2ConfigStore(
	configor DiConfig,
	cfg ConfigStore,
) (mWire.DiStage2ConfigStore, error) {
	if len(cfg.GetCli().ConfigFile) == 0 {
		return mWire.DiStage2ConfigStore{}, nil
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
	if err = cfg.GetStore().Merge(defConfig.GetStore()); err != nil {
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

var DiConfigStagesSet = wire.NewSet(ProvideStage1ConfigStore, ProvideConfigFile, ProvideStage2ConfigStore, ProvideStage3ConfigStore)

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

func ProvideStoreConfigStore(
	_ mWire.DiStage1ConfigStore,
	config ConfigStore,
) (*mStore.ConfigStore, error) {
	return config.GetStore(), nil
}

func ProvideTransportConfigStore(
	_ mWire.DiStage1ConfigStore,
	config ConfigStore,
) (*mTransport.ConfigStore, error) {
	return config.GetTransport(), nil
}
