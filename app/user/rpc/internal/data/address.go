package data

import (
	"context"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/model"
	"github.com/SliverFlow/ksmall/common/constant"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type addressRepo struct {
	logger *zap.Logger
	db     *gorm.DB
	rdb    *redis.Client
}

func NewAddressRepo(
	logger *zap.Logger,
	db *gorm.DB,
	rdb *redis.Client,
) *addressRepo {
	return &addressRepo{
		logger: logger,
		db:     db,
		rdb:    rdb,
	}
}

// Count 统计地址数量
func (ar *addressRepo) Count(ctx context.Context, userId int64) (int64, error) {
	db := ar.db.WithContext(ctx).Model(&model.Address{})
	var count int64
	if err := db.Where(model.AddressCol.UserId+" = ?", userId).Count(&count).Error; err != nil {
		return 0, err
	}
	return count, nil
}

// FindDefaultByUserId 获取用户默认地址
func (ar *addressRepo) FindDefaultByUserId(ctx context.Context, userId int64) (*model.Address, error) {
	db := ar.db.WithContext(ctx).Model(&model.Address{})
	var address model.Address
	if err := db.Where(model.AddressCol.UserId+" = ? AND "+model.AddressCol.IsDefault+" = ?", userId, model.AddressIsDefault).First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

// FindListByUserId 获取用户地址列表
func (ar *addressRepo) FindListByUserId(ctx context.Context, userId int64) ([]*model.Address, error) {
	db := ar.db.WithContext(ctx).Model(&model.Address{})
	var addresses []*model.Address
	if err := db.Where(model.AddressCol.UserId+" = ?", userId).Find(&addresses).Error; err != nil {
		return nil, err
	}
	return addresses, nil
}

// FindById 获取地址
func (ar *addressRepo) FindById(ctx context.Context, id int64) (*model.Address, error) {
	db := ar.db.WithContext(ctx).Model(&model.Address{})
	var address model.Address
	if err := db.Where(model.AddressCol.Id+" = ?", id).First(&address).Error; err != nil {
		return nil, err
	}
	return &address, nil
}

// Create 创建地址
func (ar *addressRepo) Create(ctx context.Context, address *model.Address) error {
	db := ar.db.WithContext(ctx).Model(&model.Address{})
	return db.Create(address).Error
}

// Update 更新地址
func (ar *addressRepo) Update(ctx context.Context, address *model.Address) error {
	db := ar.db.WithContext(ctx).Model(&model.Address{})
	return db.Where(model.AddressCol.Id+" = ?", address.Id).Updates(address).Error
}

// Delete 删除地址
func (ar *addressRepo) Delete(ctx context.Context, id int64) error {
	db := ar.db.WithContext(ctx).Model(&model.Address{})
	return db.Where(model.AddressCol.Id+" = ?", id).Update(model.AddressCol.DeleteFlag, constant.ModelDeleted).Error
}
