package data

import (
	"context"
	"fmt"
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/config"
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/model"
	"github.com/SliverFlow/ksmall/server/common/constant"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type userRepo struct {
	logger *zap.Logger
	db     *gorm.DB
	*userRDB
}

func NewUserRepo(
	logger *zap.Logger,
	db *gorm.DB,
	rdb *redis.Client,
	c *config.Possess,
) *userRepo {
	var user model.User
	userRepo := userRepo{
		logger: logger,
		db:     db,
		userRDB: &userRDB{
			userRDBBaseKey:    fmt.Sprintf("%s:%s:", c.Redis.Store, user.TableName()),
			userRDBBaseExpire: time.Duration(c.Redis.Expire) * time.Minute,
			rdb:               rdb,
		},
	}
	return &userRepo
}

// PageList 分页查询用户
func (ur *userRepo) PageList(ctx context.Context, limit int64, offset int64) ([]*model.User, int64, error) {
	db := ur.db.WithContext(ctx).Model(&model.User{}).Where(model.UserCol.DeleteFlag+" = ?", constant.ModelNotDeleted)
	var users []*model.User
	var total int64
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err := db.Limit(int(limit)).Offset(int(offset)).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	return users, total, nil
}

// Find 根据 id 查询用户
func (ur *userRepo) Find(ctx context.Context, id int64) (*model.User, error) {
	db := ur.db.WithContext(ctx).Model(&model.User{}).Where(model.UserCol.DeleteFlag+" = ?", constant.ModelNotDeleted)
	var user model.User
	if err := db.Where(model.UserCol.Id+" = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByEmail 根据邮箱查询用户
func (ur *userRepo) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	db := ur.db.WithContext(ctx).Model(&model.User{}).Where(model.UserCol.DeleteFlag+" = ?", constant.ModelNotDeleted)
	var user model.User
	if err := db.Where(model.UserCol.Email+" = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// FindByPhone 根据手机号查询用户
func (ur *userRepo) FindByPhone(ctx context.Context, phone string) (*model.User, error) {
	db := ur.db.WithContext(ctx).Model(&model.User{}).Where(model.UserCol.DeleteFlag+" = ?", constant.ModelNotDeleted)
	var user model.User
	if err := db.Where(model.UserCol.Phone+" = ?", phone).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

// Create 创建用户
func (ur *userRepo) Create(ctx context.Context, user *model.User) (int64, error) {
	db := ur.db.WithContext(ctx).Model(&model.User{})
	if err := db.Create(user).Error; err != nil {
		return 0, err
	}
	return user.Id, nil
}

// Delete 删除用户
func (ur *userRepo) Delete(ctx context.Context, id int64) error {
	db := ur.db.WithContext(ctx).Model(&model.User{})
	err := db.Where(model.UserCol.Id+" = ?", id).Update(model.UserCol.DeleteFlag, constant.ModelDeleted).
		Update(model.UserCol.DeleteTime, time.Now().Unix()).
		Update(model.UserCol.UpdateTime, time.Now().Unix()).
		Error
	if err != nil {
		return err
	}

	return nil
}
