package biz

import (
	"github.com/SliverFlow/core/middleware"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEtcdLocker,
	NewRoleUsecase,
	wire.Bind(new(middleware.IAuthority), new(*RoleUsecase)),
	NewUserUsecase,
	NewCategoryUsecase,
	NewGoodUsecase,
	NewAuthorityUsecase,
	NewAuthorityGroupUsecase,
)
