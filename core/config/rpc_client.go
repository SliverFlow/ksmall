package config

type RpcClient struct {
	Ip   string `mapstructure:"port" json:"port" yaml:"port"`
	Type string `mapstructure:"type" json:"type" yaml:"type"`
	Etcd Etcd   `mapstructure:"etcd" json:"etcd" yaml:"etcd"`
}
