// Code generated with jinja2 templates. DO NOT EDIT.

package server

type ServerConfig struct {
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

type Config struct {
	Server ServerConfig `json:"broker" yaml:"Server"`
}

func NewConfig() *Config {
	return &Config{
		Server: ServerConfig{
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
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Server.Enabled != def.Server.Enabled {
		d.Server.Enabled = src.Server.Enabled
	}

	if src.Server.Plugin != def.Server.Plugin {
		d.Server.Plugin = src.Server.Plugin
		d.Server.Address = src.Server.Address
		d.Server.ID = src.Server.ID
		d.Server.Name = src.Server.Name
		d.Server.Version = src.Server.Version
		d.Server.Metadata = src.Server.Metadata
	}

	return nil
}
