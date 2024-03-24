package initialize

import (
	"github.com/SliverFlow/ksmall/core/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func InitMysqlConn(c *config.Mysql) *gorm.DB {
	dsn := c.Dsn()

	mc := mysql.Config{
		DSN:                       dsn,   // DSN data source name
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}

	logLever := c.GetLog()

	DB, err := gorm.Open(mysql.New(mc), &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold: 2 * time.Second,           // 慢 SQL 阈值
				LogLevel:      logger.LogLevel(logLever), // Log level
				Colorful:      false,                     // 禁用彩色打印
			}),
	})

	if err != nil {
		panic(err)
	}

	db, _ := DB.DB()
	db.SetMaxIdleConns(c.MaxIdleConns)
	db.SetMaxOpenConns(c.MaxOpenConns)

	return DB
}
