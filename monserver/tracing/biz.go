package tracing

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

type Biz struct {
}

func (b *Biz) Tacker(c context.Context, name string) (context.Context, trace.Span) {
	tracer := otel.Tracer("service")
	ctx, span := tracer.Start(c, name)
	return ctx, span
}
