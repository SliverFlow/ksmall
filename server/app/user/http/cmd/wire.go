//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SliverFlow/ksmall/server/app/user/http/api"
	"github.com/SliverFlow/ksmall/server/app/user/http/internal/biz"
	"github.com/SliverFlow/ksmall/server/app/user/http/internal/config"
	"github.com/SliverFlow/ksmall/server/app/user/http/internal/rpc_client"
	"github.com/SliverFlow/ksmall/server/app/user/http/internal/service"
	"github.com/SliverFlow/ksmall/server/core/middleware"
	"github.com/SliverFlow/ksmall/server/core/server"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// wireApp init kratos application.
func wireApp(c *config.Possess, log *zap.Logger) *server.HttpServer {
	panic(wire.Build(middleware.ProviderSet, api.ProviderSet, service.ProviderSet, biz.ProviderSet, rpc_client.ProviderSet, server.NewHttpServer, config.NewServerConfig))
}
