package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/reply"
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
	authorityRepo      repo.AuthorityRepo
}

func NewAuthorityGroupUsecase(
	logger *zap.Logger,
	authorityGroupRepo repo.AuthorityGroupRepo,
	userRepo repo.UserRepo,
	roleRepo repo.RoleRepo,
	authorityRepo repo.AuthorityRepo,
) *AuthorityGroupUsecase {
	return &AuthorityGroupUsecase{
		logger:             logger,
		authorityGroupRepo: authorityGroupRepo,
		userRepo:           userRepo,
		roleRepo:           roleRepo,
		authorityRepo:      authorityRepo,
	}
}

// Create 创建权限组
func (uc *AuthorityGroupUsecase) Create(ctx context.Context, userId int64, param *request.CreateAuthorityGroupReq) error {
	check, err := uc.check(ctx, userId)
	if err != nil || !check {
		return err
	}

	insetAuthorityGroup := &model.AuthorityGroup{
		Name:     param.Name,
		Remark:   param.Remark,
		Sort:     param.Sort,
		Status:   model.Disable,
		UserId:   userId,
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

// FindAllHasAuthority 查询所有权限组
func (uc *AuthorityGroupUsecase) FindAllHasAuthority(ctx context.Context) ([]*reply.AuthGroupHasAuthorityReply, error) {
	authorityGroups, err := uc.authorityGroupRepo.FindAll(ctx)
	if err != nil {
		uc.logger.Error("uc.authorityGroupRepo.FindAll", zap.Error(err))
		return nil, xerror.NewWithMessage("查询权限组失败")
	}

	//// 查询权限组的创始人id
	//var authGroupUserIds []int64
	//for _, authorityGroup := range authorityGroups {
	//	authGroupUserIds = append(authGroupUserIds, authorityGroup.UserId)
	//}
	//
	//// 查询权限组的创始人
	//authGroupUsers, err := uc.userRepo.FindListByUserIds(ctx, authGroupUserIds)
	//if err != nil {
	//	uc.logger.Error("uc.userRepo.FindListByUserIds", zap.Error(err))
	//	return nil, xerror.NewWithMessage("查询权限组创始人失败")
	//}

	var authorityGroupIds []int64
	for _, authorityGroup := range authorityGroups {
		authorityGroupIds = append(authorityGroupIds, authorityGroup.Id)
	}

	authorities, err := uc.authorityRepo.FindByAuthorityGroupIds(ctx, authorityGroupIds)
	if err != nil {
		uc.logger.Error("uc.authorityRepo.FindByAuthorityGroupIds", zap.Error(err))
		return nil, xerror.NewWithMessage("查询权限失败")
	}

	//// 查询权限的创始人id
	//var authUserIds []int64
	//for _, authority := range authorities {
	//	authUserIds = append(authUserIds, authority.UserId)
	//}

	//// 查询权限的创始人
	//authUsers, err := uc.userRepo.FindListByUserIds(ctx, authUserIds)
	//if err != nil {
	//	uc.logger.Error("uc.userRepo.FindListByUserIds", zap.Error(err))
	//	return nil, xerror.NewWithMessage("查询权限创始人失败")
	//}

	// 组装数据
	var authGroupHasAuthorityReplies []*reply.AuthGroupHasAuthorityReply
	for _, authorityGroup := range authorityGroups {
		authGroupHasAuthorityReply := &reply.AuthGroupHasAuthorityReply{
			Id:            authorityGroup.Id,
			Name:          authorityGroup.Name,
			Remark:        authorityGroup.Remark,
			Sort:          authorityGroup.Sort,
			Status:        authorityGroup.Status,
			AuthorityList: make([]*reply.AuthorityReply, 0),
		}

		for _, authority := range authorities {
			if authority.AuthorityGroupId == authorityGroup.Id {
				authGroupHasAuthorityReply.AuthorityList = append(authGroupHasAuthorityReply.AuthorityList, &reply.AuthorityReply{
					Id:     authority.Id,
					Name:   authority.Name,
					Url:    authority.Url,
					Remark: authority.Remark,
					Sort:   authority.Sort,
					Status: authority.Status,
				})
			}
		}

		authGroupHasAuthorityReplies = append(authGroupHasAuthorityReplies, authGroupHasAuthorityReply)
	}

	return authGroupHasAuthorityReplies, nil
}

// Delete 删除
func (uc *AuthorityGroupUsecase) Delete(ctx context.Context, id int64) error {
	authorityGroup, err := uc.authorityGroupRepo.Find(ctx, id)
	if err != nil {
		uc.logger.Error("uc.authorityGroupRepo.FindAll", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("权限组不存在")
		}
		return xerror.NewWithMessage("查询权限组失败")
	}

	authorities, err := uc.authorityRepo.FindByAuthorityGroupIds(ctx, []int64{id})
	if err != nil {
		uc.logger.Error("uc.authorityRepo.FindByAuthorityGroupIds", zap.Error(err))
		return xerror.NewWithMessage("查询权限失败")
	}

	if len(authorities) > 0 {
		return xerror.NewWithMessage("当前组下含有未删除的权限")
	}

	err = uc.authorityGroupRepo.Delete(ctx, authorityGroup.Id)
	if err != nil {
		return xerror.NewWithMessage("权限组删除失败")
	}

	return nil
}
