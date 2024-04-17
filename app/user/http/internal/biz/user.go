package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/http/internal/model/reply"
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"go.uber.org/zap"
)

type UserUsecase struct {
	userRpc pb.UserClient
	logger  *zap.Logger
}

func NewUserUsecase(logger *zap.Logger, userRpc pb.UserClient) *UserUsecase {
	return &UserUsecase{
		logger:  logger,
		userRpc: userRpc,
	}
}

func (uu *UserUsecase) FindById(ctx context.Context, id uint) (*reply.UserInfoReply, error) {
	_, err := uu.userRpc.UserFindByUsername(ctx, &pb.UserFindByUsernameReq{Username: "id"})
	if err != nil {
		return nil, err
	}

	return &reply.UserInfoReply{}, nil
}
