package main

import (
	"flag"
	"fmt"
	"github.com/SliverFlow/core/initialize"
	"github.com/SliverFlow/ksmall/monserver/internal/config"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const path = "./etc"

func main() {

	env := flag.String("env", "dev", "")
	port := flag.Int64("port", 8080, "")
	path := flag.String("path", path, "	")
	flag.Parse()

	var conf config.Possess
	v := initialize.Viper(*path, env, nil)
	err := v.Unmarshal(&conf)
	if err != nil {
		panic(err)
	}

	if port != nil && *port != 0 {
		fmt.Println("port:", *port)
		conf.Server.Port = int(*port)
	}

	gin.SetMode(gin.ReleaseMode)
	logger := initialize.Zap(conf.Zap)
	s := wireApp(&conf, logger)
	err = s.ListenServer()
	if err != nil {
		logger.Error("服务启动失败", zap.Any("err", err))
	}

}
