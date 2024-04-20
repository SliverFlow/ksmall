// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/biz"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/config"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/data"
	"github.com/SliverFlow/ksmall/core/server"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func wireApp(c *config.Possess, logger *zap.Logger) *server.GrpcServer {
	rpcServer := config.NewServerConfig(c)
	db := data.NewDB(c)
	client := data.NewRDB(c)
	userRepo := data.NewUserRepo(logger, db, client, c)
	userUsecase := biz.NewUserUsecase(userRepo, logger)
	grpcServer := server.NewGrpcServer(rpcServer, logger, userUsecase)
	return grpcServer
}