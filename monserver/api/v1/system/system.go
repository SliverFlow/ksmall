package system

import (
	"github.com/SliverFlow/ksmall/monserver/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Group struct {
	logger          *zap.Logger
	roleService     *service.RoleService
	userService     *service.UserService
	categoryService *service.CategoryService
	goodService     *service.GoodService
}

func NewGroup(
	logger *zap.Logger,
	roleService *service.RoleService,
	userService *service.UserService,
	categoryService *service.CategoryService,
	goodService *service.GoodService,
) *Group {
	return &Group{
		logger:          logger,
		roleService:     roleService,
		userService:     userService,
		categoryService: categoryService,
		goodService:     goodService,
	}
}

func (a *Group) InitApi(group *gin.RouterGroup) {
	roleRouter := group.Group("/system/role")
	{
		roleRouter.POST("/create", a.roleService.Create)
		roleRouter.POST("/delete", a.roleService.Delete)
		roleRouter.POST("/update", a.roleService.Update)
		roleRouter.POST("/find", a.roleService.Find)
		roleRouter.POST("/list", a.roleService.List)
		roleRouter.POST("/dict", a.roleService.Dict)             // 角色字典
		roleRouter.POST("/statusDict", a.roleService.StatusDict) // 角色状态字典

	}

	userRouter := group.Group("/system/user")
	{
		userRouter.POST("/create", a.userService.Create)
	}

	categoryRouter := group.Group("/system/category")
	{
		categoryRouter.POST("/create", a.categoryService.Create)

	}

	goodRouter := group.Group("/system/good")
	{
		goodRouter.POST("/create", a.goodService.Create)
	}

	a.logger.Info("system api init success")
}
