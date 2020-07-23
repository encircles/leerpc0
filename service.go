package leerpc0

import (
	"context"
	"github.com/encircles/leerpc0/interceptor"
)

type Service interface {
	Register(string, Handler)
	Serve(*ServerOptions)
	Close()
}

type service struct {
	svr         interface{}
	ctx         context.Context
	cancel      context.CancelFunc
	serviceName string
	handlers    map[string]Handler
	opts        *ServerOptions
}

// Handler is the handler of a method
type Handler func(context.Context, interface{}, func(interface{}) error, []interceptor.ServerInterceptor) (interface{}, error)
