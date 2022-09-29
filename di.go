package microwire

import (
	"fmt"
	"os"

	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"

	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/transport"
)

type CliArgs []string

func ProvideOptions(opts []Option) *Options {
	options := &Options{
		ArgPrefix:   "",
		Name:        "",
		Description: "",
		Version:     "",
		Usage:       "",
		NoFlags:     false,
		Flags:       []mCli.Flag{},

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

	return options
}

func ProvideCLI(
	_ mWire.DiStage1ConfigStore,
	config *mCli.ConfigStore,
) (mCli.CLI, error) {
	c, err := mCli.Plugins.Get(config.Plugin)
	if err != nil {
		return nil, fmt.Errorf("unknown cli given: %v", err)
	}

	return c(), nil
}

func ProvideCliArgs() CliArgs {
	return os.Args
}

func ProvideInitializedCLI(
	// These are here because they do something with cli.CLI
	_ *mBroker.DiFlags,
	_ *mRegistry.DiFlags,
	_ *mTransport.DiFlags,

	opts *Options,
	c mCli.CLI,
	args CliArgs,
) (mCli.ParsedCli, error) {
	// User flags
	for _, f := range opts.Flags {
		if err := c.Add(f.AsOptions()...); err != nil {
			return nil, err
		}
	}

	// Initialize the CLI / parse flags
	if err := c.Parse(
		args,
		mCli.CliName(opts.Name),
		mCli.CliVersion(opts.Version),
		mCli.CliDescription(opts.Description),
		mCli.CliUsage(opts.Usage),
	); err != nil {
		return nil, err
	}

	return c, nil
}

func ProvideMicroOpts(
	opts *Options,
	c mCli.ParsedCli,

	// Just to be save and clear, all components require stage3
	_ mWire.DiStage3ConfigStore,

	broker broker.Broker,
	registry registry.Registry,
	transport transport.Transport,
) ([]micro.Option, error) {
	result := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
	}

	if broker != nil {
		result = append(result, micro.Broker(broker))
	}
	if registry != nil {
		result = append(result, micro.Registry(registry))
	}
	if transport != nil {
		result = append(result, micro.Transport(transport))
	}

	for _, fn := range opts.BeforeStart {
		result = append(result, micro.BeforeStart(fn))
	}
	for _, fn := range opts.BeforeStop {
		result = append(result, micro.BeforeStop(fn))
	}
	for _, fn := range opts.AfterStart {
		result = append(result, micro.AfterStart(fn))
	}
	for _, fn := range opts.AfterStop {
		result = append(result, micro.AfterStop(fn))
	}

	return result, nil
}

func ProvideMicroService(config ConfigStore, opts *Options, mOpts []micro.Option) (micro.Service, error) {
	service := micro.NewService(
		mOpts...,
	)

	for _, fn := range opts.Actions {
		if err := fn(config, service); err != nil {
			return nil, err
		}
	}

	return service, nil
}
