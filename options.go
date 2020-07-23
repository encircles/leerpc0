package leerpc0

import (
	"time"

	"github.com/encircles/leerpc0/interceptor"
)

type ServerOptions struct {
	address           string        // listening address, e.g. :( ip://127.0.0.1:8080、 dns://www.google.com)
	network           string        // network type, e.g. : tcp、udp
	protocol          string        // protocol typpe, e.g. : proto、json
	timeout           time.Duration // timeout
	serializationType string        // serialization type, default: proto

	selectorSvrAddr string   // service discovery server address, required when using the third-party service discovery plugin
	tracingSvrAddr  string   // tracing plugin server address, required when using the third-party tracing plugin
	tracingSpanName string   // tracing span name, required when using the third-party tracing plugin
	pluginNames     []string // plugin name
	interceptors    []interceptor.ServerInterceptor
}

type ServerOption func(*ServerOptions)

func WithAddress(address string) ServerOption {
	return func(o *ServerOptions) {
		o.address = address
	}
}

func WithNetwork(network string) ServerOption {
	return func(o *ServerOptions) {
		o.network = network
	}
}

func WithProtocol(protocol string) ServerOption {
	return func(o *ServerOptions) {
		o.protocol = protocol
	}
}

func WithTimeout(timeout time.Duration) ServerOption {
	return func(o *ServerOptions) {
		o.timeout = timeout
	}
}

func WithSerializationType(serializationType string) ServerOption {
	return func(o *ServerOptions) {
		o.serializationType = serializationType
	}
}

func WithSelectorSvrAddr(selectorSvrAddr string) ServerOption {
	return func(o *ServerOptions) {
		o.selectorSvrAddr = selectorSvrAddr
	}
}

func WithTracingSvrAddr(tracingSvrAddr string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSvrAddr = tracingSvrAddr
	}
}

func WithTracingSpanName(tracingSpanName string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSpanName = tracingSpanName
	}
}

func WithPluginNames(pluginName ...string) ServerOption {
	return func(o *ServerOptions) {
		o.pluginNames = append(o.pluginNames, pluginName...)
	}
}

func WithInterceptors(interceptors ...interceptor.ServerInterceptor) ServerOption {
	return func(o *ServerOptions) {
		o.interceptors = append(o.interceptors, interceptors...)
	}
}
