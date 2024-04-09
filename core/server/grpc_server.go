package server

import (
	"github.com/SliverFlow/ksmall/core/config"
	"go.uber.org/zap"
)

type GrpcServer struct {
}

func NewGrpcServer(c *config.RpcServer, logger *zap.Logger) *GrpcServer {
	return &GrpcServer{}
}
