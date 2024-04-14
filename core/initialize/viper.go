package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Viper 初始化 viper 配置
func Viper(path string, env *string, watch func()) *viper.Viper {
	v := viper.New()
	if env == nil {
		*env = "dev"
	}
	v.SetConfigFile(fmt.Sprintf("%s/config-%s.yaml", path, *env))
	v.SetConfigType("yaml")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	if watch != nil {
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			fmt.Println("config file changed:", e.Name)
			watch()
		})
	}
	return v
}
