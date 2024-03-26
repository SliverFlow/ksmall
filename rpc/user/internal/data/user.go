package data

import (
	"fmt"
	"github.com/SliverFlow/ksmall/rpc/user/internal/config"
	"github.com/SliverFlow/ksmall/rpc/user/internal/model"
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

func (ur *userRepo) getRdbKey(id int64) string {
	return fmt.Sprintf("%s%d", rdbBaseKey, id)
}

func (ur *userRepo) getBaseExpire() time.Duration {
	return rdbBaseExpire
}

func (ur *userRepo) getExpireMinute(expire int64) time.Duration {
	if expire <= 0 {
		return ur.getBaseExpire()
	}
	return time.Duration(expire) * time.Minute
}

func (ur *userRepo) getExpireSecond(expire int64) time.Duration {
	if expire <= 0 {
		return ur.getBaseExpire()
	}
	return time.Duration(expire) * time.Second
}
