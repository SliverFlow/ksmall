package data

import (
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/data/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	// init db and rdb
	NewDB,
	NewRDB,

	// init user repo
	NewUserRepo,
	wire.Bind(new(repo.IUserRepo), new(*userRepo)),
)
