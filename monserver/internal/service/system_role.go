package service

import (
	"github.com/SliverFlow/ksmall/monserver/common/response"
	"github.com/SliverFlow/ksmall/monserver/common/util"
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RoleService struct {
	logger  *zap.Logger
	usecase *biz.RoleUsecase
}

func NewRoleService(logger *zap.Logger, usecase *biz.RoleUsecase) *RoleService {
	return &RoleService{
		logger:  logger,
		usecase: usecase,
	}
}

func (s *RoleService) Create(c *gin.Context) {
	var req request.RoleCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Insert(c, &req)
	if err != nil {
		s.logger.Error("roleService.Insert", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

func (s *RoleService) Delete(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Delete(c, &req)
	if err != nil {
		s.logger.Error("roleService.Delete", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

func (s *RoleService) Update(c *gin.Context) {
	var req request.RoleUpdateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Update(c, &req)
	if err != nil {
		s.logger.Error("roleService.Update", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

func (s *RoleService) Find(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	role, err := s.usecase.Find(c, req.Id)
	if err != nil {
		s.logger.Error("roleService.Find", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(gin.H{"role": role}, c)
}

func (s *RoleService) List(c *gin.Context) {
	var req request.RoleListReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	roles, err := s.usecase.List(c, &req)
	if err != nil {
		s.logger.Error("roleService.List", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(gin.H{"list": roles}, c)
}

// Dict 查询角色字典
func (s *RoleService) Dict(c *gin.Context) {
	dict, err := s.usecase.Dict(c)
	if err != nil {
		s.logger.Error("roleService.Dict", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(gin.H{"dict": dict}, c)
}

// StatusDict 查询角色状态字典
func (s *RoleService) StatusDict(c *gin.Context) {
	dict := s.usecase.StatusDict(c)
	response.OkWithData(gin.H{"dict": dict}, c)
}
