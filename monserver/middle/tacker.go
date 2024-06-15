package middle

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
)

type Tacker struct {
}

func NewTacker() *Tacker {
	return &Tacker{}
}

func (t *Tacker) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		tracer := otel.Tracer("ParentTracer")
		ctx, span := tracer.Start(c.Request.Context(), c.Request.URL.Path)
		defer span.End()

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
