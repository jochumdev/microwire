package broker

import (
	"strings"

	"github.com/go-micro/microwire/generic"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/util/cmd"
)

type BrokerOptions struct {
	Name      string
	Addresses string
}

func ProvideBroker(opts *BrokerOptions) (broker.Broker, error) {
	b, err := generic.Brokers.Get(opts.Name)
	if err != nil {
		var ok bool
		if b, ok = cmd.DefaultBrokers[opts.Name]; !ok {
			return nil, err
		}
	}

	var result broker.Broker
	if len(opts.Addresses) > 0 {
		result = b(broker.Addrs(strings.Split(opts.Addresses, ",")...))
	} else {
		result = b()
	}

	return result, nil
}
