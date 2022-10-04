// Code generated with jinja2 templates. DO NOT EDIT.

package logger

import ()

type Config struct {
	Enabled         bool                   `json:"enabled" yaml:"Enabled"`
	Plugin          string                 `json:"plugin,omitempty" yaml:"Plugin,omitempty"`
	Fields          map[string]interface{} `json:"fields" yaml:"Fields"`
	Level           string                 `json:"level" yaml:"Level"`
	CallerSkipCount int                    `json:"caller_skip_count" yaml:"CallerSkipCount"`
}

type sourceConfig struct {
	Logger Config `json:"logger" yaml:"Logger"`
}

func NewConfig() *Config {
	config := &Config{
		Enabled:         false,
		Plugin:          "default",
		Fields:          make(map[string]interface{}),
		Level:           InfoLevel.String(),
		CallerSkipCount: 2,
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
		d.Fields = src.Fields
		d.Level = src.Level
		d.CallerSkipCount = src.CallerSkipCount
	}

	return nil
}
