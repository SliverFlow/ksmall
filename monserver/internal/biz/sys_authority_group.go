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
	"time"
)

type AuthorityGroupUsecase struct {
	tracing.Biz
	logger             *zap.Logger
	userRepo           repo.UserRepo
	roleRepo           repo.RoleRepo
	authorityGroupRepo repo.AuthorityGroupRepo
}

func NewAuthorityGroupUsecase(
	authorityGroupRepo repo.AuthorityGroupRepo,
	userRepo repo.UserRepo,
	roleRepo repo.RoleRepo,
) *AuthorityGroupUsecase {
	return &AuthorityGroupUsecase{
		authorityGroupRepo: authorityGroupRepo,
		userRepo:           userRepo,
		roleRepo:           roleRepo,
	}
}

// Create 创建权限组
func (uc *AuthorityGroupUsecase) Create(ctx context.Context, userId int64, param *request.CreateAuthorityGroupReq) error {
	findUser, err := uc.userRepo.Find(ctx, userId)
	if err != nil {
		uc.logger.Error("uc.userRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("用户不存在")
		}
		return err
	}
	if findUser.Status == model.Disable {
		return xerror.NewWithMessage("用户已禁用")
	}

	roleId, err := uc.userRepo.FindRoleId(ctx, userId)
	if err != nil {
		uc.logger.Error("uc.userRepo.FindRoleId", zap.Error(err))
		return err
	}
	if roleId == 0 {
		return xerror.NewWithMessage("用户未分配角色")
	}

	findRole, err := uc.roleRepo.Find(ctx, roleId)
	if err != nil {
		uc.logger.Error("uc.roleRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("角色不存在")
		}
		return err
	}
	if findRole.Status == model.Disable {
		return xerror.NewWithMessage("角色已禁用")
	}
	if findRole.Key != model.RoleKeyAdmin {
		return xerror.NewWithMessage("管理员角色不允许创建权限组")
	}

	insetAuthorityGroup := &model.AuthorityGroup{
		Name:     param.Name,
		Remark:   param.Remark,
		Sort:     param.Sort,
		Status:   model.Disable,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
		Deleted:  model.NotDeleted,
	}

	_, err = uc.authorityGroupRepo.Insert(ctx, insetAuthorityGroup)
	if err != nil {
		uc.logger.Error("uc.authorityGroupRepo.Insert", zap.Error(err))
		return err
	}
	return nil
}

// check
func (uc *AuthorityGroupUsecase) check(ctx context.Context, userId int64) (bool, error) {
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

// Update 更新权限组
func (uc *AuthorityGroupUsecase) Update(ctx context.Context, userId int64, param *request.UpdateAuthorityGroupReq) error {
	check, err := uc.check(ctx, userId)
	if err != nil || !check {
		return err
	}

	_, err = uc.authorityGroupRepo.Find(ctx, param.Id)
	if err != nil {
		uc.logger.Error("uc.authorityGroupRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("权限组不存在")
		}
		return err
	}

	updateAuthorityGroup := &model.AuthorityGroup{
		Id:       param.Id,
		Name:     param.Name,
		Remark:   param.Remark,
		Sort:     param.Sort,
		Status:   *param.Status,
		UpdateAt: time.Now().Unix(),
	}

	_, err = uc.authorityGroupRepo.Update(ctx, updateAuthorityGroup)
	if err != nil {
		uc.logger.Error("uc.authorityGroupRepo.Update", zap.Error(err))
		return err
	}

	return nil
}
