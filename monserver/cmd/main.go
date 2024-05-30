package main

import (
	"flag"
	"github.com/SliverFlow/core/initialize"
	"github.com/SliverFlow/ksmall/monserver/internal/config"
	"go.uber.org/zap"
)

const path = "./etc"

func main() {

	env := flag.String("env", "dev", "")

	var conf config.Possess
	v := initialize.Viper(path, env, nil)
	err := v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	logger := initialize.Zap(conf.Zap)
	s := wireApp(&conf, logger)
	err = s.ListenServer()
	if err != nil {
		logger.Error("服务启动失败", zap.Any("err", err))
		return
	}
}