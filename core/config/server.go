package config

// Server 服务配置
type Server struct {
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Timeout int    `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
