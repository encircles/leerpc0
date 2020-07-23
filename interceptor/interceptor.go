package interceptor

import "context"

type ServerInterceptor func(ctx context.Context, req interface{}, handler Handler) (interface{}, error)

type Handler func(ctx context.Context, req interface{}) (interface{}, error)
