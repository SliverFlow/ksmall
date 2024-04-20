package api

import (
	"github.com/SliverFlow/ksmall/server/app/user/http/api/v1"
	"github.com/SliverFlow/ksmall/server/core/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	v1.NewApiGroup,
	wire.Bind(new(server.ApiGroup), new(*v1.Group)),
)
