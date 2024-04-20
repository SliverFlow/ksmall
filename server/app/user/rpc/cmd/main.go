package main

import (
	"flag"
	"github.com/SliverFlow/ksmall/app/user/rpc/internal/config"
	"github.com/SliverFlow/ksmall/core/initialize"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
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
	go func() {
		err := s.Listen()
		if err != nil {
			panic(err)
		}
	}()
	logger.Info("grpc server started in 0.0.0.0", zap.Int("port", c.Server.Port))
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	<-quit
	logger.Info("Shutting down server...")
}
