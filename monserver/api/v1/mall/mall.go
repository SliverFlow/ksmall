package mall

import (
	"github.com/SliverFlow/ksmall/monserver/internal/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Group struct {
	logger          *zap.Logger
	categoryService *service.CategoryService
}

func NewGroup(logger *zap.Logger) *Group {
	return &Group{
		logger: logger,
	}
}

func (a *Group) InitApi(group *gin.RouterGroup) {

	categoryGroup := group.Group("/mall/category")
	{
		categoryGroup.POST("/treeList", a.categoryService.TreeList)
	}

	a.logger.Info("mall api init success")
}
