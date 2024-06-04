package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type CategoryUsecase struct {
	logger       *zap.Logger
	userRepo     repo.UserRepo
	categoryRepo repo.CategoryRepo
}

func NewCategoryUsecase(logger *zap.Logger, userRepo repo.UserRepo, categoryRepo repo.CategoryRepo) *CategoryUsecase {
	return &CategoryUsecase{
		logger:       logger,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
	}
}

// Insert 插入分类
func (uc *CategoryUsecase) Insert(ctx context.Context, userId int64, param *request.CreateCategoryReq) error {
	// 检查用户是否存在
	_, err := uc.userRepo.Find(ctx, userId)
	if err != nil {
		uc.logger.Error("userRepo.Find failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("用户不存在")
		}
		return err
	}

	var parentCategory *model.Category
	var level int64 = 1
	if *param.ParenId != 0 {
		// 检查父分类是否存在
		parentCategory, err = uc.categoryRepo.Find(ctx, *param.ParenId)
		if err != nil {
			uc.logger.Error("categoryRepo.Find failed", zap.Error(err))
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return xerror.NewWithMessage("父分类不存在")
			}
			return err
		}
		level = parentCategory.Level + 1
	}

	// 插入分类
	insertCategory := &model.Category{
		Name:     param.Name,
		ParentId: *param.ParenId,
		Icon:     param.Icon,
		IsIndex:  *param.IsIndex,
		Level:    level,
		Sort:     param.Sort,
		Status:   model.Disable,
		Deleted:  model.NotDeleted,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
	}

	err = uc.categoryRepo.Insert(ctx, insertCategory)
	if err != nil {
		uc.logger.Error("categoryRepo.Insert failed", zap.Error(err))
		return err
	}

	return nil
}

// checkExist 检查分类是否存在
func (uc *CategoryUsecase) checkExist(ctx context.Context, categoryId int64) (bool, error) {
	_, err := uc.categoryRepo.Find(ctx, categoryId)
	if err != nil {
		return false, err
	}
	return true, nil
}
