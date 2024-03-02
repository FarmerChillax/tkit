package tkit

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func StartSpan(ctx context.Context, tracerName, spanName string) (context.Context, trace.Span) {
	return otel.Tracer(tracerName).Start(ctx, spanName)
}
