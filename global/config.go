package global

import (
	"github.com/spf13/viper"
	"log"
)

//全局配置
type config struct {
	Env   string //当前环境
	DBDsn string //数据库DSN
}

//初始化配置信息
func initConf() {
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	log.Println("加载全局配置成功")
}
