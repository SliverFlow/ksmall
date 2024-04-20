package config

// RpcServer the rpc server config
type RpcServer struct {
	Port int    `mapstructure:"port" json:"port" yaml:"port"`
	Name string `mapstructure:"name" json:"name" yaml:"name"`
}
