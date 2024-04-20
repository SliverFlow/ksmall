package data

import (
	"github.com/SliverFlow/ksmall/server/app/user/rpc/internal/config"
	"github.com/SliverFlow/ksmall/server/core/initialize"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewDB(c *config.Possess) *gorm.DB {
	return initialize.Mysql(c.Mysql)
}

func NewRDB(c *config.Possess) *redis.Client {
	return initialize.Redis(c.Redis)
}
