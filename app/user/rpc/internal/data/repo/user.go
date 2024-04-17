package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/model"
)

type IUserRepo interface {
	FindByUsername(ctx context.Context, username string) (*model.User, error)
	PageList(ctx context.Context, limit int64, offset int64) ([]*model.User, int64, error)
}
