// Code generated with jinja2 templates. DO NOT EDIT.

package broker

type Config struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
}

type sourceConfig struct {
	Broker Config `json:"" yaml:"Broker"`
}

func NewConfig() *Config {
	return &Config{
		Enabled:   true,
		Plugin:    "http",
		Addresses: []string{},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Enabled != def.Enabled {
		d.Enabled = src.Enabled
	}

	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
		d.Addresses = src.Addresses
	}

	return nil
}
