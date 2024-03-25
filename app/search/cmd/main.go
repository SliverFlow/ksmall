package main

import (
	"flag"
	"github.com/SliverFlow/ksmall/app/search/internal/config"
	"github.com/SliverFlow/ksmall/core/initialize"
	"go.uber.org/zap"
)

const path = "./app/search/etc/"

func main() {

	env := flag.String("env", "dev", "运行环境")

	var c config.Possess
	v := initialize.Viper(path, env, nil)
	_ = v.Unmarshal(&c)

	// 初始化日志
	logger := initialize.Zap(c.Log)

	app := wireApp(&c, logger)
	err := app.Server.ListenAndServe()
	if err != nil {
		logger.Error("服务启动失败", zap.Error(err))
	}
}
