package service

import (
	"github.com/SliverFlow/ksmall/monserver/common/response"
	"github.com/SliverFlow/ksmall/monserver/common/util"
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserService struct {
	logger  *zap.Logger
	usecase *biz.UserUsecase
}

func NewUserService(logger *zap.Logger, usecase *biz.UserUsecase) *UserService {
	return &UserService{
		logger:  logger,
		usecase: usecase,
	}
}

// Create 创建用户
func (s *UserService) Create(c *gin.Context) {
	var req request.UserCreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Insert(c, &req)
	if err != nil {
		s.logger.Error("userService.Insert", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Delete 通过id删除用户
func (s *UserService) Delete(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := s.usecase.Delete(c, req.Id)
	if err != nil {
		s.logger.Error("userService.Delete", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Find 通过id查找用户
func (s *UserService) Find(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	user, err := s.usecase.Find(c, req.Id)
	if err != nil {
		s.logger.Error("userService.Find", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(gin.H{"user": user}, c)
}

// Update 更新用户信息
func (s *UserService) Update(c *gin.Context) {
	var req request.UpdateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		s.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}
}
