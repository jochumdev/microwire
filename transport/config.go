// Code generated with jinja2 templates. DO NOT EDIT.

package transport

import (
	"github.com/go-micro/microwire/v5/logger"
)

type Config struct {
	Enabled   bool           `json:"enabled" yaml:"Enabled"`
	Plugin    string         `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Logger    *logger.Config `json:"logger,omitempty" yaml:"Logger,omitempty"`
	Addresses []string       `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
}

type sourceConfig struct {
	Transport Config `json:"transport" yaml:"Transport"`
}

func NewConfig() *Config {
	config := &Config{
		Enabled:   true,
		Plugin:    "http",
		Addresses: []string{},
		Logger:    logger.NewConfig(),
	}

	return config
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if !d.Enabled && !def.Enabled && src.Enabled {
		d.Enabled = true
	}

	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
		d.Addresses = src.Addresses
	}

	d.Logger.Merge(src.Logger)

	return nil
}
