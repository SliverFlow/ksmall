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
