package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/SliverFlow/ksmall/monserver/tracing"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AuthorityUsecase struct {
	tracing.Biz
	logger             *zap.Logger
	authorityRepo      repo.AuthorityRepo
	authorityGroupRepo repo.AuthorityGroupRepo
	userRepo           repo.UserRepo
	roleRepo           repo.RoleRepo
}

func NewAuthorityUsecase(
	authorityRepo repo.AuthorityRepo,
	logger *zap.Logger,
) *AuthorityUsecase {
	return &AuthorityUsecase{
		authorityRepo: authorityRepo,
		logger:        logger,
	}
}

// check
func (uc *AuthorityUsecase) check(ctx context.Context, userId int64) (bool, error) {
	findUser, err := uc.userRepo.Find(ctx, userId)
	if err != nil {
		uc.logger.Error("uc.userRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, xerror.NewWithMessage("用户不存在")
		}
		return false, err
	}
	if findUser.Status == model.Disable {
		return false, xerror.NewWithMessage("用户已禁用")
	}

	roleId, err := uc.userRepo.FindRoleId(ctx, userId)
	if err != nil {
		uc.logger.Error("uc.userRepo.FindRoleId", zap.Error(err))
		return false, err
	}
	if roleId == 0 {
		return false, xerror.NewWithMessage("用户未分配角色")
	}

	findRole, err := uc.roleRepo.Find(ctx, roleId)
	if err != nil {
		uc.logger.Error("uc.roleRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, xerror.NewWithMessage("角色不存在")
		}
		return false, err
	}
	if findRole.Status == model.Disable {
		return false, xerror.NewWithMessage("角色已禁用")
	}
	if findRole.Key != model.RoleKeyAdmin {
		return false, xerror.NewWithMessage("管理员角色不允许创建权限组")
	}
	return true, nil
}

// Create 创建权限
func (uc *AuthorityUsecase) Create(ctx context.Context, userId int64, param *request.CreateAuthorityReq) error {
	ok, err := uc.check(ctx, userId)
	if err != nil || !ok {
		return err
	}

	findAuthorityGroup, err := uc.authorityGroupRepo.Find(ctx, param.AuthorityGroupId)
	if err != nil {
		uc.logger.Error("uc.authorityGroupRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("权限组不存在")
		}
		return err
	}
	if findAuthorityGroup.Status == model.Disable {
		return xerror.NewWithMessage("权限组已禁用")
	}

	authority := &model.Authority{
		AuthorityGroupId: param.AuthorityGroupId,
		Name:             param.Name,
		Remark:           param.Remark,
		Status:           model.Enable,
		Auth:             *param.Auth,
		UserId:           userId,
		Url:              param.Url,
		Sort:             param.Sort,
		CreateAt:         model.Now(),
		UpdateAt:         model.Now(),
		Deleted:          model.NotDeleted,
	}
	_, err = uc.authorityRepo.Insert(ctx, authority)
	if err != nil {
		uc.logger.Error("uc.authorityRepo.Create", zap.Error(err))
		return err
	}
	return nil
}
