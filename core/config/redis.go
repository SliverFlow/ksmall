package config

type Redis struct {
	Host       string `mapstructure:"host" json:"host" yaml:"host"`
	Port       int    `mapstructure:"port" json:"port" yaml:"port"`
	Password   string `mapstructure:"password" json:"password" yaml:"password"`
	DB         int    `mapstructure:"db" json:"db" yaml:"db"`
	BaseKey    string `mapstructure:"base-key" json:"baseKey" yaml:"base-key"`
	ExpireTime int    `mapstructure:"expire-time" json:"expireTime" yaml:"expire-time"`
}
