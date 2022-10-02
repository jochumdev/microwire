// Code generated with jinja2 templates. DO NOT EDIT.

package broker

type BrokerConfig struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
}

type Config struct {
	Broker BrokerConfig `json:"broker" yaml:"Broker"`
}

func NewConfig() *Config {
	return &Config{
		Broker: BrokerConfig{
			Enabled:   true,
			Plugin:    "http",
			Addresses: []string{},
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Broker.Enabled != def.Broker.Enabled {
		d.Broker.Enabled = src.Broker.Enabled
	}

	if src.Broker.Plugin != def.Broker.Plugin {
		d.Broker.Plugin = src.Broker.Plugin
		d.Broker.Addresses = src.Broker.Addresses
	}

	return nil
}
