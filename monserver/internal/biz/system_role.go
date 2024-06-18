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
	"strconv"
	"time"
)

type RoleUsecase struct {
	tracing.Biz
	logger        *zap.Logger
	roleRepo      repo.RoleRepo
	userRepo      repo.UserRepo
	authorityRepo repo.AuthorityRepo
}

func NewRoleUsecase(
	logger *zap.Logger,
	roleRepo repo.RoleRepo,
	userRepo repo.UserRepo,
	authorityRepo repo.AuthorityRepo,
) *RoleUsecase {
	return &RoleUsecase{
		logger:        logger,
		roleRepo:      roleRepo,
		userRepo:      userRepo,
		authorityRepo: authorityRepo,
	}
}

func (u *RoleUsecase) Insert(ctx context.Context, params *request.RoleCreateReq) error {
	_, err := u.roleRepo.FindByName(ctx, params.Name)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		u.logger.Error("[date repo err] roleRepo.FindByName", zap.Error(err))
		return err
	}

	insertRole := &model.Role{
		Name:     params.Name,
		Remark:   params.Remark,
		Status:   model.Disable,
		Key:      params.Key,
		Sorted:   params.Sorted,
		CreateAt: time.Now().Unix(),
		UpdateAt: time.Now().Unix(),
		Deleted:  model.NotDeleted,
	}
	_, err = u.roleRepo.Insert(ctx, insertRole)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.Insert", zap.Error(err))
		return xerror.NewWithMessage("角色创建失败")
	}

	return nil
}

// Delete 删除角色
func (u *RoleUsecase) Delete(ctx context.Context, param *request.IdReq) error {
	role, err := u.roleRepo.Find(ctx, param.Id)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.Find", zap.Error(err))
		return xerror.NewWithMessage("角色不存在")
	}

	err = u.roleRepo.Delete(ctx, role.Id)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.Delete", zap.Error(err))
		return xerror.NewWithMessage("角色删除失败")
	}

	return nil
}

// Update 更新角色
func (u *RoleUsecase) Update(ctx context.Context, param *request.RoleUpdateReq) error {
	role, err := u.roleRepo.Find(ctx, param.Id)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.Find", zap.Error(err))
		return xerror.NewWithMessage("角色不存在")
	}

	role.Name = param.Name
	role.Remark = param.Remark
	role.Key = param.Key
	role.Sorted = param.Sorted
	role.Status = *param.Status
	role.UpdateAt = time.Now().Unix()

	_, err = u.roleRepo.Update(ctx, role)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.Update", zap.Error(err))
		return xerror.NewWithMessage("角色更新失败")
	}

	return nil
}

// Find 查询角色
func (u *RoleUsecase) Find(ctx context.Context, id int64) (*model.Role, error) {
	role, err := u.roleRepo.Find(ctx, id)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, xerror.NewWithMessage("角色不存在")
		}
		return nil, err
	}

	return role, nil
}

// List 查询角色列表
func (u *RoleUsecase) List(ctx context.Context, param *request.RoleListReq) ([]*model.Role, error) {
	ctx, span := u.Tacker(ctx, "roleUsecase.List")
	defer span.End()

	roles, err := u.roleRepo.ListByCondition(ctx, param.Name, *param.Status)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.List", zap.Error(err))
		return nil, xerror.NewWithMessage("角色列表查询失败")
	}

	return roles, nil
}

// Dict 查询角色字典
func (u *RoleUsecase) Dict(ctx context.Context) ([]*reply.DictReply, error) {
	roles, err := u.roleRepo.ListByCondition(ctx, "", model.Enable)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.List", zap.Error(err))
		return nil, xerror.NewWithMessage("角色字典查询失败")
	}

	var dict []*reply.DictReply
	for _, role := range roles {
		dict = append(dict, &reply.DictReply{
			Id:    role.Id,
			Name:  strconv.FormatInt(role.Key, 10),
			Value: role.Name,
		})
	}

	return dict, nil
}

// StatusDict 查询角色状态字典
func (u *RoleUsecase) StatusDict(ctx context.Context) []*reply.DictReply {
	var dict []*reply.DictReply
	for k, v := range model.StatusMap {
		dict = append(dict, &reply.DictReply{
			Id:    k + 1,
			Name:  v,
			Value: strconv.FormatInt(k, 10),
		})
	}

	return dict
}

// CheckAccess 检查用户是否有权限
func (u *RoleUsecase) CheckAccess(ctx context.Context, userId int64, path string) (bool, error) {
	// 查询用户信息
	user, err := u.userRepo.Find(ctx, userId)
	if err != nil {
		u.logger.Error("[date repo err] userRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, xerror.NewWithMessage("用户不存在")
		}
		return false, err
	}

	// 查询权限信息
	authority, err := u.authorityRepo.FindByUrl(ctx, path)
	if err != nil {
		u.logger.Error("[date repo err] authorityRepo.FindByUrl", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, xerror.NewWithMessage("权限不存在")
		}
		return false, err
	}
	if authority.Status == model.Disable {
		return false, xerror.NewWithMessage("权限已禁用")
	}
	// 不鉴权
	if authority.Auth == model.AuthorityAuthNo {
		return true, nil
	}

	// 查询用户角色信息
	roleIds, err := u.userRepo.FindRoleIds(ctx, user.Id)
	if err != nil {
		u.logger.Error("[date repo err] userRepo.FindRoleIds", zap.Error(err))
		return false, xerror.NewWithMessage("用户角色查询失败")
	}

	// 查询角色权限信息
	authorityIds, err := u.authorityRepo.FindIdsByRoleIds(ctx, roleIds)
	if err != nil {
		u.logger.Error("[date repo err] authorityRepo.FindIdsByRoleIds", zap.Error(err))
		return false, xerror.NewWithMessage("角色权限查询失败")
	}

	// 鉴权
	for _, id := range authorityIds {
		if id == authority.Id {
			return true, nil
		}
	}

	return false, nil
}

// AllocationAuth 分配权限
func (u *RoleUsecase) AllocationAuth(ctx context.Context, param *request.RoleAllocationAuthReq) error {
	role, err := u.roleRepo.Find(ctx, param.RoleId)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.Find", zap.Error(err))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return xerror.NewWithMessage("角色不存在")
		}
		return err
	}
	if role.Status == model.Disable {
		return xerror.NewWithMessage("角色已禁用")
	}

	// 查询权限信息
	authorities, err := u.authorityRepo.FindByIds(ctx, param.AuthIds)
	if err != nil {
		u.logger.Error("[date repo err] authorityRepo.FindByIds", zap.Error(err))
		return xerror.NewWithMessage("权限查询失败")
	}
	if len(authorities) != len(param.AuthIds) {
		return xerror.NewWithMessage("权限不存在")
	}

	var authIds []int64
	for _, authority := range authorities {
		if authority.Status == model.Disable {
			return xerror.NewWithMessage("权限已禁用")
		}
		authIds = append(authIds, authority.Id)
	}

	// 分配权限
	err = u.roleRepo.AllocationAuth(ctx, role.Id, authIds)
	if err != nil {
		u.logger.Error("[date repo err] roleRepo.AllocationAuth", zap.Error(err))
		return xerror.NewWithMessage("权限分配失败")
	}

	return nil
}
