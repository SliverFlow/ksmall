package config

import "github.com/SliverFlow/ksmall/core/config"

type Possess struct {
	Server *config.Server `mapstructure:"server" json:"server" yaml:"server"`
}

func NewServerConfig(c *Possess) *config.Server {
	return c.Server
}
