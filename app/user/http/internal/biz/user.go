package biz

import (
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"go.uber.org/zap"
)

type UserUsecase struct {
	userRpc pb.UserClient

	logger *zap.Logger
}

func NewUserUsecase(logger *zap.Logger, userRpc pb.UserClient) *UserUsecase {
	return &UserUsecase{
		logger:  logger,
		userRpc: userRpc,
	}
}
