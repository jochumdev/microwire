package store

import (
	"github.com/go-micro/microwire/v5/util/generic"
)

var Plugins = generic.NewContainer(func(opts ...Option) Store { return nil })
