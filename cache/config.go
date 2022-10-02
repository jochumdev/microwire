// Code generated with jinja2 templates. DO NOT EDIT.

package cache

type CacheConfig struct {
	Enabled bool   `json:"enabled" yaml:"Enabled"`
	Plugin  string `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
}

type Config struct {
	Cache CacheConfig `json:"broker" yaml:"Cache"`
}

func NewConfig() *Config {
	return &Config{
		Cache: CacheConfig{
			Enabled: false,
			Plugin:  "",
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Cache.Enabled != def.Cache.Enabled {
		d.Cache.Enabled = src.Cache.Enabled
	}

	if src.Cache.Plugin != def.Cache.Plugin {
		d.Cache.Plugin = src.Cache.Plugin

	}

	return nil
}
