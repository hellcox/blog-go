package global

import (
	"github.com/jinzhu/gorm"
)

// 全局配置
var Conf config

// 数据库
var DB *gorm.DB

// 网页三要素
var TKD AllTkd

func Init() {
	initConf()
	initDB()
	initTKD()
}
