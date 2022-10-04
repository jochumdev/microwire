// Code generated with jinja2 templates. DO NOT EDIT.

package auth

import ()

type Config struct {
	Enabled    bool   `json:"enabled" yaml:"Enabled"`
	Plugin     string `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	ID         string `json:"client,omitempty" yaml:"ID,omitempty"`
	Secret     string `json:"secret,omitempty" yaml:"Secret,omitempty"`
	PublicKey  string `json:"public_key,omitempty" yaml:"PublicKey,omitempty"`
	PrivateKey string `json:"private_key,omitempty" yaml:"PrivateKey,omitempty"`
	Namespace  string `json:"namespace,omitempty" yaml:"Namespace,omitempty"`
}

type sourceConfig struct {
	Auth Config `json:"auth" yaml:"Auth"`
}

func NewConfig() *Config {
	config := &Config{
		Enabled:    false,
		Plugin:     "",
		ID:         "",
		Secret:     "",
		PublicKey:  "",
		PrivateKey: "",
		Namespace:  "",
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
		d.ID = src.ID
		d.Secret = src.Secret
		d.PublicKey = src.PublicKey
		d.PrivateKey = src.PrivateKey
		d.Namespace = src.Namespace
	}

	return nil
}
