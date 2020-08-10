//博客首页相关
package control

import (
	"blog/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func Index(c *gin.Context) {
	c.String(http.StatusOK, "hello index")
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
	fmt.Println(id)
	var art model.Article
	art = art.GetById(id)
	c.JSON(http.StatusOK, art)
}
