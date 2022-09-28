package urfave

import (
	mCli "github.com/go-micro/microwire/cli"
	"github.com/urfave/cli/v2"
)

func init() {
	mCli.Container.Add("urfave", NewCLI)
}

type FlagCLI struct {
	stringFlags map[string]*cli.StringFlag
	intFlags    map[string]*cli.IntFlag
	options     *mCli.Options
	ctx         *cli.Context
}

func NewCLI(opts ...mCli.Option) mCli.CLI {
	return &FlagCLI{
		stringFlags: make(map[string]*cli.StringFlag),
		intFlags:    make(map[string]*cli.IntFlag),
		options:     mCli.NewCLIOptions(),
	}
}

func (c *FlagCLI) AddInt(opts ...mCli.FlagOption) error {
	options := mCli.NewFlag(opts...)

	c.intFlags[options.Name] = &cli.IntFlag{
		Name:    options.Name,
		Usage:   options.Usage,
		Value:   options.DefaultInt,
		EnvVars: options.EnvVars,
	}

	return nil
}

func (c *FlagCLI) AddString(opts ...mCli.FlagOption) error {
	options := mCli.NewFlag(opts...)

	c.stringFlags[options.Name] = &cli.StringFlag{
		Name:    options.Name,
		Usage:   options.Usage,
		Value:   options.DefaultString,
		EnvVars: options.EnvVars,
	}

	return nil
}

func (c *FlagCLI) Init(args []string, opts ...mCli.Option) error {
	for _, o := range opts {
		o(c.options)
	}

	flags := []cli.Flag{}
	for _, f := range c.stringFlags {
		flags = append(flags, f)
	}

	for _, f := range c.intFlags {
		flags = append(flags, f)
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

func (c *FlagCLI) String(name string) string {
	flag, ok := c.stringFlags[name]
	if !ok {
		return ""
	}

	return flag.Get(c.ctx)
}

func (c *FlagCLI) Int(name string) int {
	flag, ok := c.intFlags[name]
	if !ok {
		return 0
	}

	return flag.Get(c.ctx)
}
