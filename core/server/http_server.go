package server

import (
	"fmt"
	"github.com/SliverFlow/ksmall/core/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"time"
)

type HttpServer struct {
	Server *http.Server
}

func NewHttpServer(logger *zap.Logger, c *config.Server, api ApiGroup) *HttpServer {

	R := gin.Default()
	api.InitApi(R)

	time.Sleep(500 * time.Millisecond)
	logger.Info(fmt.Sprintf("[项目运行于]：http://127.0.0.1:%d", c.Port))

	return &HttpServer{
		Server: &http.Server{
			Addr:           fmt.Sprintf(":%d", c.Port),
			Handler:        R,
			ReadTimeout:    20 * time.Second,
			WriteTimeout:   20 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

type ApiGroup interface {
	InitApi(r *gin.Engine)
}
