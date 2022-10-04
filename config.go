package micro

import (
	"github.com/go-micro/microwire/v5/auth"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
)

type configStruct struct {
	Auth   *auth.Config   `json:",omitempty" yaml:",omitempty"`
	Broker *broker.Config `json:",omitempty" yaml:",omitempty"`
	// Cache     *cache.Config     `json:",omitempty" yaml:",omitempty"`
	Client    *client.Config    `json:",omitempty" yaml:",omitempty"`
	Server    *server.Config    `json:",omitempty" yaml:",omitempty"`
	Store     *store.Config     `json:",omitempty" yaml:",omitempty"`
	Registry  *registry.Config  `json:",omitempty" yaml:",omitempty"`
	Transport *transport.Config `json:",omitempty" yaml:",omitempty"`
	// Runtime   *runtime.Config `json:",omitempty" yaml:",omitempty"`
	// Profile   *profile.Config `json:",omitempty" yaml:",omitempty"`
	// Logger    *logger.Config `json:",omitempty" yaml:",omitempty"`
}

// DumpConfig dumps the config of a micro.Service to a yaml byte array.
func (s *service) DumpConfig() ([]byte, error) {
	cfg := configStruct{}
	cfg.Auth = s.Options().Auth.Options().Config
	cfg.Broker = s.Options().Broker.Options().Config
	// cfg.Cache = service.Options().Cache.Options().Config
	cfg.Client = s.Client().Options().Config
	cfg.Server = s.Server().Options().Config
	cfg.Store = s.Options().Store.Options().Config
	cfg.Registry = s.Options().Registry.Options().Config
	cfg.Transport = s.Options().Transport.Options().Config

	return yaml.NewEncoder().Encode(cfg)
}
