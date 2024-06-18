package system

import (
	"github.com/SliverFlow/ksmall/monserver/internal/service"
	"github.com/SliverFlow/ksmall/monserver/middle"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Group struct {
	logger                *zap.Logger
	authority             *middle.Authority
	roleService           *service.RoleService
	userService           *service.UserService
	categoryService       *service.CategoryService
	goodService           *service.GoodService
	authorityService      *service.AuthorityService
	authorityGroupService *service.AuthorityGroupService
}

func NewGroup(
	logger *zap.Logger,
	roleService *service.RoleService,
	userService *service.UserService,
	categoryService *service.CategoryService,
	goodService *service.GoodService,
	authorityService *service.AuthorityService,
	authorityGroupService *service.AuthorityGroupService,
	authority *middle.Authority,
) *Group {
	return &Group{
		logger:                logger,
		roleService:           roleService,
		userService:           userService,
		categoryService:       categoryService,
		goodService:           goodService,
		authorityService:      authorityService,
		authorityGroupService: authorityGroupService,
		authority:             authority,
	}
}

func (a *Group) InitApi(group *gin.RouterGroup) {
	// group.Use(a.authority.Handle())

	// 角色相关
	roleRouter := group.Group("/system/role")
	{
		roleRouter.POST("/create", a.roleService.Create)
		roleRouter.POST("/delete", a.roleService.Delete)
		roleRouter.POST("/update", a.roleService.Update)
		roleRouter.POST("/find", a.roleService.Find)
		roleRouter.POST("/list", a.roleService.List)
		roleRouter.POST("/dict", a.roleService.Dict)             // 角色字典
		roleRouter.POST("/statusDict", a.roleService.StatusDict) // 角色状态字典
		roleRouter.POST("/allocationAuth", a.roleService.AllocationAuth)

	}

	// 用户相关
	userRouter := group.Group("/system/user")
	{
		userRouter.POST("/create", a.userService.Create)
	}

	// 分类相关
	categoryRouter := group.Group("/system/category")
	{
		categoryRouter.POST("/create", a.categoryService.Create)

	}

	// 商品相关
	goodRouter := group.Group("/system/good")
	{
		goodRouter.POST("/create", a.goodService.Create)
	}

	// 权限相关
	authorityRouter := group.Group("/system/authority")
	{
		authorityRouter.POST("/create", a.authorityService.Create)
	}

	// 权限组相关
	authorityGroupRouter := group.Group("/system/authorityGroup")
	{
		authorityGroupRouter.POST("/create", a.authorityGroupService.Create)
		authorityGroupRouter.POST("/update", a.authorityGroupService.Update)
	}

	a.logger.Info("system api init success")
}
