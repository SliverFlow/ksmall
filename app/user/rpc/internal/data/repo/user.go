package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/model"
)

type IUserRepo interface {
	PageList(ctx context.Context, limit int64, offset int64) ([]*model.User, int64, error)
	Create(ctx context.Context, user *model.User) (int64, error)
	FindByPhone(ctx context.Context, phone string) (*model.User, error)
	Find(ctx context.Context, id int64) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
	Delete(ctx context.Context, id int64) error
}
