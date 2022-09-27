package generic

import (
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/transport"
)

var Brokers = NewContainer(func(opts ...broker.Option) broker.Broker { return nil })
var Registries = NewContainer(func(opts ...registry.Option) registry.Registry { return nil })
var Transports = NewContainer(func(opts ...transport.Option) transport.Transport { return nil })
