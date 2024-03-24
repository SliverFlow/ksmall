package config

import "fmt"

type Mysql struct {
	DbName       string `mapstructure:"db-name" json:"db-name" yaml:"db-name"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	Host         string `mapstructure:"host" json:"host" yaml:"host"`
	Port         int    `mapstructure:"port" json:"port" yaml:"port"`
	Log          string `mapstructure:"log" json:"log" yaml:"log"`
	Config       string `mapstructure:"config" json:"config" yaml:"config"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"max-idle-conns,default:10" yaml:"max-idle-conns"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"max-open-conns,default:10" yaml:"max-open-conns"`
}

// Dsn 获取 dsn
func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", m.Username, m.Password, m.Host, m.Port, m.DbName, m.Config)
}

func (m *Mysql) GetLog() int {
	switch m.Log {
	case "Silent", "silent":
		return 1
	case "Error", "error":
		return 2
	case "Warn", "warn":
		return 3
	case "Info", "info":
		return 4
	default:
		return 4
	}
}
