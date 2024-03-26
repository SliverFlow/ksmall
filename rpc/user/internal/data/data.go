package data

import (
	"github.com/SliverFlow/ksmall/core/config"
	"github.com/SliverFlow/ksmall/core/initialize"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func NewDB(c *config.Mysql) *gorm.DB {
	return initialize.Mysql(c)
}

func NewRDB(c *config.Redis) *redis.Client {
	return initialize.Redis(c)
}
