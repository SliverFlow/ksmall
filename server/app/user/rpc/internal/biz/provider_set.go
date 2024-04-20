package biz

import (
	"github.com/SliverFlow/ksmall/server/core/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewUserUsecase,

	wire.Bind(new(server.RpcRegister), new(*UserUsecase)),
)
