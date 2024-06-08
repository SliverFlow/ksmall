package data

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
)

type categoryRepo struct {
	*Common
}

func NewCategoryRepo(common *Common) *categoryRepo {
	return &categoryRepo{
		common,
	}
}

// Insert a new category
func (c *categoryRepo) Insert(ctx context.Context, category *model.Category) error {
	tx := c.DB(ctx).Model(&model.Category{})
	if err := tx.Create(category).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Update a category
func (c *categoryRepo) Update(ctx context.Context, category *model.Category) error {
	tx := c.DB(ctx).Model(&model.Category{})
	if err := tx.Where(model.CategoryCol.Id+" = ?", category.Id).Updates(category).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Delete a category
func (c *categoryRepo) Delete(ctx context.Context, id int64) error {
	tx := c.DB(ctx).Model(&model.Category{})
	if err := tx.Where(model.CategoryCol.Id+" = ?", id).Delete(&model.Category{}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Find a category
func (c *categoryRepo) Find(ctx context.Context, id int64) (*model.Category, error) {
	category := &model.Category{}
	tx := c.DB(ctx).Model(&model.Category{})
	if err := tx.Where(model.CategoryCol.Id+" = ?", id).
		Where(model.CategoryCol.Deleted+" = ?", model.NotDeleted).
		First(category).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return category, nil
}

// FindChildrenList find all children categories
func (c *categoryRepo) FindChildrenList(ctx context.Context, parentId int64) ([]*model.Category, error) {
	var categories []*model.Category
	tx := c.DB(ctx).Model(&model.Category{})
	if err := tx.Where(model.CategoryCol.ParentId+" = ?", parentId).
		Where(model.CategoryCol.Deleted+" = ?", model.NotDeleted).
		Find(&categories).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return categories, nil
}
