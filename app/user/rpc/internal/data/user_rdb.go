package data

import (
	"context"
	"fmt"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/model"
	"github.com/SliverFlow/ksmall/common/util"
	"github.com/redis/go-redis/v9"
	"time"
)

type userRDB struct {
	userRDBBaseKey    string
	userRDBBaseExpire time.Duration
	rdb               *redis.Client
}

// getRdbKey 获取 redis 存储 key
func (ur *userRDB) getRdbKey(id int64) string {
	return fmt.Sprintf("%s:%d", ur.userRDBBaseKey, id)
}

// getExpireMinute 获取 redis 过期时间 分钟
func (ur *userRDB) getExpireMinute(expire int64) time.Duration {
	if expire <= 0 {
		return ur.userRDBBaseExpire
	}
	return time.Duration(expire) * time.Minute
}

// getExpireSecond 获取 redis 过期时间 秒
func (ur *userRDB) getExpireSecond(expire int64) time.Duration {
	if expire <= 0 {
		return ur.userRDBBaseExpire
	}
	return time.Duration(expire) * time.Second
}

// Set 缓存用户信息
func (ur *userRDB) Set(ctx context.Context, user model.User) error {
	key := ur.getRdbKey(user.Id)
	s, err := util.StructToJSON(user)
	if err != nil {
		return err
	}
	return ur.rdb.Set(ctx, key, s, ur.getExpireMinute(5)).Err()
}

// Get 获取用户信息
func (ur *userRDB) Get(ctx context.Context, id int64) (*model.User, error) {
	key := ur.getRdbKey(id)
	s, err := ur.rdb.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var user model.User
	err = util.JSONToStruct(s, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
