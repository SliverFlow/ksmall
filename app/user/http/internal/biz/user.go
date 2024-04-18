package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/http/internal/model/reply"
	"github.com/SliverFlow/ksmall/app/user/http/internal/model/request"
	"github.com/SliverFlow/ksmall/app/user/rpc/pb"
	"github.com/SliverFlow/ksmall/common/constant"
	"github.com/SliverFlow/ksmall/common/util"
	"github.com/SliverFlow/ksmall/common/zerror"
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

const (
	// LoginTypeEmail 邮箱登录类型
	LoginTypeEmail = 1
	// LoginTypePhone 手机登录类型
	LoginTypePhone = 2
)

// Captcha 发送验证码
func (uu *UserUsecase) Captcha(ctx context.Context, account string, captchaType int64) error {
	switch captchaType {
	case LoginTypeEmail:
		return uu.sendEmailCaptcha(ctx, account)
	case LoginTypePhone:
		return uu.sendPhoneCaptcha(ctx, account)
	default:
		return zerror.NewWithCode(constant.CaptchaTypeCode)
	}
}

// sendEmailCaptcha 发送邮箱验证码
func (uu *UserUsecase) sendEmailCaptcha(ctx context.Context, email string) error {
	return nil
}

// sendPhoneCaptcha 发送手机验证码
func (uu *UserUsecase) sendPhoneCaptcha(ctx context.Context, phone string) error {
	return nil
}

// Login 登录
func (uu *UserUsecase) Login(ctx context.Context, req *request.LoginReq) (*reply.UserInfoReply, error) {
	if req.Type != LoginTypeEmail && req.Type != LoginTypePhone {
		return nil, zerror.NewWithCode(constant.CaptchaTypeCode)
	}

	// 验证验证码
	if err := uu.verifyCaptcha(ctx, req.Account, req.Captcha); err != nil {
		return nil, err
	}

	// 查询用户信息
	var pbUserInfo *pb.UserInfo
	var err error
	if req.Type == LoginTypeEmail {
		pbUserInfo, err = uu.userRpc.UserFindByEmail(ctx, &pb.UserFindByEmailReq{Email: req.Account})
	} else {
		pbUserInfo, err = uu.userRpc.UserFindByPhone(ctx, &pb.UserFindByPhoneReq{Phone: req.Account})
	}

	if err != nil && !util.IsGormNotFoundErr(err) && pbUserInfo == nil {
		return nil, err
	}

	// 用户不存在，注册用户
	uuid, err := util.GenerateUUID()
	if err != nil {
		return nil, zerror.NewWithCode(constant.ServerInternalCode)
	}

	pbUserCreate := pb.UserCreateReq{
		Uuid:     uuid,
		Nickname: "昵称_" + uuid[0:5],
	}

	if req.Type == LoginTypeEmail {
		pbUserCreate.Email = req.Account
	} else {
		pbUserCreate.Phone = req.Account
	}

	pbUserId, err := uu.userRpc.UserCreate(ctx, &pbUserCreate)
	if err != nil {
		return nil, err
	}

	// 登录成功
	pbUserInfoFind, err := uu.userRpc.UserFind(ctx, &pb.IdReq{Id: pbUserId.Id})
	if err != nil {
		return nil, err
	}

	return &reply.UserInfoReply{
		Id:         pbUserInfoFind.Id,
		Nickname:   pbUserInfoFind.Nickname,
		Email:      pbUserInfoFind.Email,
		Phone:      pbUserInfoFind.Phone,
		Avatar:     pbUserInfoFind.Avatar,
		RoleId:     pbUserInfoFind.RoleId,
		Birthday:   pbUserInfoFind.Birthday,
		VIPLevel:   pbUserInfoFind.VIPLevel,
		Points:     pbUserInfoFind.Points,
		CreateTime: pbUserInfoFind.CreateTime,
	}, nil
}

// verifyCaptcha 验证验证码
func (uu *UserUsecase) verifyCaptcha(ctx context.Context, account, captcha string) error {
	return nil
}
