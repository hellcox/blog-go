package model

import (
	"blog/global"
	"github.com/jinzhu/gorm"
	"reflect"
	"time"
)

type Article struct {
	ID          int `gorm:"primary_key"`
	Title       string
	Uuid        string
	Description string
	Keyword     string
	Cover       string
	Content     string
	MdContent   string
	Views       int
	Status      int `gorm:"-"`
	CreateAt    time.Time
	UpdateAt    time.Time
}

func (a Article) IsEmpty() bool {
	return reflect.DeepEqual(a, Article{})
}

func (Article) TableName() string {
	return "bg_article"
}

func (a Article) GetById(id int) Article {
	global.DB.First(&a, id)
	return a
}

func (a Article) GetListByPage(page int, limit int) (count int, arts []Article) {
	var articles []Article
	offset := (page - 1) * limit
	global.DB.Limit(limit).Offset(offset).Where("status = ?", 1).Order("id desc").Find(&articles)
	global.DB.Model(Article{}).Where("status = ?", 1).Count(&count)
	return count, articles
}

func (a Article) IncrViews(id int) {
	global.DB.Model(&a).Where("id=?", id).Update("views", gorm.Expr("views + ?", 1))
}
