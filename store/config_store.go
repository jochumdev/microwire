package store

type ConfigStore struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
	Database  string   `json:"database,omitempty" yaml:"Database,omitempty"`
	Table     string   `json:"table,omitempty" yaml:"Table,omitempty"`
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

	d.Enabled = src.Enabled

	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
		d.Addresses = src.Addresses
		d.Database = src.Database
		d.Table = src.Table
	}

	return nil
}
