package data

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
	"time"
)

type authorityRepo struct {
	*Common
}

func NewAuthorityRepo(common *Common) *authorityRepo {
	return &authorityRepo{
		common,
	}
}

// Insert data
func (a *authorityRepo) Insert(ctx context.Context, authority *model.Authority) (*model.Authority, error) {
	tx := a.db.WithContext(ctx).Model(&model.Authority{})
	if err := tx.Create(authority).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return authority, nil
}

// Update data
func (a *authorityRepo) Update(ctx context.Context, authority *model.Authority) (*model.Authority, error) {
	tx := a.db.WithContext(ctx).Model(&model.Authority{})
	if err := tx.Where(model.AuthorityCol.Id+" = ?", authority.Id).Updates(authority).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return authority, nil
}

// Delete data
func (a *authorityRepo) Delete(ctx context.Context, id int64) error {
	tx := a.db.WithContext(ctx).Model(&model.Authority{})
	if err := tx.Where(model.AuthorityCol.Id+" = ?", id).UpdateColumns(
		map[string]any{
			model.AuthorityCol.Status:   model.Disable,
			model.AuthorityCol.UpdateAt: time.Now().Unix(),
			model.AuthorityCol.DeleteAt: time.Now().Unix(),
			model.AuthorityCol.Deleted:  model.Deleted,
		}).
		Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Find data
func (a *authorityRepo) Find(ctx context.Context, id int64) (*model.Authority, error) {
	tx := a.db.WithContext(ctx).Model(&model.Authority{})
	var authority model.Authority
	if err := tx.Where(model.AuthorityCol.Id+" = ?", id).
		Where(model.AuthorityCol.Deleted+" = ?", model.NotDeleted).
		First(&authority).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &authority, nil
}

// FindByUrl data
func (a *authorityRepo) FindByUrl(ctx context.Context, url string) (*model.Authority, error) {
	tx := a.db.WithContext(ctx).Model(&model.Authority{})
	var authority model.Authority
	if err := tx.Where(model.AuthorityCol.Url+" = ?", url).
		Where(model.AuthorityCol.Deleted+" = ?", model.NotDeleted).
		First(&authority).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &authority, nil
}

// FindIdsByRoleIds data
func (a *authorityRepo) FindIdsByRoleIds(ctx context.Context, roleIds []int64) ([]int64, error) {
	var ids []int64
	tx := a.db.WithContext(ctx).Model(&model.RoleAuthRef{})
	if err := tx.Where(model.RoleAuthRefCol.RoleId+" in (?)", roleIds).
		Pluck(model.RoleAuthRefCol.AuthorityId, &ids).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return ids, nil
}

// FindByIds data
func (a *authorityRepo) FindByIds(ctx context.Context, ids []int64) ([]*model.Authority, error) {
	var authorities []*model.Authority
	tx := a.db.WithContext(ctx).Model(&model.Authority{})
	if err := tx.Where(model.AuthorityCol.Id+" in (?)", ids).
		Where(model.AuthorityCol.Deleted+" = ?", model.NotDeleted).
		Find(&authorities).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return authorities, nil
}
