package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type CategoryRepo interface {
	Insert(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*model.Category, error)
	FindChildrenList(ctx context.Context, parentId int64) ([]*model.Category, error)
	FindAll(ctx context.Context) ([]*model.Category, error)
}
