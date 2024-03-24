//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SliverFlow/ksmall/app/order/api"
	"github.com/SliverFlow/ksmall/app/order/internal/config"
	"github.com/SliverFlow/ksmall/core/server"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// wireApp init kratos application.
func wireApp(c *config.Possess, log *zap.Logger) *server.HttpServer {

	panic(wire.Build(api.ProviderSet, server.NewHttpServer, config.NewServerConfig))
}
