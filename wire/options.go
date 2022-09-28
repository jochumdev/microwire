package wire

import (
	"github.com/go-micro/microwire/cli"
	"go-micro.dev/v4"
)

type HookFunc func() error
type ActionFunc func(cli.CLI, micro.Service) error

type Options struct {
	ArgPrefix   string
	Name        string
	Description string
	Version     string
	Usage       string
	Flags       []cli.Flag

	Components map[string]string

	// Livecycle
	Actions     []ActionFunc
	BeforeStart []HookFunc
	BeforeStop  []HookFunc
	AfterStart  []HookFunc
	AfterStop   []HookFunc
}

type Option func(*Options)

func ArgPrefix(n string) Option {
	return func(o *Options) {
		o.ArgPrefix = n
	}
}

func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

func Description(n string) Option {
	return func(o *Options) {
		o.Description = n
	}
}

func Version(n string) Option {
	return func(o *Options) {
		o.Version = n
	}
}

func Usage(n string) Option {
	return func(o *Options) {
		o.Usage = n
	}
}

func Flags(n []cli.Flag) Option {
	return func(o *Options) {
		o.Flags = n
	}
}

func Component(name string, plugin string) Option {
	return func(o *Options) {
		o.Components[name] = plugin
	}
}

func Action(fn ActionFunc) Option {
	return func(o *Options) {
		o.Actions = append(o.Actions, fn)
	}
}

// Before and Afters

// BeforeStart run funcs before service starts
func BeforeStart(fn HookFunc) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

// BeforeStop run funcs before service stops
func BeforeStop(fn HookFunc) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

// AfterStart run funcs after service starts
func AfterStart(fn HookFunc) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

// AfterStop run funcs after service stops
func AfterStop(fn HookFunc) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
