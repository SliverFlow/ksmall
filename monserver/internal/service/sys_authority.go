package service

import (
	"github.com/SliverFlow/core/tracing"
	"github.com/SliverFlow/ksmall/monserver/common/response"
	"github.com/SliverFlow/ksmall/monserver/common/util"
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type AuthorityService struct {
	tracing.Service
	logger  *zap.Logger
	usecase *biz.AuthorityUsecase
}

func NewAuthorityService(usecase *biz.AuthorityUsecase, logger *zap.Logger) *AuthorityService {
	return &AuthorityService{
		usecase: usecase,
		logger:  logger,
	}
}

// Create 创建一个权限
func (a *AuthorityService) Create(c *gin.Context) {
	var req request.CreateAuthorityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := a.usecase.Create(c, 2, &req)
	if err != nil {
		a.logger.Error("authorityService.Insert", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Delete 删除权限
func (a *AuthorityService) Delete(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := a.usecase.Delete(c, req.Id)
	if err != nil {
		a.logger.Error("authorityService.Insert", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Update 更新权限
func (a *AuthorityService) Update(c *gin.Context) {
	var req request.UpdateAuthorityReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := a.usecase.Update(c, &req)
	if err != nil {
		a.logger.Error("authorityService.update", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Find 查询权限
func (a *AuthorityService) Find(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	reply, err := a.usecase.Find(c, req.Id)
	if err != nil {
		a.logger.Error("authorityService.Insert", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(reply, c)
}
