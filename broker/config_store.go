package broker

import (
	"golang.org/x/exp/slices"
)

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

	if src.Enabled != def.Enabled {
		d.Enabled = src.Enabled
	}
	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
	}
	if slices.Compare(src.Addresses, def.Addresses) != 0 {
		d.Addresses = src.Addresses
	}

	return nil
}
