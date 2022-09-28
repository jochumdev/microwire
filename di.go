package microwire

import (
	"fmt"
	"os"

	"github.com/go-micro/microwire/broker"
	"github.com/go-micro/microwire/cli"
	"github.com/go-micro/microwire/registry"
	"github.com/go-micro/microwire/transport"
	mWire "github.com/go-micro/microwire/wire"
	"go-micro.dev/v4"
	"go-micro.dev/v4/errors"
)

func ProvideCLI(opts *mWire.Options) (cli.CLI, error) {
	c, err := cli.Container.Get(opts.Components[mWire.ComponentCli])
	if err != nil {
		return nil, fmt.Errorf("unknown cli given: %v", err)
	}

	return c(), nil
}

func ProvideCliArgs() mWire.CliArgs {
	return os.Args
}

func ProvideInitializedCLI(opts *mWire.Options, c cli.CLI, args mWire.CliArgs) (mWire.InitializedCli, error) {

	for n, _ := range opts.Components {
		switch n {
		case mWire.ComponentCli:
			continue
		case mWire.ComponentBroker:
			if err := broker.InjectFlags(opts, c); err != nil {
				return nil, err
			}
		case mWire.ComponentRegistry:
			if err := registry.InjectFlags(opts, c); err != nil {
				return nil, err
			}
		case mWire.ComponentTransport:
			if err := transport.InjectFlags(opts, c); err != nil {
				return nil, err
			}
		}
	}

	// User flags
	for _, f := range opts.Flags {
		switch f.FlagType {
		case cli.FlagTypeString:
			c.AddString(f.AsOptions()...)
		case cli.FlagTypeInt:
			c.AddInt(f.AsOptions()...)
		default:
			return nil, errors.InternalServerError("USER_FLAG_WITHOUT_A_DEFAULTOPTION", "found a flag without a default option")
		}
	}

	// Initialize the CLI / parse flags
	if err := c.Init(
		args,
		cli.CliName(opts.Name),
		cli.CliVersion(opts.Version),
		cli.CliDescription(opts.Description),
		cli.CliUsage(opts.Usage),
	); err != nil {
		return nil, err
	}

	return c, nil
}

func ProvideMicroOpts(opts *mWire.Options, c mWire.InitializedCli) ([]micro.Option, error) {
	result := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
	}

	for n, _ := range opts.Components {
		switch n {
		case mWire.ComponentCli:
			continue
		case mWire.ComponentBroker:
			b, err := broker.Inject(opts, c)
			if err != nil {
				return result, fmt.Errorf("unknown broker: %v", err)
			}
			result = append(result, micro.Broker(b))
		case mWire.ComponentRegistry:
			r, err := registry.Inject(opts, c)
			if err != nil {
				return result, fmt.Errorf("unknown registry: %v", err)
			}
			result = append(result, micro.Registry(r))
		case mWire.ComponentTransport:
			t, err := transport.Inject(opts, c)
			if err != nil {
				return result, fmt.Errorf("unknown transport: %v", err)
			}
			result = append(result, micro.Transport(t))
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

func ProvideMicroService(opts *mWire.Options, c cli.CLI, mOpts []micro.Option) (micro.Service, error) {
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
