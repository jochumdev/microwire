// Code generated with jinja2 templates. DO NOT EDIT.

package client

import (
	"github.com/go-micro/microwire/v5/logger"
)

type Config struct {
	Enabled            bool           `json:"enabled" yaml:"Enabled"`
	Plugin             string         `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Logger             *logger.Config `json:"logger,omitempty" yaml:"Logger,omitempty"`
	PoolSize           int            `json:"pool_size,omitempty" yaml:"PoolSize,omitempty"`
	PoolTTL            string         `json:"pool_ttl,omitempty" yaml:"PoolTTL,omitempty"`
	PoolRequestTimeout string         `json:"pool_request_timeout,omitempty" yaml:"PoolRequestTimeout,omitempty"`
	PoolRetries        int            `json:"pool_retries,omitempty" yaml:"PoolRetries,omitempty"`
	WrapCall           []CallWrapper  `json:"-" yaml:"-"`
}

type sourceConfig struct {
	Client Config `json:"client" yaml:"Client"`
}

func NewConfig() *Config {
	config := &Config{
		Enabled:            true,
		Plugin:             "rpc",
		PoolSize:           1,
		PoolTTL:            "1m",
		PoolRequestTimeout: "5s",
		PoolRetries:        1,
		WrapCall:           []CallWrapper{},
		Logger:             logger.NewConfig(),
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
		d.PoolSize = src.PoolSize
		d.PoolTTL = src.PoolTTL
		d.PoolRequestTimeout = src.PoolRequestTimeout
		d.PoolRetries = src.PoolRetries
	}

	d.Logger.Merge(src.Logger)

	return nil
}
