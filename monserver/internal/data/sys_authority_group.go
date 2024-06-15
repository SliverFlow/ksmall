package data

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
	"time"
)

type authorityGroupRepo struct {
	*Common
}

func NewAuthorityGroupRepo(common *Common) *authorityGroupRepo {
	return &authorityGroupRepo{
		common,
	}
}

// Insert data
func (a *authorityGroupRepo) Insert(ctx context.Context, authorityGroup *model.AuthorityGroup) (*model.AuthorityGroup, error) {
	tx := a.db.WithContext(ctx).Model(&model.AuthorityGroup{})
	if err := tx.Create(authorityGroup).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return authorityGroup, nil
}

// Update data
func (a *authorityGroupRepo) Update(ctx context.Context, authorityGroup *model.AuthorityGroup) (*model.AuthorityGroup, error) {
	tx := a.db.WithContext(ctx).Model(&model.AuthorityGroup{})
	if err := tx.Where(model.AuthorityCol.Id+" = ?", authorityGroup.Id).Updates(authorityGroup).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return authorityGroup, nil
}

// Delete data
func (a *authorityGroupRepo) Delete(ctx context.Context, id int64) error {
	tx := a.db.WithContext(ctx).Model(&model.AuthorityGroup{})
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
func (a *authorityGroupRepo) Find(ctx context.Context, id int64) (*model.AuthorityGroup, error) {
	tx := a.db.WithContext(ctx).Model(&model.AuthorityGroup{})
	var authorityGroup model.AuthorityGroup
	if err := tx.Where(model.AuthorityCol.Id+" = ?", id).First(&authorityGroup).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return &authorityGroup, nil
}
