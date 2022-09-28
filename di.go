package microwire

import (
	"fmt"
	"os"

	mBroker "github.com/go-micro/microwire/broker"
	mCli "github.com/go-micro/microwire/cli"
	mRegistry "github.com/go-micro/microwire/registry"
	mTransport "github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"github.com/google/wire"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/errors"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/transport"
)

type CliArgs []string

func ProvideOptions(opts []mWire.Option) *mWire.Options {
	options := &mWire.Options{
		ArgPrefix:   "",
		Name:        "",
		Description: "",
		Version:     "",
		Usage:       "",
		Flags:       []mCli.Flag{},

		Components: make(map[string]string),

		Actions:     []mWire.ActionFunc{},
		BeforeStart: []mWire.HookFunc{},
		BeforeStop:  []mWire.HookFunc{},
		AfterStart:  []mWire.HookFunc{},
		AfterStop:   []mWire.HookFunc{},
	}

	for _, o := range opts {
		o(options)
	}

	// Set default components
	defaultComponents := map[string]string{
		mBroker.ComponentName:    "http",
		mCli.ComponentName:       "urfave",
		mRegistry.ComponentName:  "mdns",
		mTransport.ComponentName: "http",
	}
	for n, v := range defaultComponents {
		if _, ok := options.Components[n]; !ok {
			options.Components[n] = v
		}
	}

	return options
}

func ProvideCLI(opts *mWire.Options) (mCli.CLI, error) {
	c, err := mCli.Plugins.Get(opts.Components[mCli.ComponentName])
	if err != nil {
		return nil, fmt.Errorf("unknown cli given: %v", err)
	}

	return c(), nil
}

func ProvideCliArgs() CliArgs {
	return os.Args
}

func ProvideInitializedCLI(
	opts *mWire.Options,
	c mCli.CLI,
	args CliArgs,
	_ *mBroker.DiFlags,
	_ *mRegistry.DiFlags,
	_ *mTransport.DiFlags,
) (mWire.InitializedCli, error) {
	// User flags
	for _, f := range opts.Flags {
		switch f.FlagType {
		case mCli.FlagTypeString:
			c.AddString(f.AsOptions()...)
		case mCli.FlagTypeInt:
			c.AddInt(f.AsOptions()...)
		default:
			return nil, errors.InternalServerError("USER_FLAG_WITHOUT_A_DEFAULTOPTION", "found a flag without a default option")
		}
	}

	// Initialize the CLI / parse flags
	if err := c.Init(
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
	opts *mWire.Options,
	c mWire.InitializedCli,
	broker broker.Broker,
	registry registry.Registry,
	transport transport.Transport,
) ([]micro.Option, error) {
	result := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
	}

	for n := range opts.Components {
		switch n {
		case mCli.ComponentName:
			continue
		case mBroker.ComponentName:
			result = append(result, micro.Broker(broker))
		case mRegistry.ComponentName:
			result = append(result, micro.Registry(registry))
		case mTransport.ComponentName:
			result = append(result, micro.Transport(transport))
		}
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

func ProvideMicroService(opts *mWire.Options, c mCli.CLI, mOpts []micro.Option) (micro.Service, error) {
	service := micro.NewService(
		mOpts...,
	)

	for _, fn := range opts.Actions {
		if err := fn(c, service); err != nil {
			return nil, err
		}
	}

	return service, nil
}

var AllComponentsSet = wire.NewSet(
	mBroker.Provide,
	mRegistry.Provide,
	mTransport.Provide,
)
