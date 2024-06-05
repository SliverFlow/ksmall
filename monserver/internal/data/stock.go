package data

import (
	"context"
	"github.com/SliverFlow/ksmall/monserver/internal/model"
	"github.com/pkg/errors"
	"time"
)

type stockRepo struct {
	*Common
}

func NewStockRepo(common *Common) *stockRepo {
	return &stockRepo{common}
}

// Insert 新建库存
func (r *stockRepo) Insert(ctx context.Context, stock *model.Stock) error {
	tx := r.db.WithContext(ctx).Model(&model.Stock{})
	if err := tx.Create(stock).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Update 更新库存
func (r *stockRepo) Update(ctx context.Context, stock *model.Stock) error {
	tx := r.db.WithContext(ctx).Model(&model.Stock{})
	if err := tx.Where(model.StockCol.Id+" = ?", stock.Id).Updates(stock).Error; err != nil {
		return errors.WithStack(err)
	}

	return nil
}

// Delete 删除库存
func (r *stockRepo) Delete(ctx context.Context, id int) error {
	tx := r.DB(ctx).Model(&model.Stock{})
	if err := tx.Where(model.StockCol.Id+" = ?", id).
		UpdateColumns(map[string]any{
			model.StockCol.Status:   model.Disable,
			model.StockCol.UpdateAt: time.Now().Unix(),
			model.StockCol.DeleteAt: time.Now().Unix(),
			model.StockCol.Deleted:  model.Deleted,
		}).Error; err != nil {
		return errors.WithStack(err)
	}
	return nil
}

// Find 查询库存
func (r *stockRepo) Find(ctx context.Context, id int) (*model.Stock, error) {
	stock := &model.Stock{}
	tx := r.DB(ctx).Model(&model.Stock{})
	if err := tx.Where(model.StockCol.Id+" = ?", id).
		Where(model.StockCol.Deleted+" = ?", model.NotDeleted).
		First(stock).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return stock, nil
}

// FindByGoodsId 根据商品 id 查询所有使用的库存，不包括已删除的，可能有多个
func (r *stockRepo) FindByGoodsId(ctx context.Context, goodsId int) ([]*model.Stock, error) {
	stocks := make([]*model.Stock, 0)
	tx := r.DB(ctx).Model(&model.Stock{})
	if err := tx.Where(model.StockCol.GoodsId+" = ?", goodsId).
		Where(model.StockCol.Status+" = ?", model.StockIsActive).
		Where(model.StockCol.Deleted+" = ?", model.NotDeleted).
		Find(&stocks).Error; err != nil {
		return nil, errors.WithStack(err)
	}
	return stocks, nil
}
