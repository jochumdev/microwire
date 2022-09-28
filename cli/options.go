//go:build go1.18
// +build go1.18

package cli

import (
	"os"
	"path/filepath"
)

type CLIOptions struct {
	// Name is the name of the app
	Name string

	// Description is the description of the app
	Description string

	// Version is the Version of the app
	Version string

	// Usage is the apps usage string
	Usage string
}

type CLIOption func(*CLIOptions)

func CliName(n string) CLIOption {
	return func(o *CLIOptions) {
		o.Name = n
	}
}

func CliDescription(n string) CLIOption {
	return func(o *CLIOptions) {
		o.Description = n
	}
}

func CliVersion(n string) CLIOption {
	return func(o *CLIOptions) {
		o.Version = n
	}
}

func CliUsage(n string) CLIOption {
	return func(o *CLIOptions) {
		o.Usage = n
	}
}

func NewCLIOptions(opts ...CLIOption) *CLIOptions {
	options := &CLIOptions{
		Name:        filepath.Base(os.Args[0]),
		Description: "",
	}

	for _, o := range opts {
		o(options)
	}

	return options
}

type Options struct {
	Name    string
	EnvVars []string
	Usage   string

	DefaultString string
	DefaultInt    int
}

type Option func(*Options)

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func EnvVars(n ...string) Option {
	return func(o *Options) {
		o.EnvVars = n
	}
}

func Usage(n string) Option {
	return func(o *Options) {
		o.Usage = n
	}
}

func DefaultValue[T any](n T) Option {
	return func(o *Options) {
		switch any(n).(type) {
		case string:
			o.DefaultString = any(n).(string)
		case int:
			o.DefaultInt = any(n).(int)
		}
	}
}

func NewOptions(opts ...Option) *Options {
	options := &Options{
		Name:          "",
		EnvVars:       []string{},
		Usage:         "",
		DefaultString: "",
		DefaultInt:    0,
	}

	for _, o := range opts {
		o(options)
	}

	return options
}
