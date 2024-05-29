package data

import (
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDB,
	NewRDB,
	NewCommon,

	NewRoleRepo,
	wire.Bind(new(repo.RoleRepo), new(*roleRepo)),
)
