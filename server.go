package leerpc0

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/encircles/leerpc0/interceptor"
	"github.com/encircles/leerpc0/logger"
	"github.com/encircles/leerpc0/plugin"
)

type Server struct {
	opts    *ServerOptions
	service Service
	plugins []plugin.Plugin

	closing bool // whether the server is closing
}

func NewServer(opt ...ServerOption) *Server {
	s := &Server{
		opts: &ServerOptions{},
	}

	for _, o := range opt {
		o(s.opts)
	}

	s.service = NewService(s.opts)

	for name, p := range plugin.PluginMap {
		if !containPlugin(name, s.opts.pluginNames) {
			continue
		}
		s.plugins = append(s.plugins, p)
	}

	return s
}

func NewService(opts *ServerOptions) Service {
	return &service{
		opts: opts,
	}
}

func containPlugin(pluginName string, plugins []string) bool {
	for _, p := range plugins {
		if pluginName == p {
			return true
		}
	}

	return false
}

type emptyInterface interface{}

func (s *Server) RegisterService(serviceName string, svr interface{}) error {
	svrType := reflect.TypeOf(svr)
	svrValue := reflect.ValueOf(svr)

	sd := &ServiceDesc{
		Svr:         svr,
		ServiceName: serviceName,
		// for compatibility with code generation
		HandlerType: (*emptyInterface)(nil),
	}

	methods, err := getServiceMethods(svrType, svrValue)
	if err != nil {
		return err
	}

	sd.Methods = methods

	s.Register(sd, svr)

	return nil
}

func getServiceMethods(serviceType reflect.Type, serviceValue reflect.Value) ([]*MethodDesc, error) {

	var methods []*MethodDesc

	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)

		if err := checkMethod(method.Type); err != nil {
			return nil, err
		}

		methodHandler := func(ctx context.Context, svr interface{}, dec func(interface{}) error, ceps []interceptor.ServerInterceptor) (interface{}, error) {
			reqType := method.Type.In(2)

			// determine type
			req := reflect.New(reqType.Elem()).Interface()

			if err := dec(req); err != nil {
				return nil, err
			}

			if len(ceps) == 0 {
				values := method.Func.Call([]reflect.Value{
					serviceValue,
					reflect.ValueOf(ctx),
					reflect.ValueOf(req),
				})
				// determine error
				return values[0].Interface(), nil
			}

			handler := func(ctx context.Context, reqBody interface{}) (interface{}, error) {
				values := method.Func.Call([]reflect.Value{
					serviceValue,
					reflect.ValueOf(ctx),
					reflect.ValueOf(req),
				})

				return values[0].Interface(), nil
			}

			return interceptor.ServerIntercept(ctx, req, ceps, handler)
		}

		methods = append(methods, &MethodDesc{
			MethodName: method.Name,
			Handler:    methodHandler,
		})
	}

	return methods, nil
}

func checkMethod(method reflect.Type) error {

	// params num must >= 2 , needs to be combined with itself
	if method.NumIn() < 3 {
		return fmt.Errorf("method %s invalid, the number of params < 2", method.Name())
	}

	// return values nums must be 2
	if method.NumOut() != 2 {
		return fmt.Errorf("method %s invalid, the number of return values != 2", method.Name())
	}

	// the first parameter must be context
	ctxType := method.In(1)
	var contextType = reflect.TypeOf((*context.Context)(nil)).Elem()
	if !ctxType.Implements(contextType) {
		return fmt.Errorf("method %s invalid, first param is not context", method.Name())
	}

	// the second parameter type must be pointer
	replyType := method.Out(0)
	if replyType.Kind() != reflect.Ptr {
		return fmt.Errorf("method %s invalid, reply type is not a pointer", method.Name())
	}

	// the second return value must be an error
	errType := method.Out(1)
	var errorType = reflect.TypeOf((*error)(nil)).Elem()
	if !errType.Implements(errorType) {
		return fmt.Errorf("method %s invalid, return %s, not error", method.Name(), errType.Name())
	}

	return nil
}

func (s *Server) Register(sd *ServiceDesc, svr interface{}) {
	if sd == nil || svr == nil {
		return
	}
	ht := reflect.TypeOf(sd.HandlerType).Elem()
	st := reflect.TypeOf(svr)
	if !st.Implements(ht) {
		logger.Fatalf("handlerType %v not match service : %v ", ht, st)
	}

	ser := &service{
		svr:         svr,
		serviceName: sd.ServiceName,
		handlers:    make(map[string]Handler),
	}

	for _, method := range sd.Methods {
		ser.handlers[method.MethodName] = method.Handler
	}

	s.service = ser
}

func (s *Server) Serve() {
	err := s.InitPlugins()
	if err != nil {
		panic(err)
	}

	s.service.Serve(s.opts)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGSEGV)
	<-ch

	s.Close()
}

type emptyService struct{}

func (s *Server) ServeHttp() {
	if err := s.RegisterService("/http", new(emptyService)); err != nil {
		panic(err)
	}

	s.Serve()
}

func (s *Server) Close() {
	s.closing = false

	s.service.Close()
}

func (s *Server) InitPlugins() error {
	// init plugins
	for _, p := range s.plugins {
		switch val := p.(type) {

		case plugin.ResolverPlugin:
			var services []string
			services = append(services, s.service.Name())

			pluginOpts := []plugin.Option{
				plugin.WithSelectorSvrAddr(s.opts.selectorSvrAddr),
				plugin.WithSvrAddr(s.opts.address),
				plugin.WithServices(services),
			}
			if err := val.Init(pluginOpts...); err != nil {
				logger.Errorf("resolver init error, %v", err)
				return err
			}

		case plugin.TracingPlugin:
			pluginOpts := []plugin.Option{
				plugin.WithTracingSvrAddr(s.opts.tracingSvrAddr),
			}

			_, err := val.Init(pluginOpts...)
			if err != nil {
				logger.Errorf("tracing init error, %v", err)
				return err
			}

		// TODO jaeger
		// s.opts.interceptors = append(s.opts.interceptors, jaeger.OpenTracingServerInterceptor(tracer, s.opts.tracingSpanName))

		default:

		}
	}

	return nil
}
