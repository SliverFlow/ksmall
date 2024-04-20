package initialize

import (
	"fmt"
	"github.com/SliverFlow/ksmall/common/util"
	"github.com/SliverFlow/ksmall/core/config"
	"github.com/SliverFlow/ksmall/core/initialize/internal"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

// Zap 获取 zap.Logger
// Author [SliverHorn](https://github.com/SliverHorn)
func Zap(c *config.Zap) (logger *zap.Logger) {
	if c == nil {
		panic("zap config is nil")
	}

	if ok, _ := util.PathExists(c.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", c.Director)
		_ = os.Mkdir(c.Director, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores(*c)
	logger = zap.New(zapcore.NewTee(cores...))

	if c.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}
