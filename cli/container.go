package cli

import (
	"github.com/go-micro/microwire/util/generic"
)

var Container = generic.NewContainer(func(opts ...Option) CLI { return nil })
