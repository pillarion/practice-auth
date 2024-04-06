package interceptor

import (
	"context"

	"google.golang.org/grpc"
)

// RateLimiter intercepts gRPC requests and limits them.
func (i *Interceptor) RateLimiter(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

	return handler(ctx, req)
}
