package data

import (
	"context"
	"fmt"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
	"time"
)

type roleRepo struct {
	*Common
}

func NewRoleRepo(common *Common) *roleRepo {
	return &roleRepo{
		Common: common,
	}
}

func (r *roleRepo) Insert(ctx context.Context, role *model.Role) (*model.Role, error) {
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.Create(role).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return role, nil
}

func (r *roleRepo) Update(ctx context.Context, role *model.Role) (*model.Role, error) {
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.Where(model.RoleCol.Id+" = ?", role.Id).Updates(role).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return role, nil
}

func (r *roleRepo) Delete(ctx context.Context, id int64) error {
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.Where(model.RoleCol.Id+" = ?", id).
		UpdateColumns(map[string]any{
			model.RoleCol.Status:   model.Disable,
			model.RoleCol.UpdateAt: time.Now().Unix(),
			model.RoleCol.DeleteAt: time.Now().Unix(),
			model.RoleCol.Deleted:  model.Deleted,
		}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func (r *roleRepo) Find(ctx context.Context, id int64) (*model.Role, error) {
	role := &model.Role{}
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.Where(model.RoleCol.Id+" = ?", id).
		Where(model.RoleCol.Deleted+" = ?", model.NotDeleted).
		First(role).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return role, nil
}

func (r *roleRepo) List(ctx context.Context) ([]*model.Role, error) {
	var list []*model.Role
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.
		Where(model.RoleCol.Deleted+" = ?", model.NotDeleted).
		Find(&list).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return list, nil
}

// FindByName 根据角色名称查找角色
func (r *roleRepo) FindByName(ctx context.Context, name string) (*model.Role, error) {
	role := &model.Role{}
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.Where(model.RoleCol.Name+" = ?", name).
		Where(model.RoleCol.Deleted+" = ?", model.NotDeleted).
		First(role).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return role, nil
}

// ListByCondition 根据条件查询角色列表
func (r *roleRepo) ListByCondition(ctx context.Context, name string, status int64) ([]*model.Role, error) {
	var list []*model.Role
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.Where(model.RoleCol.Name+" like ?", fmt.Sprintf("%%%s%%", name)).
		Where(model.RoleCol.Status+" = ?", status).
		Where(model.RoleCol.Deleted+" = ?", model.NotDeleted).
		Find(&list).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return list, nil
}

// FindByKey 根据key查找角色
func (r *roleRepo) FindByKey(ctx context.Context, key int64) (*model.Role, error) {
	role := &model.Role{}
	tx := r.DB(ctx).Model(&model.Role{})
	if err := tx.Where(model.RoleCol.Key+" = ?", key).
		Where(model.RoleCol.Deleted+" = ?", model.NotDeleted).
		First(role).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return role, nil
}
