package interceptor

import "context"

type ServerInterceptor func(ctx context.Context, req interface{}, handler Handler) (interface{}, error)

type Handler func(ctx context.Context, req interface{}) (interface{}, error)

func ServerIntercept(ctx context.Context, req interface{}, ceps []ServerInterceptor, handler Handler) (interface{}, error) {
	if len(ceps) == 0 {
		return handler(ctx, req)
	}

	return ceps[0](ctx, req, getHandler(0, ceps, handler))
}

func getHandler(cur int, ceps []ServerInterceptor, handler Handler) Handler {
	if cur == len(ceps)-1 {
		return handler
	}

	return func(ctx context.Context, req interface{}) (i interface{}, err error) {
		return ceps[cur+1](ctx, req, getHandler(cur+1, ceps, handler))
	}
}
