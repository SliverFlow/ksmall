package biz

import (
	"go.uber.org/zap"
)

type UserUsecase struct {
	logger *zap.Logger
}

func NewUserUsecase(logger *zap.Logger) *UserUsecase {
	return &UserUsecase{
		logger: logger,
	}
}
