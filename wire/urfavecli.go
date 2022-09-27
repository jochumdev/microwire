package wire

import (
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"
)

type InternalFlags []cli.Flag

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

func ProvideMyFlags(opts *Options) InternalFlags {
	result := []cli.Flag{}
	result = append(result, opts.Flags...)
	return InternalFlags(result)
}

func ProvideMicroService(opts *Options, ctx *cli.Context, mOpts []micro.Option) micro.Service {
	return micro.NewService(
		mOpts...,
	)
}
