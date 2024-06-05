package data

import (
	"context"
	"github.com/SliverFlow/core/initialize"
	"github.com/SliverFlow/ksmall/monserver/internal/config"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Common struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewCommon(db *gorm.DB, rdb *redis.Client) *Common {
	return &Common{
		db:  db,
		rdb: rdb,
	}
}

func (c *Common) DB(ctx context.Context) *gorm.DB {
	return c.db.WithContext(ctx)
}

func (c *Common) RDB() *redis.Client {
	return c.rdb
}

func NewDB(c *config.Possess) *gorm.DB {
	db, err := initialize.Mysql(c.Mysql)
	if err != nil {
		panic(err)
		return nil
	}
	// _ = db.AutoMigrate(model.Role{})
	// _ = db.AutoMigrate(model.User{})
	// _ = db.AutoMigrate(model.Category{})
	//_ = db.AutoMigrate(model.Good{})
	//_ = db.AutoMigrate(model.Stock{})
	return db
}

func NewRDB(c *config.Possess) *redis.Client {
	client, err := initialize.Redis(c.Redis)
	if err != nil {
		panic(err)
		return nil
	}
	return client
}
