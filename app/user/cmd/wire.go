//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SliverFlow/ksmall/app/user/api"
	"github.com/SliverFlow/ksmall/app/user/internal/biz"
	"github.com/SliverFlow/ksmall/app/user/internal/config"
	"github.com/SliverFlow/ksmall/app/user/internal/service"
	"github.com/SliverFlow/ksmall/core/server"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// wireApp init kratos application.
func wireApp(c *config.Possess, log *zap.Logger) *server.HttpServer {
	panic(wire.Build(api.ProviderSet, service.ProviderSet, biz.ProviderSet, server.NewHttpServer, config.NewServerConfig))
}
