package leerpc0

import (
	"github.com/encircles/leerpc0/interceptor"
	"time"
)

type ServerOptions struct {
	address           string
	network           string
	protocol          string
	timeout           time.Duration
	serializationType string

	selectorSvrAddr string
	tracingSvrAddr  string
	tracingSpanName string
	pluginNames     []string
	interceptors    []interceptor.ServerInterceptor
}
