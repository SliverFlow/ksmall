package system

import (
	"github.com/SliverFlow/ksmall/monserver/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Group struct {
	logger      *zap.Logger
	roleService *service.RoleService
}

func NewGroup(
	logger *zap.Logger,
	roleService *service.RoleService,
) *Group {
	return &Group{
		logger:      logger,
		roleService: roleService,
	}
}

func (a *Group) InitApi(group *gin.RouterGroup) {
	roleRouter := group.Group("/role")
	{
		roleRouter.POST("/create", a.roleService.Create)
		roleRouter.POST("/delete", a.roleService.Delete)
		roleRouter.POST("/update", a.roleService.Update)
		roleRouter.POST("/find", a.roleService.Find)
		roleRouter.POST("/list", a.roleService.List)
		roleRouter.POST("/dict", a.roleService.Dict)             // 角色字典
		roleRouter.POST("/statusDict", a.roleService.StatusDict) // 角色状态字典

		a.logger.Info("role api init success")
	}
	a.logger.Info("system api init success")
}
