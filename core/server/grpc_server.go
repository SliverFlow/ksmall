package server

import (
	"fmt"
	"github.com/SliverFlow/ksmall/core/config"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type GrpcServer struct {
	server *grpc.Server
	c      *config.RpcServer
	logger *zap.Logger
}

type RpcRegister interface {
	RegisterServer(server *grpc.Server) *grpc.Server
}

func NewGrpcServer(c *config.RpcServer, logger *zap.Logger, rpcReg RpcRegister) *GrpcServer {
	return &GrpcServer{
		c:      c,
		logger: logger,
		server: rpcReg.RegisterServer(grpc.NewServer()),
	}
}

func (s *GrpcServer) Listen() error {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", s.c.Port))
	if err != nil {
		return err
	}
	err = s.server.Serve(listen)
	if err != nil {
		return err
	}
	return nil
}
