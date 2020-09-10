package model

import (
	"blog/global"
	"reflect"
)

type Tools struct {
	ID     int `gorm:"primary_key"`
	Url    string
	Sort   int
	Title string
	IsShow int
}

func (m Tools) TableName() string {
	return "bg_tools"
}

func (m Tools) IsEmpty() bool {
	return reflect.DeepEqual(m, Tools{})
}

func (m Tools) GetAllShow() []Tools {
	var tools []Tools
	global.DB.Where("is_show=?", 1).Order("sort asc").Find(&tools)
	return tools
}
