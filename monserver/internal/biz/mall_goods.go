package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"go.uber.org/zap"
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
	return nil
}
