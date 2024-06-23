package service

import (
	"github.com/SliverFlow/ksmall/monserver/common/response"
	"github.com/SliverFlow/ksmall/monserver/common/util"
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CategoryService struct {
	logger  *zap.Logger
	usecase *biz.CategoryUsecase
}

func NewCategoryService(logger *zap.Logger, usecase *biz.CategoryUsecase) *CategoryService {
	return &CategoryService{
		logger:  logger,
		usecase: usecase,
	}
}

// Create 创建分类
func (a *CategoryService) Create(c *gin.Context) {
	var req request.CreateCategoryReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := a.usecase.Insert(c, 2, &req)
	if err != nil {
		a.logger.Error("categoryService.Insert", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// TreeList 获取分类树
func (a *CategoryService) TreeList(c *gin.Context) {
	list, err := a.usecase.TreeList(c)
	if err != nil {
		a.logger.Error("categoryService.TreeList", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(list, c)
}

// Delete 删除分类
func (a *CategoryService) Delete(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	err := a.usecase.Delete(c, req.Id)
	if err != nil {
		a.logger.Error("categoryService.Delete", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.Ok(c)
}

// Find 获取分类
func (a *CategoryService) Find(c *gin.Context) {
	var req request.IdReq
	if err := c.ShouldBindJSON(&req); err != nil {
		a.logger.Error("param bind err", zap.Error(err))
		response.FailWithMessage(util.ValidaMsg(err, &req), c)
		return
	}

	category, err := a.usecase.Find(c, req.Id)
	if err != nil {
		a.logger.Error("categoryService.Find", zap.Error(err))
		response.FailWithError(err, c)
		return
	}

	response.OkWithData(category, c)
}
