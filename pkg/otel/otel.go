package otel

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

const (
	NoTracer = ""
)

func Start(ctx context.Context, tracer, span string, opts ...trace.SpanStartOption) (context.Context, trace.Span) {
	return otel.Tracer(tracer).Start(ctx, span, opts...)
}
