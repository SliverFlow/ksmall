package data

import (
	"context"
	"fmt"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/config"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/model"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"time"
)

type userRepo struct {
	logger *zap.Logger
	db     *gorm.DB
	rdb    *redis.Client
}

var (
	rdbBaseKey    string        = ""
	rdbBaseExpire time.Duration = 1 * time.Minute
)

func NewUserRepo(
	logger *zap.Logger,
	db *gorm.DB,
	rdb *redis.Client,
	c *config.Possess,
) *userRepo {

	var user model.User

	rdbBaseKey = fmt.Sprintf("%s:%s:", c.Redis.Store, user.TableName())
	rdbBaseExpire = time.Duration(c.Redis.Expire) * time.Minute

	return &userRepo{
		logger: logger,
		db:     db,
		rdb:    rdb,
	}
}

// getRdbKey 获取 redis 存储 key
func (ur *userRepo) getRdbKey(id int64) string {
	return fmt.Sprintf("%s%d", rdbBaseKey, id)
}

// getBaseExpire 获取 redis 过期时间
func (ur *userRepo) getBaseExpire() time.Duration {
	return rdbBaseExpire
}

// getBaseExpire 获取 redis 过期时间 分钟
func (ur *userRepo) getExpireMinute(expire int64) time.Duration {
	if expire <= 0 {
		return ur.getBaseExpire()
	}
	return time.Duration(expire) * time.Minute
}

// getBaseExpire 获取 redis 过期时间 秒
func (ur *userRepo) getExpireSecond(expire int64) time.Duration {
	if expire <= 0 {
		return ur.getBaseExpire()
	}
	return time.Duration(expire) * time.Second
}

// PageList 分页查询用户
func (ur *userRepo) PageList(ctx context.Context, limit int64, offset int64) ([]*model.User, int64, error) {
	db := ur.db.WithContext(ctx).Model(&model.User{})
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
