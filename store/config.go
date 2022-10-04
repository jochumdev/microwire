// Code generated with jinja2 templates. DO NOT EDIT.

package store

import ()

type Config struct {
	Enabled   bool     `json:"enabled" yaml:"Enabled"`
	Plugin    string   `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Addresses []string `json:"addresses,omitempty" yaml:"Addresses,omitempty"`
	Database  string   `json:"database,omitempty" yaml:"Database,omitempty"`
	Table     string   `json:"table,omitempty" yaml:"Table,omitempty"`
}

type sourceConfig struct {
	Store Config `json:"store" yaml:"Store"`
}

func NewConfig() *Config {
	config := &Config{
		Enabled:   false,
		Plugin:    "",
		Addresses: []string{},
		Database:  "",
		Table:     "",
	}

	return config
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if !d.Enabled && !def.Enabled && src.Enabled {
		d.Enabled = true
	}

	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
		d.Addresses = src.Addresses
		d.Database = src.Database
		d.Table = src.Table
	}

	return nil
}
