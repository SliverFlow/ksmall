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

func (ur *userRepo) FindByUsername(ctx context.Context, username string) (*model.User, error) {
	db := ur.db.WithContext(ctx).Model(&model.User{})
	var user model.User
	if err := db.Where(model.UserCol.Username, username).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
