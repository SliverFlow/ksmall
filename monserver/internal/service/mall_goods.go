package service

import (
	"github.com/SliverFlow/ksmall/monserver/common/response"
	"github.com/SliverFlow/ksmall/monserver/common/util"
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type GoodService struct {
	logger  *zap.Logger
	usecase *biz.GoodUsecase
}

func NewGoodService(logger *zap.Logger, usecase *biz.GoodUsecase) *GoodService {
	return &GoodService{
		logger:  logger,
		usecase: usecase,
	}
}

// Create 创建商品
func (a *GoodService) Create(c *gin.Context) {
	var req request.CreateGoodReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := a.usecase.Insert(c, 2, &req)
	if err != nil {
		a.logger.Error("goodService.Insert", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}
