package v1

import (
	"github.com/SliverFlow/core/middleware"
	"github.com/SliverFlow/ksmall/monserver/api/v1/mall"
	"github.com/SliverFlow/ksmall/monserver/api/v1/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Group struct {
	logger *zap.Logger

	SystemApi *system.Group
	MallApi   *mall.Group
	tacker    *middleware.Tacker
}

func NewGroup(logger *zap.Logger, systemApi *system.Group, mallApi *mall.Group, tacker *middleware.Tacker) *Group {
	return &Group{
		logger:    logger,
		SystemApi: systemApi,
		MallApi:   mallApi,
		tacker:    tacker,
	}
}

func (v *Group) InitApi(r *gin.Engine) {
	// r.Use(v.tacker.Handle())
	group := r.Group("/api/v1")
	v.logger.Info("api v1 init start")

	{
		v.SystemApi.InitApi(group)
		v.MallApi.InitApi(group)
	}
}
