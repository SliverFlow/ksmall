// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/SliverFlow/ksmall/app/order/api"
	"github.com/SliverFlow/ksmall/app/order/internal/config"
	"github.com/SliverFlow/ksmall/core/server"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(c *config.Possess, log *zap.Logger) *server.HttpServer {
	configServer := config.NewServerConfig(c)
	group := api.NewApiGroup()
	httpServer := server.NewHttpServer(log, configServer, group)
	return httpServer
}
