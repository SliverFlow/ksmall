package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"go.uber.org/zap"
)

type UserUsecase struct {
	logger   *zap.Logger
	userRepo repo.UserRepo
	roleRepo repo.RoleRepo
}

func NewUserUsecase(logger *zap.Logger, userRepo repo.UserRepo, roleRepo repo.RoleRepo) *UserUsecase {
	return &UserUsecase{
		logger:   logger,
		userRepo: userRepo,
		roleRepo: roleRepo,
	}
}

// Insert 插入用户
func (u *UserUsecase) Insert(ctx context.Context, param *request.UserCreateReq) error {
	return nil
}
