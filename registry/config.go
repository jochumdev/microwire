package registry

type RegistryConfig struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
}

type Config struct {
	Registry RegistryConfig `json:"broker" yaml:"Registry"`
}

func NewConfig() *Config {
	return &Config{
		Registry: RegistryConfig{
			Enabled:   true,
			Plugin:    "mdns",
			Addresses: []string{},
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	d.Registry.Enabled = src.Registry.Enabled

	if src.Registry.Plugin != def.Registry.Plugin {
		d.Registry.Plugin = src.Registry.Plugin
		d.Registry.Addresses = src.Registry.Addresses
	}

	return nil
}
