//博客首页相关
package control

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func Index(c *gin.Context){
	c.String(http.StatusOK,"hello index")
}

func Json(c *gin.Context) {
	c.JSON(200, gin.H{
		"a": "hello world",
	})
}

func ArtDetail(c *gin.Context){
	artId := c.Param("id")
	fmt.Println(artId)
	c.String(http.StatusOK,artId)
}


