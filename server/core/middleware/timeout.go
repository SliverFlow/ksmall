package middleware

import (
	"github.com/SliverFlow/ksmall/server/common/constant"
	"github.com/SliverFlow/ksmall/server/common/response"
	"github.com/SliverFlow/ksmall/server/common/zerror"
	"github.com/SliverFlow/ksmall/server/core/config"
	gtimeout "github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type Timeout struct {
	logger *zap.Logger
	config *config.HttpServer
}

// NewTimeoutMiddleware 创建超时中间件
func NewTimeoutMiddleware(logger *zap.Logger, c *config.HttpServer) *Timeout {
	return &Timeout{
		logger: logger,
		config: c,
	}
}

func (t *Timeout) Handle() gin.HandlerFunc {
	tot := t.config.Timeout
	if tot <= 2 || tot >= 30 {
		tot = 5
	}
	return gtimeout.New(
		gtimeout.WithTimeout(time.Duration(tot)*time.Second),
		gtimeout.WithHandler(func(c *gin.Context) {
			c.Next()
		}),
		gtimeout.WithResponse(func(c *gin.Context) {
			t.logger.Error("Request timeout", zap.String("path", c.Request.URL.Path))
			response.FailWithError(zerror.NewWithCode(constant.RequestTimeoutCode), c)
		}),
	)
}
