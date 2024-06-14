package data

import (
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewDB,
	NewRDB,
	NewCommon,
	NewEtcd,

	NewRoleRepo,
	wire.Bind(new(repo.RoleRepo), new(*roleRepo)),
	NewUserRepo,
	wire.Bind(new(repo.UserRepo), new(*userRepo)),
	NewCategoryRepo,
	wire.Bind(new(repo.CategoryRepo), new(*categoryRepo)),
	NewGoodRepo,
	wire.Bind(new(repo.GoodRepo), new(*goodRepo)),
	NewStockRepo,
	wire.Bind(new(repo.StockRepo), new(*stockRepo)),
)
