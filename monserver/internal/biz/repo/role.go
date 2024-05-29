package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type RoleRepo interface {
	Insert(ctx context.Context, role *model.Role) (*model.Role, error)
	Update(ctx context.Context, role *model.Role) (*model.Role, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*model.Role, error)
	List(ctx context.Context) ([]*model.Role, error)
	FindByName(ctx context.Context, name string) (*model.Role, error)
	ListByCondition(ctx context.Context, name string, status int64) ([]*model.Role, error)
}
