package cli

import (
	"github.com/go-micro/microwire/util/generic"
)

const ComponentName = "cli"

var Plugins = generic.NewContainer(func(opts ...Option) CLI { return nil })
