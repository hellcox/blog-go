package model

import "blog/global"

type Article struct {
	Id          int    `json:"id" gorm:"primary_key"`
	Title       string `json:"title"`
	Uuid        string `json:"uuid"`
	Description string `json:"description"`
	Keyword     string `json:"keyword"`
	Cover       string `json:"cover"`
	Content     string `json:"content"`
	MdContent   string `json:"md_content"`
	GmtCreate   string `json:"gmt_create"`
	GmtUpdate   string `json:"gmt_update"`
	Views       int    `json:"views"`
	Status      int    `json:"status"`
}

func (a Article) GetById(id int) Article {
	global.DB.Table("bg_article").First(&a, id)
	return a
}
