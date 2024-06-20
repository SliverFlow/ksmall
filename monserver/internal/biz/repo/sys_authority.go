package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type AuthorityRepo interface {
	Insert(ctx context.Context, authority *model.Authority) (*model.Authority, error)
	Update(ctx context.Context, authority *model.Authority) (*model.Authority, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*model.Authority, error)
	FindByUrl(ctx context.Context, url string) (*model.Authority, error)
	FindIdsByRoleIds(ctx context.Context, roleIds []int64) ([]int64, error)
	FindByIds(ctx context.Context, ids []int64) ([]*model.Authority, error)
	FindByAuthorityGroupIds(ctx context.Context, authorityGroupIds []int64) ([]*model.Authority, error)
}
