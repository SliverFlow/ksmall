package service

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	NewRoleService,
	NewUserService,
	NewCategoryService,
	NewGoodService,
	NewAuthorityService,
	NewAuthorityGroupService,
)
