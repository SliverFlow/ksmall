package config

// RpcClient the rpc client config
type RpcClient struct {
	Ip   string `mapstructure:"ip" json:"ip" yaml:"ip"`
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	Etcd Etcd   `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
}
