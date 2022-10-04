package micro

import (
	"context"

	"github.com/go-micro/microwire/v5/auth"
	"github.com/go-micro/microwire/v5/broker"
	"github.com/go-micro/microwire/v5/cache"
	mCli "github.com/go-micro/microwire/v5/cli"
	"github.com/go-micro/microwire/v5/client"
	"github.com/go-micro/microwire/v5/config"
	"github.com/go-micro/microwire/v5/debug/profile"
	"github.com/go-micro/microwire/v5/debug/trace"
	"github.com/go-micro/microwire/v5/logger"
	"github.com/go-micro/microwire/v5/registry"
	"github.com/go-micro/microwire/v5/runtime"
	"github.com/go-micro/microwire/v5/selector"
	"github.com/go-micro/microwire/v5/server"
	"github.com/go-micro/microwire/v5/store"
	"github.com/go-micro/microwire/v5/transport"
)

type HookFunc func(Service) error

// Options for micro service.
type Options struct {
	ArgPrefix        string
	Address          string
	Name             string
	Description      string
	Version          string
	Usage            string
	NoFlags          bool
	ConfigFile       string
	RegisterTTL      int
	RegisterInterval int
	Metadata         map[string]string
	Flags            []mCli.Flag

	// References to get them through service.Xyz
	Auth      auth.Auth
	Broker    broker.Broker
	Cache     cache.Cache
	Config    config.Config
	Client    client.Client
	Server    server.Server
	Store     store.Store
	Registry  registry.Registry
	Runtime   runtime.Runtime
	Transport transport.Transport
	Profile   profile.Profile
	Logger    logger.Logger

	// Wrappers
	WrapSubscriber []server.SubscriberWrapper
	WrapHandler    []server.HandlerWrapper
	WrapCall       []client.CallWrapper
	WrapClient     []client.Wrapper
	OrigClient     client.Client

	// Before and After funcs
	Actions     []HookFunc
	BeforeStart []HookFunc
	BeforeStop  []HookFunc
	AfterStart  []HookFunc
	AfterStop   []HookFunc

	// Other options for implementations of the interface
	// can be stored in a context
	Context context.Context

	Signal bool
}

func NewOptions(opts ...Option) *Options {
	opt := &Options{
		ArgPrefix:        "",
		Address:          "",
		Name:             "",
		Description:      "",
		Version:          "",
		Usage:            "",
		NoFlags:          false,
		ConfigFile:       "",
		RegisterTTL:      30,
		RegisterInterval: 60,
		Metadata:         make(map[string]string),
		Flags:            []mCli.Flag{},

		Actions:     []HookFunc{},
		BeforeStart: []HookFunc{},
		BeforeStop:  []HookFunc{},
		AfterStart:  []HookFunc{},
		AfterStop:   []HookFunc{},

		Auth:    auth.DefaultAuth,
		Cache:   cache.DefaultCache,
		Config:  config.DefaultConfig,
		Store:   store.DefaultStore,
		Runtime: runtime.DefaultRuntime,
		Context: context.Background(),
		Signal:  true,
		Logger:  logger.DefaultLogger,
	}

	for _, o := range opts {
		o(opt)
	}

	return opt
}

// Broker to be used for service.
func Broker(b broker.Broker) Option {
	return func(o *Options) {
		o.Broker = b
	}
}

func Cache(c cache.Cache) Option {
	return func(o *Options) {
		o.Cache = c
	}
}

// Client to be used for service.
func Client(c client.Client) Option {
	return func(o *Options) {
		o.OrigClient = c

		// apply in reverse
		o.Client = c
		for i := len(o.WrapClient); i > 0; i-- {
			o.Client = o.WrapClient[i-1](o.Client)
		}
	}
}

// Context specifies a context for the service.
// Can be used to signal shutdown of the service and for extra option values.
func Context(ctx context.Context) Option {
	return func(o *Options) {
		o.Context = ctx
	}
}

// HandleSignal toggles automatic installation of the signal handler that
// traps TERM, INT, and QUIT.  Users of this feature to disable the signal
// handler, should control liveness of the service through the context.
func HandleSignal(b bool) Option {
	return func(o *Options) {
		o.Signal = b
	}
}

// Profile to be used for debug profile.
func Profile(p profile.Profile) Option {
	return func(o *Options) {
		o.Profile = p
	}
}

// Server to be used for service.
func Server(s server.Server) Option {
	return func(o *Options) {
		o.Server = s
	}
}

// Store sets the store to use.
func Store(s store.Store) Option {
	return func(o *Options) {
		o.Store = s
	}
}

// Registry sets the registry for the service
// and the underlying components.
func Registry(r registry.Registry) Option {
	return func(o *Options) {
		o.Registry = r
	}
}

// Tracer sets the tracer for the service.
func Tracer(t trace.Tracer) Option {
	return func(o *Options) {
		o.Server.Init(server.Tracer(t))
	}
}

// Auth sets the auth for the service.
func Auth(a auth.Auth) Option {
	return func(o *Options) {
		o.Auth = a
	}
}

// Config sets the config for the service.
func Config(c config.Config) Option {
	return func(o *Options) {
		o.Config = c
	}
}

// Selector sets the selector for the service client.
func Selector(s selector.Selector) Option {
	return func(o *Options) {
		o.Client.Init(client.Selector(s))
	}
}

// Transport sets the transport for the service
// and the underlying components.
func Transport(t transport.Transport) Option {
	return func(o *Options) {
		o.Transport = t
	}
}

// Runtime sets the runtime.
func Runtime(r runtime.Runtime) Option {
	return func(o *Options) {
		o.Runtime = r
	}
}

// Convenience options

// Address sets the address of the server.
func Address(addr string) Option {
	return func(o *Options) {
		o.Address = addr
	}
}

// Name of the service.
func Name(n string) Option {
	return func(o *Options) {
		o.Name = n
	}
}

// Version of the service.
func Version(v string) Option {
	return func(o *Options) {
		o.Version = v
	}
}

// ArgPrefix is the cli prefix for args.
func ArgPrefix(n string) Option {
	return func(o *Options) {
		o.ArgPrefix = n
	}
}

// Description is the Description in cli usage.
func Description(n string) Option {
	return func(o *Options) {
		o.Description = n
	}
}

// Usage is the Usage in cli.
func Usage(n string) Option {
	return func(o *Options) {
		o.Usage = n
	}
}

// NoFlags is a marker that no micro flags should be there.
func NoFlags() Option {
	return func(o *Options) {
		o.NoFlags = true
	}
}

// Flags is a list of additional flags to add you can parse them in Hooks.
func Flags(n []mCli.Flag) Option {
	return func(o *Options) {
		o.Flags = n
	}
}

// ConfigFile is the config file to read in.
func ConfigFile(n string) Option {
	return func(o *Options) {
		o.ConfigFile = n
	}
}

// Metadata associated with the service.
func Metadata(md map[string]string) Option {
	return func(o *Options) {
		o.Metadata = md
	}
}

// RegisterTTL specifies the TTL to use when registering the service.
func RegisterTTL(n int) Option {
	return func(o *Options) {
		o.RegisterTTL = n
	}
}

// RegisterInterval specifies the interval on which to re-register.
func RegisterInterval(n int) Option {
	return func(o *Options) {
		o.RegisterInterval = n
	}
}

// WrapClient is a convenience method for wrapping a Client with
// some middleware component. A list of wrappers can be provided.
// Wrappers are applied in reverse order so the last is executed first.
func WrapClient(w ...client.Wrapper) Option {
	return func(o *Options) {
		o.WrapClient = append(o.WrapClient, w...)

		// apply in reverse
		o.Client = o.OrigClient
		for i := len(o.WrapClient); i > 0; i-- {
			o.Client = o.WrapClient[i-1](o.Client)
		}
	}
}

// WrapCall is a convenience method for wrapping a Client CallFunc.
func WrapCall(w ...client.CallWrapper) Option {
	return func(o *Options) {
		o.WrapCall = append(o.WrapCall, w...)
	}
}

// WrapHandler adds a handler Wrapper to a list of options passed into the server.
func WrapHandler(w ...server.HandlerWrapper) Option {
	return func(o *Options) {
		o.WrapHandler = append(o.WrapHandler, w...)
	}
}

// WrapSubscriber adds a subscriber Wrapper to a list of options passed into the server.
func WrapSubscriber(w ...server.SubscriberWrapper) Option {
	return func(o *Options) {
		o.WrapSubscriber = append(o.WrapSubscriber, w...)
	}
}

// Before and Afters

// BeforeStart run funcs before service starts.
func BeforeStart(fn HookFunc) Option {
	return func(o *Options) {
		o.BeforeStart = append(o.BeforeStart, fn)
	}
}

// BeforeStop run funcs before service stops.
func BeforeStop(fn HookFunc) Option {
	return func(o *Options) {
		o.BeforeStop = append(o.BeforeStop, fn)
	}
}

// AfterStart run funcs after service starts.
func AfterStart(fn HookFunc) Option {
	return func(o *Options) {
		o.AfterStart = append(o.AfterStart, fn)
	}
}

// Action is an alias for AfterStart.
func Action(fn HookFunc) Option {
	return func(o *Options) {
		o.AfterStart = append(o.Actions, fn)
	}
}

// AfterStop run funcs after service stops.
func AfterStop(fn HookFunc) Option {
	return func(o *Options) {
		o.AfterStop = append(o.AfterStop, fn)
	}
}

// Logger sets the logger for the service.
func Logger(l logger.Logger) Option {
	return func(o *Options) {
		o.Logger = l
	}
}
