package broker

import (
	"github.com/go-micro/microwire/util/generic"
	"go-micro.dev/v4/broker"
)

var Container = generic.NewContainer(func(opts ...broker.Option) broker.Broker { return nil })
