//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/biz"
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/config"
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/data"
	"github.com/SliverFlow/ksmall/server/core/server"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func wireApp(c *config.Possess, logger *zap.Logger) *server.GrpcServer {
	panic(wire.Build(server.NewGrpcServer, config.NewServerConfig, biz.ProviderSet, data.ProviderSet))
}
