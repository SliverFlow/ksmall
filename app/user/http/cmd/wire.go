//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SliverFlow/ksmall/app/user/http/api"
	"github.com/SliverFlow/ksmall/app/user/http/internal/biz"
	"github.com/SliverFlow/ksmall/app/user/http/internal/config"
	"github.com/SliverFlow/ksmall/app/user/http/internal/rpc_server"
	"github.com/SliverFlow/ksmall/app/user/http/internal/service"
	"github.com/SliverFlow/ksmall/core/middleware"
	"github.com/SliverFlow/ksmall/core/server"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// wireApp init kratos application.
func wireApp(c *config.Possess, log *zap.Logger) *server.HttpServer {
	panic(wire.Build(middleware.ProviderSet, api.ProviderSet, service.ProviderSet, biz.ProviderSet, rpc_server.ProviderSet, server.NewHttpServer, config.NewServerConfig))
}
