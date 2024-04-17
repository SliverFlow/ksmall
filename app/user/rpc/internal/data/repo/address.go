package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/model"
)

type IAddressRepo interface {
	Count(ctx context.Context, userId int64) (int64, error)
	FindDefaultByUserId(ctx context.Context, userId int64) (*model.Address, error)
	FindListByUserId(ctx context.Context, userId int64) ([]*model.Address, error)
	FindById(ctx context.Context, id int64) (*model.Address, error)
	Update(ctx context.Context, address *model.Address) error
	Create(ctx context.Context, address *model.Address) error
	Delete(ctx context.Context, id int64) error
}
