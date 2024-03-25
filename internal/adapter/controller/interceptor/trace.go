package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

func (i *Interceptor) TraceInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	ctx, span := i.tracer.Start(ctx, info.FullMethod)
	defer span.End()

	res, err := handler(ctx, req)
	if err != nil {
		span.RecordError(err)
	}

	return res, err
}
