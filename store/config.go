package store

type StoreConfig struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
	Database  string   `json:"database,omitempty" yaml:"Database,omitempty"`
	Table     string   `json:"table,omitempty" yaml:"Table,omitempty"`
}

type Config struct {
	Store StoreConfig `json:"broker" yaml:"Store"`
}

func NewConfig() *Config {
	return &Config{
		Store: StoreConfig{
			Enabled:   false,
			Plugin:    "",
			Addresses: []string{},
			Database:  "",
			Table:     "",
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	d.Store.Enabled = src.Store.Enabled

	if src.Store.Plugin != def.Store.Plugin {
		d.Store.Plugin = src.Store.Plugin
		d.Store.Addresses = src.Store.Addresses
		d.Store.Database = src.Store.Database
		d.Store.Table = src.Store.Table
	}

	return nil
}
