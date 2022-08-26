package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func Init() (err error) {
	viper.SetConfigName("config")  // 指定配置文件(不需要带后缀)
	viper.SetConfigType("yaml")    // 指定配置文件类型
	viper.AddConfigPath("./conf/") // 指定查找配置文件路径
	err = viper.ReadInConfig()     // 读取配置文件
	if err != nil {
		// 读取配置信息失败
		fmt.Printf("viper.ReadInConfig() failed, err:%v\n", err)
		return
	}
	// 线上监控配置文件更改
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件修改了......")
	})
	return
}
