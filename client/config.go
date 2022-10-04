// Code generated with jinja2 templates. DO NOT EDIT.

package client

type ClientConfig struct {
	Enabled            bool   `json:"enabled" yaml:"Enabled"`
	Plugin             string `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	PoolSize           int    `json:"pool_size,omitempty" yaml:"PoolSize,omitempty"`
	PoolTTL            string `json:"pool_ttl,omitempty" yaml:"PoolTTL,omitempty"`
	PoolRequestTimeout string `json:"pool_request_timeout,omitempty" yaml:"PoolRequestTimeout,omitempty"`
	PoolRetries        int    `json:"pool_retries,omitempty" yaml:"PoolRetries,omitempty"`
}

type Config struct {
	Client ClientConfig `json:"broker" yaml:"Client"`
}

func NewConfig() *Config {
	return &Config{
		Client: ClientConfig{
			Enabled:            true,
			Plugin:             "rpc",
			PoolSize:           1,
			PoolTTL:            "1m",
			PoolRequestTimeout: "5s",
			PoolRetries:        1,
		},
	}
}

func (d *Config) Merge(src *Config) error {
	def := NewConfig()

	if src.Client.Enabled != def.Client.Enabled {
		d.Client.Enabled = src.Client.Enabled
	}

	if src.Client.Plugin != def.Client.Plugin {
		d.Client.Plugin = src.Client.Plugin
		d.Client.PoolSize = src.Client.PoolSize
		d.Client.PoolTTL = src.Client.PoolTTL
		d.Client.PoolRequestTimeout = src.Client.PoolRequestTimeout
		d.Client.PoolRetries = src.Client.PoolRetries
	}

	return nil
}
