package store

import (
	"github.com/go-micro/microwire/util/generic"
	"go-micro.dev/v4/store"
)

var Plugins = generic.NewContainer(func(opts ...store.Option) store.Store { return nil })
