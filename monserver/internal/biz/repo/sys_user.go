package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type UserRepo interface {
	Insert(ctx context.Context, user *model.User) (*model.User, error)
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*model.User, error)
	FindRoleId(ctx context.Context, id int64) (int64, error)
}
