package config

// HttpServer the http server config
type HttpServer struct {
	Port    int    `mapstructure:"port" json:"port" yaml:"port"`
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Timeout int    `mapstructure:"timeout" json:"timeout" yaml:"timeout"`
}
