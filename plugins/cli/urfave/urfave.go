package urfave

import (
	mCli "github.com/go-micro/microwire/cli"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/errors"
)

func init() {
	_ = mCli.Plugins.Add("urfave", NewCLI)
}

type FlagCLI struct {
	stringFlags map[string]*cli.StringFlag
	intFlags    map[string]*cli.IntFlag
	options     *mCli.Options
	ctx         *cli.Context
}

func NewCLI(opts ...mCli.Option) mCli.Cli {
	return &FlagCLI{
		stringFlags: make(map[string]*cli.StringFlag),
		intFlags:    make(map[string]*cli.IntFlag),
		options:     mCli.NewCLIOptions(),
	}
}

func (c *FlagCLI) Add(opts ...mCli.FlagOption) error {
	options, err := mCli.NewFlag(opts...)
	if err != nil {
		return err
	}

	switch options.FlagType {
	case mCli.FlagTypeInt:
		f := &cli.IntFlag{
			Name:        options.Name,
			Usage:       options.Usage,
			Value:       options.DefaultInt,
			EnvVars:     options.EnvVars,
			Destination: options.DestinationInt,
		}
		c.intFlags[options.Name] = f
	case mCli.FlagTypeString:
		f := &cli.StringFlag{
			Name:        options.Name,
			Usage:       options.Usage,
			Value:       options.DefaultString,
			EnvVars:     options.EnvVars,
			Destination: options.DestinationString,
		}
		c.stringFlags[options.Name] = f
	default:
		return errors.InternalServerError("USER_FLAG_WITHOUT_A_DEFAULTOPTION", "found a flag without a default option")
	}

	return nil
}

func (c *FlagCLI) Parse(args []string, opts ...mCli.Option) error {
	for _, o := range opts {
		o(c.options)
	}

	i := 0
	flags := make([]cli.Flag, len(c.stringFlags)+len(c.intFlags))
	for _, f := range c.stringFlags {
		flags[i] = f
		i += 1
	}
	for _, f := range c.intFlags {
		flags[i] = f
		i += 1
	}

	var ctx *cli.Context
	app := &cli.App{
		Version:     c.options.Version,
		Description: c.options.Description,
		Usage:       c.options.Usage,
		Flags:       flags,
		Action: func(fCtx *cli.Context) error {
			// Extract the ctx from the urfave app
			ctx = fCtx

			return nil
		},
	}
	if len(c.options.Version) < 1 {
		app.HideVersion = true
	}

	if err := app.Run(args); err != nil {
		return err
	}
	c.ctx = ctx

	return nil
}

func (c *FlagCLI) String() string {
	return "urfave"
}
