package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/common/util"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
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
	roleKey, err := strconv.ParseInt(param.RoleId, 10, 64)
	if err != nil {
		u.logger.Error("[strconv err] strconv.ParseInt", zap.Error(err))
		return xerror.NewWithMessage("角色ID格式错误")
	}
	role, err := u.roleRepo.FindByKey(ctx, roleKey)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.FindByKey", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("角色不存在")
		}
		return err
	}

	if role.Status == model.Disable {
		return xerror.NewWithMessage("角色已禁用")
	}

	insertUser := &model.User{
		Username: param.Username,
		Password: util.BcryptHash(param.Password),
		Nickname: param.Nickname,
		Phone:    param.Phone,
		Email:    param.Email,
		Male:     *param.Male,
		Avatar:   param.Avatar,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
		Deleted:  model.NotDeleted,
		Status:   model.Enable,
	}
	_, err = u.userRepo.Insert(ctx, insertUser)
	if err != nil {
		u.logger.Error("[date repo err] userRepo.Insert", zap.Error(err))
		return xerror.NewWithMessage("用户创建失败")
	}
	return nil
}
