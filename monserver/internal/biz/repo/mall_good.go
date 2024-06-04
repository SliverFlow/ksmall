package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type GoodRepo interface {
	Insert(ctx context.Context, good *model.Good) (*model.Good, error)
}
