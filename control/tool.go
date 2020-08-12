//实用工具
package control

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func ToolTmp(c *gin.Context) {
	c.HTML(http.StatusOK, "toolIndex.html", gin.H{})
}

func ToolMd5Tmp(c *gin.Context) {
	c.HTML(http.StatusOK, "toolMd5.html", gin.H{})
}

func ToolTimeTmp(c *gin.Context) {
	now := time.Now()
	c.HTML(http.StatusOK, "toolTime.html", gin.H{
		"timestamp": now.Unix(),
		"datetime":  now.Format("2006-01-02 15:04:05"),
	})
}
