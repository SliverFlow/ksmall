package biz

import (
	"context"
	"github.com/SliverFlow/core/tracing"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/reply"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
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
	authorityGroupRepo repo.AuthorityGroupRepo,
	userRepo repo.UserRepo,
	roleRepo repo.RoleRepo,
) *AuthorityUsecase {
	return &AuthorityUsecase{
		authorityRepo:      authorityRepo,
		logger:             logger,
		authorityGroupRepo: authorityGroupRepo,
		userRepo:           userRepo,
		roleRepo:           roleRepo,
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

// Delete 删除权限
func (uc *AuthorityUsecase) Delete(ctx context.Context, id int64) error {
	authority, err := uc.authorityRepo.Find(ctx, id)
	if err != nil {
		uc.logger.Error("uc.authorityRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("此权限不存在")
		}
		return xerror.NewWithMessage("权限查询失败")
	}

	err = uc.authorityRepo.Delete(ctx, authority.Id)
	if err != nil {
		return xerror.NewWithMessage("权限删除失败")
	}

	return nil
}

// Find 查询权限
func (uc *AuthorityUsecase) Find(ctx context.Context, id int64) (*reply.AuthorityFindReply, error) {
	authority, err := uc.authorityRepo.Find(ctx, id)
	if err != nil {
		uc.logger.Error("uc.authorityRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.NewWithMessage("此权限不存在")
		}
		return nil, xerror.NewWithMessage("权限查询失败")
	}

	resp := &reply.AuthorityFindReply{
		Id:       authority.Id,
		Name:     authority.Name,
		Remark:   authority.Remark,
		Url:      authority.Url,
		Auth:     authority.Auth,
		Status:   authority.Status,
		Sort:     authority.Sort,
		Username: "未知用户",
		RoleName: "未知角色",
		CreateAt: authority.CreateAt,
	}

	user, err := uc.userRepo.Find(ctx, authority.UserId)
	if err != nil {
		return resp, nil
	}
	roleId, err := uc.userRepo.FindRoleId(ctx, user.Id)
	if err != nil {
		return resp, nil
	}
	role, err := uc.roleRepo.Find(ctx, roleId)
	if err != nil {
		return resp, nil
	}

	resp.Username = user.Nickname
	resp.RoleName = role.Name

	return resp, nil
}

// Update 更新权限
func (uc *AuthorityUsecase) Update(ctx context.Context, param *request.UpdateAuthorityReq) error {
	authorityGroup, err := uc.authorityGroupRepo.Find(ctx, param.AuthorityGroupId)
	if err != nil {
		uc.logger.Error("uc.authorityGroupRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("权限组不存在")
		}
		return xerror.NewWithMessage("权限组查询失败")
	}
	if authorityGroup.Status != model.Enable {
		return xerror.NewWithMessage("权限组状态未开启")
	}

	authority, err := uc.authorityRepo.Find(ctx, param.Id)
	if err != nil {
		uc.logger.Error("uc.authorityRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("权限不存在")
		}
		return xerror.NewWithMessage("权限查询失败")
	}
	updateAuthority := model.Authority{
		Id:               authority.Id,
		UserId:           authority.UserId,
		AuthorityGroupId: param.AuthorityGroupId,
		Name:             param.Name,
		Url:              param.Url,
		Auth:             *param.Auth,
		Remark:           param.Remark,
		Status:           param.Sort,
		Sort:             param.Sort,
		UpdateAt:         model.Now(),
	}
	_, err = uc.authorityRepo.Update(ctx, &updateAuthority)
	if err != nil {
		uc.logger.Error("uc.authorityRepo.Update", zap.Error(err))
		return xerror.NewWithMessage("权限更新失败")
	}

	return nil
}
