package data

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
)

type roleAuthRefRepo struct {
	*Common
}

func NewRoleAuthRefRepo(common *Common) *roleAuthRefRepo {
	return &roleAuthRefRepo{
		common,
	}
}

// FindByRoleId 根据角色ID查询角色权限关联
func (r *roleAuthRefRepo) FindByRoleId(ctx context.Context, roleId int64) ([]*model.RoleAuthRef, error) {
	tx := r.db.WithContext(ctx).Model(&model.RoleAuthRef{})
	var roleAuthRefs []*model.RoleAuthRef
	if err := tx.Where(model.RoleAuthRefCol.RoleId+" = ?", roleId).
		Find(&roleAuthRefs).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return roleAuthRefs, nil
}
