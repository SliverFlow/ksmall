// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/SliverFlow/ksmall/app/user/http/api/v1"
	"github.com/SliverFlow/ksmall/app/user/http/internal/biz"
	"github.com/SliverFlow/ksmall/app/user/http/internal/config"
	"github.com/SliverFlow/ksmall/app/user/http/internal/rpc_server"
	"github.com/SliverFlow/ksmall/app/user/http/internal/service"
	"github.com/SliverFlow/ksmall/core/middleware"
	"github.com/SliverFlow/ksmall/core/server"
	"go.uber.org/zap"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(c *config.Possess, log *zap.Logger) *server.HttpServer {
	httpServer := config.NewServerConfig(c)
	userClient := rpc_server.NewUserRpc(c)
	userUsecase := biz.NewUserUsecase(log, userClient)
	userService := service.NewUserService(log, userUsecase)
	cors := middleware.NewCorsMiddleware(log)
	timeout := middleware.NewTimeoutMiddleware(log, httpServer)
	group := v1.NewApiGroup(userService, cors, timeout)
	serverHttpServer := server.NewHttpServer(log, httpServer, group)
	return serverHttpServer
}
