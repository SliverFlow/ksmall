package config

import "github.com/SliverFlow/core/config"

// Possess 配置文件结构体
type Possess struct {
	Server *config.HttpServer `mapstructure:"server" yaml:"server"`
	Zap    *config.Zap        `mapstructure:"zap" yaml:"zap"`
	Mysql  *config.Mysql      `mapstructure:"mysql" yaml:"mysql"`
	Redis  *config.Redis      `mapstructure:"redis" yaml:"redis"`
}

func NewHttpServerConfig(c *Possess) *config.HttpServer {
	return c.Server
}
