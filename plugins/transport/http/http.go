// Package http returns a http2 transport using net/http
package http

import (
	"github.com/go-micro/microwire/v5/transport"
	mTransport "github.com/go-micro/microwire/v5/transport"
)

// The only reason this lives here atm.
func init() {
	mTransport.Plugins.Add("http", NewTransport)
}

// NewTransport returns a new http transport using net/http and supporting http2
func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewHTTPTransport(opts...)
}
