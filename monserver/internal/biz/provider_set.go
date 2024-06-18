package biz

import (
	"github.com/SliverFlow/ksmall/monserver/middle"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewEtcdLocker,
	NewRoleUsecase,
	wire.Bind(new(middle.IAuthority), new(*RoleUsecase)),
	NewUserUsecase,
	NewCategoryUsecase,
	NewGoodUsecase,
	NewAuthorityUsecase,
	NewAuthorityGroupUsecase,
)
