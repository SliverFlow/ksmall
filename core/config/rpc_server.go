package config

// RpcServer rpc 服务配置
type RpcServer struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}
