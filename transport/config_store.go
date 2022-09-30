package transport

type ConfigStore struct {
	Enabled   bool     `json:"enabled" yaml:"enabled"`
	Plugin    string   `json:"plugin" yaml:"Plugin"`
	Addresses []string `json:"addresses" yaml:"Addresses"`
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
