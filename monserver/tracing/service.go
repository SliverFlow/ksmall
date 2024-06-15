package tracing

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
)

type Service struct{}

func (s *Service) Tacker(c *gin.Context, name string) (context.Context, trace.Span) {
	tracer := otel.Tracer("service")
	ctx, span := tracer.Start(c.Request.Context(), name)
	return ctx, span
}
