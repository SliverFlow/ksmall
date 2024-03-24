package main

import (
	"github.com/SliverFlow/ksmall/app/shopcart/internal/config"
	"github.com/SliverFlow/ksmall/core/initialize"
	"go.uber.org/zap"
)

const path = "./app/shopcart/etc/config.yaml"

func main() {

	var c config.Possess
	v := initialize.Viper(path, nil)
	_ = v.Unmarshal(&c)

	// 初始化日志
	logger := initialize.Zap(c.Log)

	app := wireApp(&c, logger)
	err := app.Server.ListenAndServe()
	if err != nil {
		logger.Error("服务启动失败", zap.Error(err))
	}
}
