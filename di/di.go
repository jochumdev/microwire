package di

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-micro/plugins/v4/config/encoder/toml"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	uJson "go-micro.dev/v4/config/reader/json"
	uFile "go-micro.dev/v4/util/file"

	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/source/file"
)

// DiFlags is a marker that the config has been loaded from compiled in opts
type DiFlags struct{}

// DiConfigData is a marker that the config has been loaded from different sources (yaml,json,toml,name it here) on top of DiStage1Config
type DiConfigData struct{}

type DiConfig string

type DiConfigor config.Config

func ProvideConfigor(
	configFile DiConfig,
) (DiConfigor, error) {
	if configFile == "" {
		// Ignore no configFile
		return nil, nil
	}

	// Guess the file extension
	strFp := strings.ToLower(string(configFile))
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
		if err := configor.Load(); err != nil {
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
		if err := configor.Load(); err != nil {
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
		if err := configor.Load(); err != nil {
			return nil, err
		}

		return configor, nil
	default:
		return nil, fmt.Errorf("unknown file extension '%s'", filepath.Ext(strFp))
	}
}
