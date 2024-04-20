package biz

import (
	"context"
	"errors"
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/data/repo"
	"github.com/SliverFlow/ksmall/server/app/user/rpc/pb"
	"github.com/SliverFlow/ksmall/server/common/constant"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"gorm.io/gorm"
)

// UserUsecase  并实现 将 grpc server 注册的接口
type UserUsecase struct {
	iUserRepo repo.IUserRepo
	pb.UnimplementedUserServer

	logger *zap.Logger
}

func NewUserUsecase(
	iUserRepo repo.IUserRepo,
	logger *zap.Logger,
) *UserUsecase {
	return &UserUsecase{
		iUserRepo: iUserRepo,
		logger:    logger,
	}
}

// RegisterServer 注册 grpc server
func (uu *UserUsecase) RegisterServer(server *grpc.Server) *grpc.Server {
	pb.RegisterUserServer(server, uu)
	return server
}

// UserPageList
// @Author:  [github.com/SliverFlow]
// @Desc: 用户分页查询
func (uu *UserUsecase) UserPageList(ctx context.Context, req *pb.UserPageListReq) (*pb.UserPageListReply, error) {
	users, total, err := uu.iUserRepo.PageList(ctx, req.Limit, req.Offset)
	if err != nil {
		return nil, status.Errorf(constant.UserPageListErrorCode, "get user list error: %v", err)
	}
	var userInfos []*pb.UserInfo
	for _, user := range users {
		userInfos = append(userInfos, &pb.UserInfo{
			Id:         user.Id,
			WxOpenId:   user.WxOpenId,
			Nickname:   user.Nickname,
			Email:      user.Email,
			Phone:      user.Phone,
			RoleId:     user.RoleId,
			Avatar:     user.Avatar,
			Status:     user.Status,
			CreateTime: user.CreateTime,
			UpdateTime: user.UpdateTime,
		})
	}
	return &pb.UserPageListReply{
		Total: total,
		List:  userInfos,
	}, nil
}

// UserDelete
// @Author: [github.com/SliverFlow]
// @Desc: 删除用户
func (uu *UserUsecase) UserDelete(ctx context.Context, req *pb.IdReq) (*emptypb.Empty, error) {
	uu.logger.Info("[rpc] delete user", zap.Int64("id", req.Id))

	_, err := uu.iUserRepo.Find(ctx, req.Id)
	if err != nil {
		uu.logger.Error("[rpc] find user error", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(constant.GormNotFoundCode, "user not found")
		}
		return nil, err
	}

	err = uu.iUserRepo.Delete(ctx, req.Id)
	if err != nil {
		uu.logger.Error("[rpc] delete user error", zap.Error(err))
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

// UserFind
// @Author: [github.com/SliverFlow]
// @Desc: 查询用户
func (uu *UserUsecase) UserFind(ctx context.Context, req *pb.IdReq) (*pb.UserInfo, error) {
	uu.logger.Info("[rpc] find user", zap.Int64("id", req.Id))

	user, err := uu.iUserRepo.Find(ctx, req.Id)
	if err != nil {
		uu.logger.Error("[rpc] find user error", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(constant.GormNotFoundCode, "user not found")
		}
		return nil, err
	}

	return &pb.UserInfo{
		Id:         user.Id,
		WxOpenId:   user.WxOpenId,
		Nickname:   user.Nickname,
		Email:      user.Email,
		Phone:      user.Phone,
		RoleId:     user.RoleId,
		Avatar:     user.Avatar,
		Status:     user.Status,
		Male:       user.Male,
		Birthday:   user.Birthday,
		VIPLevel:   user.VIPLevel,
		Points:     user.Points,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}, nil
}

// UserFindByEmailReq
// @Author: [github.com/SliverFlow]
// @Desc: 根据邮箱查询用户
func (uu *UserUsecase) UserFindByEmailReq(ctx context.Context, req *pb.UserFindByEmailReq) (*pb.UserInfo, error) {
	uu.logger.Info("[rpc] find user by email", zap.String("email", req.Email))

	user, err := uu.iUserRepo.FindByEmail(ctx, req.Email)
	if err != nil {
		uu.logger.Error("[rpc] find user by email error", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(constant.GormNotFoundCode, "user not found")
		}
		return nil, err
	}

	return &pb.UserInfo{
		Id:         user.Id,
		WxOpenId:   user.WxOpenId,
		Nickname:   user.Nickname,
		Email:      user.Email,
		Phone:      user.Phone,
		RoleId:     user.RoleId,
		Avatar:     user.Avatar,
		Status:     user.Status,
		Male:       user.Male,
		Birthday:   user.Birthday,
		VIPLevel:   user.VIPLevel,
		Points:     user.Points,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}, nil
}

// UserFindByPhoneReq
// @Author: [github.com/SliverFlow]
// @Desc: 根据手机号查询用户
func (uu *UserUsecase) UserFindByPhoneReq(ctx context.Context, req *pb.UserFindByPhoneReq) (*pb.UserInfo, error) {
	uu.logger.Info("[rpc] find user by phone", zap.String("phone", req.Phone))

	user, err := uu.iUserRepo.FindByPhone(ctx, req.Phone)
	if err != nil {
		uu.logger.Error("[rpc] find user by phone error", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Errorf(constant.GormNotFoundCode, "user not found")
		}
		return nil, err
	}

	return &pb.UserInfo{
		Id:         user.Id,
		WxOpenId:   user.WxOpenId,
		Nickname:   user.Nickname,
		Email:      user.Email,
		Phone:      user.Phone,
		RoleId:     user.RoleId,
		Avatar:     user.Avatar,
		Status:     user.Status,
		Male:       user.Male,
		Birthday:   user.Birthday,
		VIPLevel:   user.VIPLevel,
		Points:     user.Points,
		CreateTime: user.CreateTime,
		UpdateTime: user.UpdateTime,
	}, nil
}
