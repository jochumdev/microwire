// Code generated with jinja2 templates. DO NOT EDIT.

package auth

type AuthConfig struct {
	Enabled    bool   `json:"enabled" yaml:"Enabled"`
	Plugin     string `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	ID         string `json:"client,omitempty" yaml:"ID,omitempty"`
	Secret     string `json:"secret,omitempty" yaml:"Secret,omitempty"`
	PublicKey  string `json:"public_key,omitempty" yaml:"PublicKey,omitempty"`
	PrivateKey string `json:"private_key,omitempty" yaml:"PrivateKey,omitempty"`
	Namespace  string `json:"namespace,omitempty" yaml:"Namespace,omitempty"`
}

type Config struct {
	Auth AuthConfig `json:"broker" yaml:"Auth"`
}

func NewConfig() *Config {
	return &Config{
		Auth: AuthConfig{
			Enabled:    false,
			Plugin:     "",
			ID:         "",
			Secret:     "",
			PublicKey:  "",
			PrivateKey: "",
			Namespace:  "",
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Auth.Enabled != def.Auth.Enabled {
		d.Auth.Enabled = src.Auth.Enabled
	}

	if src.Auth.Plugin != def.Auth.Plugin {
		d.Auth.Plugin = src.Auth.Plugin
		d.Auth.ID = src.Auth.ID
		d.Auth.Secret = src.Auth.Secret
		d.Auth.PublicKey = src.Auth.PublicKey
		d.Auth.PrivateKey = src.Auth.PrivateKey
		d.Auth.Namespace = src.Auth.Namespace
	}

	return nil
}
