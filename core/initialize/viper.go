package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper(path string, watch func()) *viper.Viper {
	v := viper.New()
	v.SetConfigFile(path)
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
