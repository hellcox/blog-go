package main

import (
	"blog/control"
	"blog/global"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

func init() {
	global.Init()
}

func main() {
	defer global.DB.Close()
	//环境设置 线上环境为gin.ReleaseMode
	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	//静态资源
	r.LoadHTMLGlob("view/**/*")
	r.StaticFS("/asset", http.Dir("asset"))
	r.StaticFS("/upload", http.Dir("upload"))
	r.StaticFile("/favicon.ico", "./asset/favicon.ico")

	//中间件
	r.Use()

	//注册路由
	router(r)

	//启动服务&打印日志
	log.Fatal(r.Run(":80"))
}

func router(r *gin.Engine) {
	//首页
	r.GET("/", control.Index)     //首页
	r.GET("/p/:id", control.Json) //首页页码

	//博文
	r.GET("/d/:id", control.ArtDetail) //博文详情

	//工具栏
	r.GET("/tool", control.Json)       //工具首页
	r.GET("/tool/:name", control.Json) //工具详情页
}
