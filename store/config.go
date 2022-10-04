// Code generated with jinja2 templates. DO NOT EDIT.

package store

type Config struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
	Database  string   `json:"database,omitempty" yaml:"Database,omitempty"`
	Table     string   `json:"table,omitempty" yaml:"Table,omitempty"`
}

type sourceConfig struct {
	Store Config `json:"" yaml:"Store"`
}

func NewConfig() *Config {
	return &Config{
		Enabled:   false,
		Plugin:    "",
		Addresses: []string{},
		Database:  "",
		Table:     "",
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

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
