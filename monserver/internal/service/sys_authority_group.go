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

type AuthorityGroupService struct {
	tracing.Service
	logger  *zap.Logger
	usecase *biz.AuthorityGroupUsecase
}

func NewAuthorityGroupService(usecase *biz.AuthorityGroupUsecase, logger *zap.Logger) *AuthorityGroupService {
	return &AuthorityGroupService{
		usecase: usecase,
		logger:  logger,
	}
}

// Create 创建权限组
func (s *AuthorityGroupService) Create(c *gin.Context) {

	var req request.CreateAuthorityGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Create(c, 2, &req)
	if err != nil {
		s.logger.Error("authorityGroupService.Create", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Update 更新权限组
func (s *AuthorityGroupService) Update(c *gin.Context) {

	var req request.UpdateAuthorityGroupReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Update(c, 2, &req)
	if err != nil {
		s.logger.Error("authorityGroupService.Update", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Delete 删除权限组
func (s *AuthorityGroupService) Delete(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Delete(c, req.Id)
	if err != nil {
		s.logger.Error("authorityGroupService.Update", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Find 查询权限组
func (s *AuthorityGroupService) Find(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	reply, err := s.usecase.Find(c, req.Id)
	if err != nil {
		s.logger.Error("authorityGroupService.Find", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(reply, c)
}
