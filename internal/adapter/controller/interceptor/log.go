package interceptor

import (
	"context"
	"time"

	"github.com/pillarion/practice-platform/pkg/logger"
	"google.golang.org/grpc"
)

// LogInterceptor intercepts gRPC requests and logs them.
func (i *Interceptor) LogInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	now := time.Now()

	res, err := handler(ctx, req)
	if err != nil {
		logger.Err(err).Str("method", info.FullMethod).Any("req", req)
	}

	logger.Info().Str("method", info.FullMethod).Any("req", req).Any("res", res).Dur("duration", time.Since(now)).Msg("request")

	return res, err
}
