package wire

import (
	"github.com/go-micro/microwire/cli"
)

type MyFlags bool
type InitializedCli cli.CLI
type CliArgs []string

func ProvideOptions(opts []Option) *Options {
	options := &Options{
		ArgPrefix:   "",
		Name:        "",
		Description: "",
		Version:     "",
		Usage:       "",
		Flags:       []cli.Flag{},

		Components: make(map[string]string),

		Actions:     []ActionFunc{},
		BeforeStart: []HookFunc{},
		BeforeStop:  []HookFunc{},
		AfterStart:  []HookFunc{},
		AfterStop:   []HookFunc{},
	}

	for _, o := range opts {
		o(options)
	}

	// Set default components
	defaultComponents := map[string]string{
		ComponentBroker:    "http",
		ComponentCli:       "urfave",
		ComponentRegistry:  "mdns",
		ComponentTransport: "http",
	}
	for n, v := range defaultComponents {
		if _, ok := options.Components[n]; !ok {
			options.Components[n] = v
		}
	}

	return options
}
