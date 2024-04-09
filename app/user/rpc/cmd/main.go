package main

import (
	"flag"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/config"
	"github.com/SliverFlow/ksmall/core/initialize"
)

const configPath = "./app/user/rpc/etc/"

func main() {

	env := flag.String("env", "dev", "set env")

	v := initialize.Viper(configPath, env, nil)
	var c config.Possess
	if err := v.Unmarshal(&c); err != nil {
		panic(err)
	}

	logger := initialize.Zap(c.Log)

	s := wireApp(&c, logger)
	if s == nil {
		logger.Info("s is nil")
	}

}
