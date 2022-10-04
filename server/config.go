// Code generated with jinja2 templates. DO NOT EDIT.

package server

type Config struct {
	Enabled          bool                `json:"enabled" yaml:"Enabled"`
	Plugin           string              `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Address          string              `json:"address,omitempty" yaml:"Address,omitempty"`
	ID               string              `json:"id,omitempty" yaml:"ID,omitempty"`
	Name             string              `json:"name,omitempty" yaml:"Name,omitempty"`
	Version          string              `json:"version,omitempty" yaml:"Version,omitempty"`
	Metadata         map[string]string   `json:"metadata,omitempty" yaml:"Metadata,omitempty"`
	RegisterTTL      int                 `json:"register_ttl,omitempty" yaml:"RegisterTTL,omitempty"`
	RegisterInterval int                 `json:"register_interval,omitempty" yaml:"RegisterInterval,omitempty"`
	WrapSubscriber   []SubscriberWrapper `json:"-" yaml:"-"`
	WrapHandler      []HandlerWrapper    `json:"-" yaml:"-"`
}

type sourceConfig struct {
	Server Config `json:"broker" yaml:"Server"`
}

func NewConfig() *Config {
	return &Config{
		Enabled:          true,
		Plugin:           "rpc",
		Address:          "",
		ID:               "",
		Name:             "",
		Version:          "",
		Metadata:         make(map[string]string),
		RegisterTTL:      60,
		RegisterInterval: 30,
		WrapSubscriber:   []SubscriberWrapper{},
		WrapHandler:      []HandlerWrapper{},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Enabled != def.Enabled {
		d.Enabled = src.Enabled
	}

	if src.Plugin != def.Plugin {
		d.Plugin = src.Plugin
		d.Address = src.Address
		d.ID = src.ID
		d.Name = src.Name
		d.Version = src.Version
		d.Metadata = src.Metadata
	}

	return nil
}
