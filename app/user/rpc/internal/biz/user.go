package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/data/repo"
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"github.com/SliverFlow/ksmall/common/constant"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// UserUsecase 重写 grpc server 方法 并实现 将 grpc server 注册的接口
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

// UserFindByUsername 重写 grpc server 方法
// @Author:  [github.com/SliverFlow]
// @Desc: 根据用户名查询用户
func (uu *UserUsecase) UserFindByUsername(ctx context.Context, req *pb.UserFindByUsernameReq) (*pb.UserInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserFindByUsername not implemented")
}

// UserPageList 重写 grpc server 方法
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
			Username:   user.Username,
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

// UserDelete 重写 grpc server 方法
// @Author: [github.com/SliverFlow]
// @Desc: 删除用户
func (uu *UserUsecase) UserDelete(ctx context.Context, req *pb.IdReq) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserDelete not implemented")
}
