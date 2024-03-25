package interceptor

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Interceptor intercepts gRPC requests
type Interceptor struct {
	tracer trace.Tracer
}

// NewInterceptor returns new interceptor
func NewInterceptor(tracename string) *Interceptor {
	t := otel.Tracer(tracename)

	return &Interceptor{
		tracer: t,
	}
}
