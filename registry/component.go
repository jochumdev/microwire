package registry

import (
	"github.com/go-micro/microwire/util/generic"
	"go-micro.dev/v4/registry"
)

var Plugins = generic.NewContainer(func(opts ...registry.Option) registry.Registry { return nil })
