package config

import (
	"github.com/SliverFlow/ksmall/server/core/config"
)

type Possess struct {
	Server  *config.HttpServer `mapstructure:"server" json:"server" yaml:"server"`
	Log     *config.Zap        `mapstructure:"log" json:"log" yaml:"log"`
	UserRpc *config.RpcClient  `mapstructure:"userRpc" json:"userRpc" yaml:"userRpc"`
}

func NewServerConfig(c *Possess) *config.HttpServer {
	return c.Server
}
