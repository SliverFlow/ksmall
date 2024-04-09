package api

import (
	"github.com/SliverFlow/ksmall/core/server"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewApiGroup,
	wire.Bind(new(server.ApiGroup), new(*Group)),
)
