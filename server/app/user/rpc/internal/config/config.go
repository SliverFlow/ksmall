package config

import "github.com/SliverFlow/ksmall/server/core/config"

type Possess struct {
	Server *config.RpcServer `mapstructure:"server" json:"server" yaml:"server"`
	Log    *config.Zap       `mapstructure:"log" json:"log" yaml:"log"`
	Mysql  *config.Mysql     `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Redis  *config.Redis     `mapstructure:"redis" json:"redis" yaml:"redis"`
}

func NewServerConfig(c *Possess) *config.RpcServer {
	return c.Server
}
