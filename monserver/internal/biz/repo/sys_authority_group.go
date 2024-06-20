package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type AuthorityGroupRepo interface {
	Insert(ctx context.Context, authority *model.AuthorityGroup) (*model.AuthorityGroup, error)
	Update(ctx context.Context, authority *model.AuthorityGroup) (*model.AuthorityGroup, error)
	Delete(ctx context.Context, id int64) error
	Find(ctx context.Context, id int64) (*model.AuthorityGroup, error)
	FindAll(ctx context.Context) ([]*model.AuthorityGroup, error)
}
