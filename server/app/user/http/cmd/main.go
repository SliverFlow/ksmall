package main

import (
	"flag"
	"github.com/SliverFlow/ksmall/server/app/user/http/internal/config"
	"github.com/SliverFlow/ksmall/server/core/initialize"
	"go.uber.org/zap"
)

const configPath = "./app/user/http/etc/"

func main() {

	env := flag.String("env", "dev", "")

	var c config.Possess
	v := initialize.Viper(configPath, env, nil)
	_ = v.Unmarshal(&c)

	// 初始化日志
	logger := initialize.Zap(c.Log)

	app := wireApp(&c, logger)
	err := app.Server.ListenAndServe()
	if err != nil {
		logger.Error("服务启动失败", zap.Error(err))
	}
}
