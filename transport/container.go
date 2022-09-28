package transport

import (
	"github.com/go-micro/microwire/util/generic"
	"go-micro.dev/v4/transport"
)

var Container = generic.NewContainer(func(opts ...transport.Option) transport.Transport { return nil })