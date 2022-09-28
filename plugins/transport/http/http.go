// Package http returns a http2 transport using net/http
package http

import (
	mTransport "github.com/go-micro/microwire/transport"
	"go-micro.dev/v4/transport"
)

func init() {
	mTransport.Plugins.Add("http", NewTransport)
}

// NewTransport returns a new http transport using net/http and supporting http2
func NewTransport(opts ...transport.Option) transport.Transport {
	return transport.NewHTTPTransport(opts...)
}
