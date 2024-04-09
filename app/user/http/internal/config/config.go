package config

import (
	"github.com/SliverFlow/ksmall/core/config"
)

type Possess struct {
	Server *config.HttpServer `mapstructure:"server" json:"server" yaml:"server"`
	Log    *config.Zap        `mapstructure:"log" json:"log" yaml:"log"`
}

func NewServerConfig(c *Possess) *config.HttpServer {
	return c.Server
}
