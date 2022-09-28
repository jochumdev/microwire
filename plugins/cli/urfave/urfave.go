package urfave

import (
	mCli "github.com/go-micro/microwire/cli"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4/errors"
)

func init() {
	mCli.Plugins.Add("urfave", NewCLI)
}

type FlagCLI struct {
	flags       map[string]cli.Flag
	stringFlags map[string]*cli.StringFlag
	intFlags    map[string]*cli.IntFlag
	options     *mCli.Options
	ctx         *cli.Context
}

func NewCLI(opts ...mCli.Option) mCli.CLI {
	return &FlagCLI{
		flags:       make(map[string]cli.Flag),
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
			Name:    options.Name,
			Usage:   options.Usage,
			Value:   options.DefaultInt,
			EnvVars: options.EnvVars,
		}
		c.flags[options.Name] = f
		c.intFlags[options.Name] = f
	case mCli.FlagTypeString:
		f := &cli.StringFlag{
			Name:    options.Name,
			Usage:   options.Usage,
			Value:   options.DefaultString,
			EnvVars: options.EnvVars,
		}
		c.flags[options.Name] = f
		c.stringFlags[options.Name] = f
	default:
		return errors.InternalServerError("USER_FLAG_WITHOUT_A_DEFAULTOPTION", "found a flag without a default option")
	}

	return nil
}

func (c *FlagCLI) Init(args []string, opts ...mCli.Option) error {
	for _, o := range opts {
		o(c.options)
	}

	i := 0
	flags := make([]cli.Flag, len(c.flags))
	for _, f := range c.flags {
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

func (c *FlagCLI) StringValue(name string) string {
	flag, ok := c.stringFlags[name]
	if !ok {
		return ""
	}

	return flag.Get(c.ctx)
}

func (c *FlagCLI) IntValue(name string) int {
	flag, ok := c.intFlags[name]
	if !ok {
		return 0
	}

	return flag.Get(c.ctx)
}
