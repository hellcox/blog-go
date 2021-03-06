//实用工具
package control

import (
	"blog/global"
	"blog/model"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func ToolTmp(c *gin.Context) {
	tools := model.Tools{}.GetAllShow()
	c.HTML(http.StatusOK, "toolIndex.html", gin.H{
		"tkd":   global.TKD.Tools,
		"tools": tools,
	})
}

func ToolMd5Tmp(c *gin.Context) {
	c.HTML(http.StatusOK, "toolMd5.html", gin.H{"tkd": global.TKD.ToolMd5,})
}

func ToolBase64Tmp(c *gin.Context) {
	c.HTML(http.StatusOK, "toolBase64.html", gin.H{"tkd": global.TKD.ToolBase64,})
}

func ToolUrlEncodeTmp(c *gin.Context) {
	c.HTML(http.StatusOK, "toolUrlEncode.html", gin.H{"tkd": global.TKD.ToolUrlEncode,})
}

func ToolIpTmp(c *gin.Context) {
	ip := c.ClientIP()
	rip := c.GetHeader("Remote_addr")
	if rip != "" {
		ip = rip
	}
	if ip == "::1" {
		ip = "127.0.0.1"
	}
	c.HTML(http.StatusOK, "toolIp.html", gin.H{
		"ip":  ip,
		"tkd": global.TKD.ToolIp,
	})
}

func ToolTimeTmp(c *gin.Context) {
	now := time.Now()
	c.HTML(http.StatusOK, "toolTime.html", gin.H{
		"timestamp": now.Unix(),
		"datetime":  now.Format("2006-01-02 15:04:05"),
		"tkd":       global.TKD.ToolTime,
	})
}

func ToolColorTmp(c *gin.Context) {
	c.HTML(http.StatusOK, "toolColor.html", gin.H{
		"tkd": global.TKD.ToolColor,
	})
}

//加密操作
func Encode(c *gin.Context) {
	doType, _ := c.GetPostForm("type")
	value, _ := c.GetPostForm("value")
	fmt.Println(doType)
	fmt.Println(reflect.TypeOf(value), value)
	switch doType {
	case "md5":
		h := md5.New()
		h.Write([]byte(value))
		lower32 := hex.EncodeToString(h.Sum(nil))
		lower16 := lower32[8:24]

		//加密值入库
		m := model.Md5{}
		if m.GetByM32(lower32).IsEmpty() {
			m.M32 = lower32
			m.M16 = lower16
			m.Content = value
			fmt.Println(m.Save())
		}

		type Values struct {
			Upper32 string
			Lower32 string
			Upper16 string
			Lower16 string
		}
		values := Values{
			Upper32: strings.ToUpper(lower32),
			Lower32: lower32,
			Upper16: strings.ToUpper(lower16),
			Lower16: lower16,
		}
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"data": values,
		})
	case "base64":
	default:
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "加密类型错误",
		})
	}
}

//解密操作
func Decode(c *gin.Context) {
	doType, _ := c.GetPostForm("type")
	value, _ := c.GetPostForm("value")
	switch doType {
	case "md5":
		m := model.Md5{}
		if len(value) == 16 {
			m = m.GetByM16(value)
		} else if len(value) == 32 {
			m = m.GetByM32(value)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "数据格式错误",
			})
			return
		}
		if m.IsEmpty() {
			c.JSON(http.StatusOK, gin.H{
				"code": 1,
				"msg":  "解密失败",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "",
				"data": struct{ Content string }{m.Content},
			})
		}

	default:
		c.JSON(http.StatusOK, gin.H{
			"code": 1,
			"msg":  "解密类型错误",
		})
	}
}
