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

// Find 根据id查询用户
func (uu *UserUsecase) Find(ctx context.Context, id int64) (*reply.UserInfoReply, error) {
	userInfo, err := uu.userRpc.UserFind(ctx, &pb.IdReq{Id: id})
	if err != nil {
		uu.logger.Error("[UserUsecase] failed to find user", zap.Error(err))
		return nil, err
	}
	return &reply.UserInfoReply{
		Id:         userInfo.Id,
		Nickname:   userInfo.Nickname,
		Email:      userInfo.Email,
		Phone:      userInfo.Phone,
		Avatar:     userInfo.Avatar,
		RoleId:     userInfo.RoleId,
		Male:       userInfo.Male,
		Birthday:   userInfo.Birthday,
		VIPLevel:   userInfo.VIPLevel,
		Points:     userInfo.Points,
		CreateTime: userInfo.CreateTime,
	}, nil
}
