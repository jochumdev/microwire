package broker

type ConfigStore struct {
	Enabled   bool     `json:"enabled" yaml:"enabled"`
	Plugin    string   `json:"plugin" yaml:"Plugin"`
	Addresses []string `json:"addresses" yaml:"Addresses"`
}

func DefaultConfigStore() ConfigStore {
	return ConfigStore{
		Enabled:   true,
		Plugin:    "http",
		Addresses: []string{},
	}
}

func (d *ConfigStore) Merge(src *ConfigStore) error {
	def := DefaultConfigStore()

	if src.Plugin != def.Plugin {
		d.Enabled = src.Enabled
		d.Plugin = src.Plugin
		d.Addresses = src.Addresses
	}

	return nil
}
