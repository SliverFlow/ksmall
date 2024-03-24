package service

import (
	"github.com/SliverFlow/ksmall/app/user/internal/biz"
	"go.uber.org/zap"
)

type UserService struct {
	logger      *zap.Logger
	userUsecase *biz.UserUsecase
}

func NewUserService(logger *zap.Logger, userUsecase *biz.UserUsecase) *UserService {
	return &UserService{logger: logger, userUsecase: userUsecase}
}
