package mall

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Group struct {
	logger *zap.Logger
}

func NewGroup(logger *zap.Logger) *Group {
	return &Group{
		logger: logger,
	}
}

func (a *Group) InitApi(group *gin.RouterGroup) {

	a.logger.Info("mall api init success")
}
