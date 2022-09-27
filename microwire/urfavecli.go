package microwire

import (
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
)

type InternalFlags []cli.Flag

func ProvideDefaultServiceInitializer() InitializeServiceFunc {
	return DefaultService
}

func ProvideApp(opts *Options, flags InternalFlags) *cli.App {
	app := cli.NewApp()
	app.Name = opts.Name
	app.Version = opts.Version
	app.Usage = opts.Usage
	app.Flags = []cli.Flag(flags)

	app.Before = func(c *cli.Context) error {
		return nil
	}
	app.After = func(c *cli.Context) error {
		return nil
	}
	app.Action = func(c *cli.Context) error {
		s, err := opts.InitService(c, opts)
		if err != nil {
			return err
		}

		for _, fn := range opts.Actions {
			if err := fn(c, s); err != nil {
				return err
			}
		}

		return s.Run()
	}

	if len(opts.Name) == 0 {
		app.HideVersion = true
	}

	return app
}

func ProvideDefaultFlags(opts *Options, brokerFlags BrokerFlags) InternalFlags {
	result := []cli.Flag{}
	result = append(result, []cli.Flag(brokerFlags)...)
	result = append(result, opts.Flags...)
	return InternalFlags(result)
}

func ProvideMyFlags(opts *Options) InternalFlags {
	result := []cli.Flag{}
	result = append(result, opts.Flags...)
	return InternalFlags(result)
}

func ProvideDefaultMicroOpts(opts *Options, broker broker.Broker) []micro.Option {
	result := []micro.Option{
		micro.Name(opts.Name),
		micro.Version(opts.Version),
		micro.Broker(broker),
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

	return result
}

func ProvideMicroService(opts *Options, ctx *cli.Context, mOpts []micro.Option) micro.Service {
	return micro.NewService(
		mOpts...,
	)
}
