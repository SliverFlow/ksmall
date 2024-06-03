package repo

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
)

type StockRepo interface {
	Insert(ctx context.Context, stock *model.Stock) error
	Update(ctx context.Context, stock *model.Stock) error
	Delete(ctx context.Context, id int) error
	Find(ctx context.Context, id int) (*model.Stock, error)
}
