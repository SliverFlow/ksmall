//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/SliverFlow/core/server"
	"github.com/SliverFlow/ksmall/monserver/api"
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/internal/config"
	"github.com/SliverFlow/ksmall/monserver/internal/data"
	"github.com/SliverFlow/ksmall/monserver/internal/service"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// wireApp init
func wireApp(c *config.Possess, log *zap.Logger) *server.Http {
	panic(wire.Build(config.NewHttpServerConfig, server.NewHttp, api.ProviderSet, service.ProviderSet, biz.ProviderSet, data.ProviderSet))
}
