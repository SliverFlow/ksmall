package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/data/repo"
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"go.uber.org/zap"
	"google.golang.org/grpc"
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

func (uu *UserUsecase) RegisterServer(server *grpc.Server) *grpc.Server {
	pb.RegisterUserServer(server, uu)
	return server
}

func (uu *UserUsecase) FindByUsername(ctx context.Context, req *pb.FindByUsernameReq) (*pb.UserInfo, error) {
	uu.logger.Info("user rpc", zap.Any("request", req))
	return &pb.UserInfo{}, nil
}
func (uu *UserUsecase) Delete(ctx context.Context, req *pb.IdReq) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
