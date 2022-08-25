package initialize

import (
	"fmt"
	"github.com/Gpihuier/gpihuier_blog/utils"

	"github.com/Gpihuier/gpihuier_blog/global"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Viper() *viper.Viper {
	v := viper.New()
	rootPath := utils.GetRootPath()
	v.AddConfigPath(rootPath)
	v.SetConfigType("yaml")
	v.SetConfigName("config")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控并重新读取配置文件
	v.WatchConfig()

	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("config file changed:", e.Name)
		if err := v.Unmarshal(&global.CONFIG); err != nil {
			panic(fmt.Errorf("Fatal error config file: %s \n", err))
		}
	})
	// 读取配置文件
	// 反序列化 把值解析到结构体中
	// mapstructure viper 通过 mapstructure 读取配置文件
	if err := v.Unmarshal(&global.CONFIG); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return v
}
