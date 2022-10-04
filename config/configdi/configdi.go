package configdi

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/config/reader"
	uJson "github.com/go-micro/microwire/v5/config/reader/json"
	"github.com/go-micro/microwire/v5/config/source/file"
	"github.com/go-micro/microwire/v5/di"
	uFile "github.com/go-micro/microwire/v5/util/file"
	"github.com/go-micro/plugins/v4/config/encoder/toml"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
)

func ProvideConfigor(
	configFile di.DiConfig,
) (config.Config, error) {
	if configFile == "" {
		// Ignore no configFile
		return nil, nil
	}

	// Guess the file extension
	strFp := strings.ToLower(string(configFile))
	if ok, err2 := uFile.Exists(strFp); !ok || err2 != nil {
		found := false
		for _, ext := range []string{".toml", ".tml", ".yaml", ".yml"} {

			if ok, err2 := uFile.Exists(fmt.Sprintf("%s%s", strFp, ext)); ok && err2 == nil {
				strFp = fmt.Sprintf("%s%s", strFp, ext)
				found = true
				break
			}
		}

		if !found {
			return nil, fmt.Errorf("unknown config file '%s' with extension '%s' given", strFp, filepath.Ext(strFp))
		}
	}

	// Provide config.Config based on the file extension
	switch filepath.Ext(strFp) {
	case ".toml", ".tml":
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
	case ".yaml", ".yml":
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
