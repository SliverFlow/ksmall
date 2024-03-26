package main

import (
	"flag"
	"github.com/SliverFlow/ksmall/core/initialize"
	"github.com/SliverFlow/ksmall/rpc/user/internal/config"
)

const path = "./rpc/user/etc/"

func main() {

	env := flag.String("env", "dev", "set env")

	v := initialize.Viper(path, env, nil)
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
