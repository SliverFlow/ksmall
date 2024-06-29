package api

import (
	"github.com/SliverFlow/core/middleware"
	"github.com/SliverFlow/core/server"
	v1 "github.com/SliverFlow/ksmall/monserver/api/v1"
	"github.com/SliverFlow/ksmall/monserver/api/v1/mall"
	"github.com/SliverFlow/ksmall/monserver/api/v1/system"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	middleware.NewTacker,
	middleware.NewAuthority,
	v1.NewGroup,
	wire.Bind(new(server.ApiGroup), new(*v1.Group)),
	system.NewGroup,
	mall.NewGroup,
)
