package micro

import (
	mCli "github.com/go-micro/microwire/v5/cli"
)

type HookFunc func() error
type ActionFunc func(Service) error

type MwOptions struct {
	ArgPrefix   string
	Name        string
	Description string
	Version     string
	Usage       string
	NoFlags     bool
	Config      string
	Flags       []mCli.Flag

	// Livecycle
	Actions     []ActionFunc
	BeforeStart []HookFunc
	BeforeStop  []HookFunc
	AfterStart  []HookFunc
	AfterStop   []HookFunc
}

func NewMwOptions(opts ...MwOption) *MwOptions {
	options := &MwOptions{
		ArgPrefix:   "",
		Name:        "",
		Description: "",
		Version:     "",
		Usage:       "",
		NoFlags:     false,
		Flags:       []mCli.Flag{},

		Actions:     []ActionFunc{},
		BeforeStart: []HookFunc{},
		BeforeStop:  []HookFunc{},
		AfterStart:  []HookFunc{},
		AfterStop:   []HookFunc{},
	}

	for _, o := range opts {
		o(options)
	}

	return options
}

type MwOption func(*MwOptions)

func MwArgPrefix(n string) MwOption {
	return func(o *MwOptions) {
		o.ArgPrefix = n
	}
}

func MwName(n string) MwOption {
	return func(o *MwOptions) {
		o.Name = n
	}
}

func MwDescription(n string) MwOption {
	return func(o *MwOptions) {
		o.Description = n
	}
}

func MwVersion(n string) MwOption {
	return func(o *MwOptions) {
		o.Version = n
	}
}

func MwUsage(n string) MwOption {
	return func(o *MwOptions) {
		o.Usage = n
	}
}

func MwNoFlags() MwOption {
	return func(o *MwOptions) {
		o.NoFlags = true
	}
}

func MwFlags(n []mCli.Flag) MwOption {
	return func(o *MwOptions) {
		o.Flags = n
	}
}

func MwConfig(n string) MwOption {
	return func(o *MwOptions) {
		o.Config = n
	}
}

func MwAction(fn ActionFunc) MwOption {
	return func(o *MwOptions) {
		o.Actions = append(o.Actions, fn)
	}
}

// Before and Afters

// BeforeStart run funcs before service starts
func MwBeforeStart(fn HookFunc) MwOption {
	return func(o *MwOptions) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

// BeforeStop run funcs before service stops
func MwBeforeStop(fn HookFunc) MwOption {
	return func(o *MwOptions) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

// AfterStart run funcs after service starts
func MwAfterStart(fn HookFunc) MwOption {
	return func(o *MwOptions) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

// AfterStop run funcs after service stops
func MwAfterStop(fn HookFunc) MwOption {
	return func(o *MwOptions) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}
