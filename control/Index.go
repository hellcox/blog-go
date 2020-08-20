//博客首页相关
package control

import (
	"blog/global"
	"blog/model"
	"blog/util"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"strconv"
)

//首页
func IndexTmp(c *gin.Context) {
	page, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		page = 1
	}
	limit := 10
	count, arts := model.Article{}.GetListByPage(page, limit)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"total": count,
		"list":  arts,
		"page":  template.HTML(util.Page(page, limit, count)),
		"tkd":   global.TKD.Home,
	})
}

//博客详情
func ArtDetailTmp(c *gin.Context) {
	emsg := ""
	artId := c.Param("id")
	id, err := strconv.Atoi(artId)
	if err != nil {
		c.HTML(http.StatusOK, "detail.html", gin.H{
			"emsg": "(ง •̀_•́)ง 文章地址可能错啦！",
		})
		return
	}
	var art model.Article
	art = art.GetById(id)
	if !art.IsEmpty() {
		art.IncrViews(id)
	} else {
		emsg = "(ง •̀_•́)ง 页面找不到啦！"
	}
	c.HTML(http.StatusOK, "detail.html", gin.H{
		"emsg":    emsg,
		"data":    art,
		"content": template.HTML(art.Content),
		"tkd":     global.Tkd{T: art.Title + "-小豆豆博客", K: art.Keyword, D: art.Description},
	})
}
