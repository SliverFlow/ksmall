package biz

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/common/xerror"
	"github.com/SliverFlow/ksmall/monserver/internal/biz/repo"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/SliverFlow/ksmall/monserver/internal/model/reply"
	"github.com/SliverFlow/ksmall/monserver/internal/model/request"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"time"
)

type RoleUsecase struct {
	logger   *zap.Logger
	roleRepo repo.RoleRepo
}

func NewRoleUsecase(logger *zap.Logger, roleRepo repo.RoleRepo) *RoleUsecase {
	return &RoleUsecase{
		logger:   logger,
		roleRepo: roleRepo,
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
		return nil, xerror.NewWithMessage("角色不存在")
	}

	return role, nil
}

// List 查询角色列表
func (u *RoleUsecase) List(ctx context.Context, param *request.RoleListReq) ([]*model.Role, error) {
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
