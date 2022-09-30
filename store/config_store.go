package store

type ConfigStore struct {
	Enabled   bool     `json:"enabled" yaml:"enabled"`
	Plugin    string   `json:"plugin" yaml:"Plugin"`
	Addresses []string `json:"addresses" yaml:"Addresses"`
	Database  string   `json:"database" yaml:"Database"`
	Table     string   `json:"table" yaml:"Table"`
}

func DefaultConfigStore() ConfigStore {
	return ConfigStore{
		Enabled:   false,
		Plugin:    "",
		Addresses: []string{},
		Database:  "",
		Table:     "",
	}
}

func (d *ConfigStore) Merge(src *ConfigStore) error {
	def := DefaultConfigStore()

	if src.Plugin != def.Plugin {
		d.Enabled = src.Enabled
		d.Plugin = src.Plugin
		if src.Addresses != nil {
			d.Addresses = src.Addresses
		}
		d.Database = src.Database
		d.Table = src.Table
	}

	return nil
}
