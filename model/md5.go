package model

import (
	"blog/global"
	"reflect"
)

type Md5 struct {
	ID      int `gorm:"primary_key"`
	M32     string
	M16     string
	Content string
}

func (m Md5) TableName() string {
	return "bg_md5"
}

func (m Md5) IsEmpty() bool {
	return reflect.DeepEqual(m, Md5{})
}

func (m Md5) GetByM32(m32 string) Md5 {
	global.DB.Where("m32=?", m32).First(&m)
	return m
}

func (m Md5) Save() int {
	global.DB.Save(&m)
	return m.ID
}
