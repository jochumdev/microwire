// Code generated with jinja2 templates. DO NOT EDIT.

package cache

import ()

type Config struct {
	Enabled bool   `json:"enabled" yaml:"Enabled"`
	Plugin  string `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
}

type sourceConfig struct {
	Cache Config `json:"cache" yaml:"Cache"`
}

func NewConfig() *Config {
	config := &Config{
		Enabled: false,
		Plugin:  "",
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

	}

	return nil
}
