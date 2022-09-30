package broker

type ConfigStore struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
}

func NewConfigStore() ConfigStore {
	return ConfigStore{
		Enabled:   true,
		Plugin:    "http",
		Addresses: []string{},
	}
}

func (d *ConfigStore) Merge(src *ConfigStore) error {
	def := NewConfigStore()

	if src.Enabled != def.Enabled {
		d.Enabled = src.Enabled
	}

	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
		d.Addresses = src.Addresses

	}

	return nil
}
