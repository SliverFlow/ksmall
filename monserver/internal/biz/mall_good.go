package biz

import (
	"context"
	"errors"
	"github.com/SliverFlow/ksmall/monserver/common/util"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type GoodUsecase struct {
	logger       *zap.Logger
	goodRepo     repo.GoodRepo
	categoryRepo repo.CategoryRepo
	userRepo     repo.UserRepo
	stockRepo    repo.StockRepo
}

func NewGoodUsecase(logger *zap.Logger,
	goodRepo repo.GoodRepo,
	categoryRepo repo.CategoryRepo,
	userRepo repo.UserRepo,
	stockRepo repo.StockRepo,
) *GoodUsecase {
	return &GoodUsecase{
		logger:       logger,
		goodRepo:     goodRepo,
		categoryRepo: categoryRepo,
		userRepo:     userRepo,
		stockRepo:    stockRepo,
	}
}

// Insert a good
func (uc *GoodUsecase) Insert(ctx context.Context, userId int64, param *request.CreateGoodReq) error {
	// Check if the user exists
	_, err := uc.userRepo.Find(ctx, userId)
	if err != nil {
		uc.logger.Error("userRepo.Find failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("用户不存在")
		}
		return err
	}

	// Check if the category exists
	category, err := uc.categoryRepo.Find(ctx, param.CategoryId)
	if err != nil {
		uc.logger.Error("categoryRepo.Find failed", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("分类不存在")
		}
		return err
	}
	if category.Status == model.Disable {
		return xerror.NewWithMessage("当前分类已禁用")
	}

	// Insert a good
	insertGood := &model.Good{
		Name:        param.Name,
		CategoryId:  param.CategoryId,
		Sn:          strconv.FormatInt(util.SnowWorker.NextId(), 10),
		Price:       param.Price,
		OriginPrice: param.OriginPrice,
		Cover:       param.Cover,
		Desc:        param.Desc,
		Status:      model.Disable,
		Deleted:     model.NotDeleted,
		CreateAt:    time.Now().Unix(),
		UpdateAt:    time.Now().Unix(),
	}
	imageStr, err := util.StructToJSON(param.Image)
	contentImageStr, err := util.StructToJSON(param.ContentImage)
	if err != nil {
		return xerror.NewWithMessage("内容图片转换失败")
	}
	insertGood.Image = imageStr
	insertGood.ContentImage = contentImageStr

	insertGood, err = uc.goodRepo.Insert(ctx, insertGood)
	if err != nil {
		uc.logger.Error("goodRepo.Insert failed", zap.Error(err))
		return err
	}

	// 异步给商品创建默认 stock
	go func() {
		insertStock := &model.Stock{
			GoodId:   insertGood.Id,
			Name:     "默认库存",
			Remark:   "默认库存",
			Num:      0,
			Version:  0,
			Status:   model.Disable,
			Deleted:  model.NotDeleted,
			CreateAt: time.Now().Unix(),
			UpdateAt: time.Now().Unix(),
		}
		err := uc.stockRepo.Insert(ctx, insertStock)
		if err != nil {
			uc.logger.Error("stockRepo.Insert failed", zap.Error(err))
		}
	}()

	return nil
}
