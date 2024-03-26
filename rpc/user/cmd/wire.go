//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SliverFlow/ksmall/core/server"
	"github.com/SliverFlow/ksmall/rpc/user/internal/config"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func wireApp(c *config.Possess, logger *zap.Logger) *server.GrpcServer {
	panic(wire.Build(server.NewGrpcServer, config.NewServerConfig))
}
