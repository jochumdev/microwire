package store

type ConfigStore struct {
	Enabled   bool     `json:"enabled" yaml:"enabled"`
	Plugin    string   `json:"plugin" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses" yaml:"Addresses,omitempty"`
	Database  string   `json:"database" yaml:"Database,omitempty"`
	Table     string   `json:"table" yaml:"Table,omitempty"`
}

func NewConfigStore() ConfigStore {
	return ConfigStore{
		Enabled:   false,
		Plugin:    "",
		Addresses: []string{},
		Database:  "",
		Table:     "",
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
		d.Database = src.Database
		d.Table = src.Table
	}

	return nil
}
