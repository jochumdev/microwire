package urfave

import (
	"flag"
	"fmt"

	mCli "github.com/go-micro/microwire/cli"
	"github.com/urfave/cli/v2"
)

func init() {
	mCli.Container.Add("urfave", NewCLI)
}

type FlagCLI struct {
	stringFlags map[string]cli.StringFlag
	intFlags    map[string]cli.IntFlag
	options     *mCli.Options
	ctx         *cli.Context
}

func NewCLI(opts ...mCli.Option) mCli.CLI {
	return &FlagCLI{
		stringFlags: make(map[string]cli.StringFlag),
		intFlags:    make(map[string]cli.IntFlag),
		options:     mCli.NewCLIOptions(),
	}
}

func (c *FlagCLI) AddInt(opts ...mCli.FlagOption) error {
	options := mCli.NewFlag(opts...)

	c.intFlags[options.Name] = cli.IntFlag{
		Name:    options.Name,
		Usage:   options.Usage,
		Value:   options.DefaultInt,
		EnvVars: options.EnvVars,
	}

	return nil
}

func (c *FlagCLI) AddString(opts ...mCli.FlagOption) error {
	options := mCli.NewFlag(opts...)

	c.stringFlags[options.Name] = cli.StringFlag{
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

	set := flag.NewFlagSet(c.options.Name, flag.ContinueOnError)

	for _, f := range c.stringFlags {
		if err := f.Apply(set); err != nil {
			return err
		}
	}

	for _, f := range c.intFlags {
		if err := f.Apply(set); err != nil {
			return err
		}
	}

	app := &cli.App{
		Version:     c.options.Version,
		Description: c.options.Description,
		Usage:       c.options.Usage,
	}
	if len(c.options.Version) < 1 {
		app.HideVersion = true
	}
	c.ctx = cli.NewContext(app, set, nil)

	if err := set.Parse(args); err != nil {
		return err
	}
	if len(set.Args()) > 1 {
		return fmt.Errorf("unknown flags '%v' given", set.Args())
	}

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
