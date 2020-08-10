package main

import (
	"blog/control"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

//全局配置
type config struct {
	Env   string //当前环境
	DBDsn string //数据库DSN
}

var Conf config

var DB *gorm.DB

func init() {
	initConf()
	initDB()
}

//初始化配置信息
func initConf() {
	viper.SetConfigFile("./config/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(err)
	}
	log.Println("加载全局配置成功")
}

//初始化数据库
func initDB() {
	var err error
	DB, err = gorm.Open("mysql", Conf.DBDsn)
	if err != nil {
		panic(err)
	}
	defer DB.Close()
	err = DB.DB().Ping()
	if err != nil {
		panic(err)
	}
	log.Println("连接数据库成功")
}

func main() {
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
