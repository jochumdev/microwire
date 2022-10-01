package transport

type TransportConfig struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
}

type Config struct {
	Transport TransportConfig `json:"broker" yaml:"Transport"`
}

func NewConfig() *Config {
	return &Config{
		Transport: TransportConfig{
			Enabled:   true,
			Plugin:    "http",
			Addresses: []string{},
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	d.Transport.Enabled = src.Transport.Enabled

	if src.Transport.Plugin != def.Transport.Plugin {
		d.Transport.Plugin = src.Transport.Plugin
		d.Transport.Addresses = src.Transport.Addresses
	}

	return nil
}
