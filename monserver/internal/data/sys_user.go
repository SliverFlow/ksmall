package data

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
	"time"
)

type userRepo struct {
	*Common
}

func NewUserRepo(common *Common) *userRepo {
	return &userRepo{
		Common: common,
	}
}

func (r *userRepo) Insert(ctx context.Context, user *model.User) (*model.User, error) {
	tx := r.DB(ctx).Model(&model.User{})
	if err := tx.Create(user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

func (r *userRepo) Update(ctx context.Context, user *model.User) (*model.User, error) {
	tx := r.DB(ctx).Model(&model.User{})
	if err := tx.Where(model.UserCol.Id+" = ?", user.Id).Updates(user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

func (r *userRepo) Delete(ctx context.Context, id int64) error {
	tx := r.DB(ctx).Model(&model.User{})
	if err := tx.Where(model.UserCol.Id+" = ?", id).
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

func (r *userRepo) Find(ctx context.Context, id int64) (*model.User, error) {
	user := &model.User{}
	tx := r.DB(ctx).Model(&model.User{})
	if err := tx.Where(model.UserCol.Id+" = ?", id).
		Where(model.UserCol.Deleted+" = ?", model.NotDeleted).
		First(user).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return user, nil
}

// FindRoleId 根据用户id查询角色id
func (r *userRepo) FindRoleId(ctx context.Context, id int64) (int64, error) {
	ref := model.UserRoleRef{}
	tx := r.DB(ctx).Model(&model.UserRoleRef{})
	if err := tx.Where(model.UserRoleRefCol.UserId+" = ?", id).
		First(&ref).Error; err != nil {
		return 0, errors.WithStack(err)
	}
	return ref.RoleId, nil
}

// FindRoleIds 根据用户id查询角色id
func (r *userRepo) FindRoleIds(ctx context.Context, id int64) ([]int64, error) {
	var refs []model.UserRoleRef
	tx := r.DB(ctx).Model(&model.UserRoleRef{})
	if err := tx.Where(model.UserRoleRefCol.UserId+" = ?", id).
		Find(&refs).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	var roleIds []int64
	for _, ref := range refs {
		roleIds = append(roleIds, ref.RoleId)
	}
	return roleIds, nil
}

// FindListByUserIds 根据用户ids查询角色列表
func (r *userRepo) FindListByUserIds(ctx context.Context, userIds []int64) ([]*model.User, error) {
	var users []*model.User
	tx := r.DB(ctx).Model(&model.User{})
	if err := tx.Where(model.UserCol.Id+" in (?)", userIds).
		Where(model.UserCol.Deleted+" = ?", model.NotDeleted).
		Find(&users).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return users, nil
}
