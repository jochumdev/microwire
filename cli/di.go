package cli

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-micro/microwire/di"
	"github.com/go-micro/plugins/v4/config/encoder/toml"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
	"github.com/google/wire"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	uJson "go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"
	uFile "go-micro.dev/v4/util/file"
)

type DiConfig string

type DiConfigor config.Config

type DiParsed struct {
	ConfigFile string
}

func ProvideCli(
	config *Config,
) (Cli, error) {
	c, err := Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown cli given: %v", err)
	}

	return c(), nil
}

func ProvideParsed(
	config *Config,
	c Cli,
) (*DiParsed, error) {
	result := &DiParsed{}

	// User flags
	for _, f := range config.Flags {
		if err := c.Add(f.AsOptions()...); err != nil {
			return result, err
		}
	}

	if config.NoFlags {
		if err := c.Add(
			Name(PrefixName(config.ArgPrefix, "config")),
			Usage("Config file"),
			Default(config.ConfigFile),
			EnvVars(PrefixEnv(config.ArgPrefix, "config")),
			Destination(&result.ConfigFile),
		); err != nil {
			return result, err
		}
	}

	// Initialize the CLI / parse flags
	if err := c.Parse(
		os.Args,
		CliName(config.Name),
		CliVersion(config.Version),
		CliDescription(config.Description),
		CliUsage(config.Usage),
	); err != nil {
		return result, err
	}

	return result, nil
}

func ProvideConfig(
	_ di.DiFlags,
	diFlags *DiParsed,
	cfg *Config,
) (DiConfig, error) {
	if cfg.NoFlags {
		// Defined silently ignore that
		defConfig := NewConfig()
		defConfig.ConfigFile = diFlags.ConfigFile
		if err := cfg.Merge(defConfig); err != nil {
			return "", err
		}
	}

	return DiConfig(cfg.ConfigFile), nil
}

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

var DiSet = wire.NewSet(ProvideParsed, ProvideConfig)
