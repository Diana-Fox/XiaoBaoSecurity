package ioc

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)
import _ "github.com/spf13/viper/remote"

func InitViper() {
	viper.SetConfigName("dev")  //配置文件的名字
	viper.SetConfigType("yaml") //后缀
	viper.AddConfigPath("./config")
	//实时监听配置文件变更
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		//配置变更的事件,只会说变更了，不会说变更了什么内容
		fmt.Println(in.Name, in.Op)
		fmt.Println(viper.GetString("mysql.dsn"))
	})
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	} //读取配置进入viper,加载进内存
}
func InitViperRemote() {
	viper.SetConfigType("yaml") //后缀
	err := viper.AddRemoteProvider("etcd3", "127.0.0.1:12379", "/security")
	if err != nil {
		panic(err)
	}
	err = viper.WatchRemoteConfig()
	if err != nil {
		panic(err)
	}
	err = viper.ReadRemoteConfig()
	if err != nil {
		return
	}

	err = viper.ReadRemoteConfig()
	if err != nil {
		panic(err)
	}
}
