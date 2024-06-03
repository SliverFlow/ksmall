package service

import (
	"github.com/SliverFlow/ksmall/monserver/internal/biz"
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

}

// TreeList 获取分类树
func (a *CategoryService) TreeList(c *gin.Context) {

}
