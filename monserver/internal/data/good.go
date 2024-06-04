package data

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
)

type goodRepo struct {
	*Common
}

func NewGoodRepo(common *Common) *goodRepo {
	return &goodRepo{common}
}

// Insert a new good
func (c *goodRepo) Insert(ctx context.Context, good *model.Good) (*model.Good, error) {
	tx := c.DB(ctx).Model(&model.Good{})
	if err := tx.Create(good).Error; err != nil {
		return nil, errors.WithStack(err)
	}

	return good, nil
}
