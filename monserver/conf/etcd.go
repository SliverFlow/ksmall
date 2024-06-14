package conf

// Etcd is the configuration for etcd.
// 最终会进行迁移
type Etcd struct {
	Endpoints []string `mapstructure:"endpoints" yaml:"endpoints"`
	Username  string   `mapstructure:"username" yaml:"username"`
	Password  string   `mapstructure:"password" yaml:"password"`
	Timeout   int      `mapstructure:"timeout" yaml:"timeout"`
}
