package global

import (
	"github.com/jinzhu/gorm"
	"log"
)

//初始化数据库
func initDB() {
	var err error
	DB, err = gorm.Open("mysql", Conf.DBDsn)
	if err != nil {
		panic(err)
	}
	err = DB.DB().Ping()
	if err != nil {
		panic(err)
	}
	log.Println("连接数据库成功")
}
