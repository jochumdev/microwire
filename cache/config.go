// Code generated with jinja2 templates. DO NOT EDIT.

package cache

type Config struct {
	Enabled bool   `json:"enabled" yaml:"Enabled"`
	Plugin  string `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
}

type sourceConfig struct {
	Cache Config `json:"" yaml:"Cache"`
}

func NewConfig() *Config {
	return &Config{
		Enabled: false,
		Plugin:  "",
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Enabled != def.Enabled {
		d.Enabled = src.Enabled
	}

	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin

	}

	return nil
}
