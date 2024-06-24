package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/reply"
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
	locker       *EtcdLocker
}

func NewCategoryUsecase(logger *zap.Logger, userRepo repo.UserRepo, categoryRepo repo.CategoryRepo, locker *EtcdLocker) *CategoryUsecase {
	return &CategoryUsecase{
		logger:       logger,
		userRepo:     userRepo,
		categoryRepo: categoryRepo,
		locker:       locker,
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

// TreeList 获取分类树
func (uc *CategoryUsecase) TreeList(ctx context.Context) ([]*reply.CategoryTreeListReply, error) {
	categories, err := uc.categoryRepo.FindAll(ctx)
	if err != nil {
		uc.logger.Error("categoryRepo.FindAll failed", zap.Error(err))
		return nil, xerror.NewWithMessage("获取分类列表失败")
	}

	// 构建分类树
	list := uc.buildChildrenList(ctx, 0, categories)
	return list, nil
}

// buildChildrenList 构建子分类列表
func (uc *CategoryUsecase) buildChildrenList(ctx context.Context, parentId int64, list []*model.Category) []*reply.CategoryTreeListReply {
	var childrenList []*reply.CategoryTreeListReply
	for _, category := range list {
		if category.ParentId == parentId {
			childrenList = append(childrenList, &reply.CategoryTreeListReply{
				Id:       category.Id,
				Name:     category.Name,
				Icon:     category.Icon,
				IsIndex:  category.IsIndex,
				Level:    category.Level,
				Sort:     category.Sort,
				Children: uc.buildChildrenList(ctx, category.Id, list),
			})
		}
	}
	return childrenList
}

// Delete 删除分类
func (uc *CategoryUsecase) Delete(ctx context.Context, categoryId int64) error {
	// 检查分类是否存在
	exist, err := uc.checkExist(ctx, categoryId)
	if err != nil && !exist {
		uc.logger.Error("checkExist failed", zap.Error(err))
		return xerror.NewWithMessage("分类不存在")
	}

	// 删除分类
	err = uc.categoryRepo.Delete(ctx, categoryId)
	if err != nil {
		uc.logger.Error("categoryRepo.Delete failed", zap.Error(err))
		return xerror.NewWithMessage("删除分类失败")
	}

	return nil
}

// Find 获取分类
func (uc *CategoryUsecase) Find(ctx context.Context, categoryId int64) (*model.Category, error) {
	category, err := uc.categoryRepo.Find(ctx, categoryId)
	if err != nil {
		uc.logger.Error("categoryRepo.Find failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.NewWithMessage("分类不存在")
		}
		return nil, xerror.NewWithMessage("获取分类失败")
	}
	return category, nil
}

// Update 更新分类
func (uc *CategoryUsecase) Update(ctx context.Context, param *request.UpdateCategoryReq) error {

	category, err := uc.categoryRepo.Find(ctx, param.Id)
	if err != nil {
		uc.logger.Error("categoryRepo.Find failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("当前分类不存在")
		}
		return xerror.NewWithMessage("查询分类失败")
	}

	parentCategory, err := uc.categoryRepo.Find(ctx, *param.ParenId)
	if err != nil {
		uc.logger.Error("categoryRepo.Find failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("所选父分类分类不存在")
		}
		return xerror.NewWithMessage("查询分类失败")
	}
	if parentCategory.Status == model.Disable {
		return xerror.NewWithMessage("所选父分类禁用中")
	}

	updateCategory := &model.Category{
		Id:       category.Id,
		Name:     param.Name,
		ParentId: *param.ParenId,
		Level:    parentCategory.Level + 1,
		Icon:     param.Icon,
		IsIndex:  *param.IsIndex,
		Sort:     param.Sort,
		Status:   *param.Status,
		UpdateAt: model.Now(),
	}

	err = uc.categoryRepo.Update(ctx, updateCategory)
	if err != nil {
		return xerror.NewWithMessage("分类更新失败")
	}

	return nil
}
