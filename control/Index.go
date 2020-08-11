//博客首页相关
package control

import (
	"blog/model"
	"blog/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		page = 1
	}
	limit := 10
	count, arts := model.Article{}.GetListByPage(page, limit)
	fmt.Println(count)
	fmt.Println(len(arts))
	c.HTML(http.StatusOK, "index.html", gin.H{
		"total": count,
		"list":  arts,
		"page":  template.HTML(util.Page(page, limit, count)),
	})
}

func Json(c *gin.Context) {
	c.JSON(200, gin.H{
		"a": "hello world",
	})
}

func ArtDetail(c *gin.Context) {
	artId := c.Param("id")
	id, err := strconv.Atoi(artId)
	if err != nil {
		c.JSON(200, gin.H{
			"a": "hello world",
		})
	}
	var art model.Article
	art = art.GetById(id)
	fmt.Println(art.IsEmpty())
	c.JSON(http.StatusOK, gin.H{
		"msg": "",
		"row": art,
	})
}
