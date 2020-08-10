package global

import (
	"github.com/jinzhu/gorm"
)

var Conf config

var DB *gorm.DB

func Init() {
	initConf()
	initDB()
}
