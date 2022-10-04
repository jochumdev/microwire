package micro

import (
	"github.com/go-micro/microwire/v5/auth"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
	"github.com/go-micro/plugins/v4/config/encoder/yaml"
)

type configStruct struct {
	Auth   *auth.Config   `json:"auth,omitempty" yaml:"Auth,omitempty"`
	Broker *broker.Config `json:"broker,omitempty" yaml:"Broker,omitempty"`
	// Cache     *cache.Config     `json:"cache,omitempty" yaml:"Cache,omitempty"`
	Client    *client.Config    `json:"client,omitempty" yaml:"Client,omitempty"`
	Server    *server.Config    `json:"server,omitempty" yaml:"Server,omitempty"`
	Store     *store.Config     `json:"store,omitempty" yaml:"Store,omitempty"`
	Logger    *logger.Config    `json:"logger,omitempty" yaml:"Logger,omitempty"`
	Registry  *registry.Config  `json:"registry,omitempty" yaml:"Registry,omitempty"`
	Transport *transport.Config `json:"transport,omitempty" yaml:"Transport,omitempty"`
	// Runtime   *runtime.Config `json:"runtime,omitempty" yaml:"Runtime,omitempty"`
	// Profile   *profile.Config `json:"profile,omitempty" yaml:"Profile,omitempty"`
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
	cfg.Logger = s.Options().Logger.Options().Config
	cfg.Registry = s.Options().Registry.Options().Config
	cfg.Transport = s.Options().Transport.Options().Config

	return yaml.NewEncoder().Encode(cfg)
}
