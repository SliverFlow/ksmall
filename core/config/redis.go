package config

// Redis the redis config
type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
	Store    string `mapstructure:"store" json:"store" yaml:"store"`
	Expire   int    `mapstructure:"expire" json:"expire" yaml:"expire"`
}
