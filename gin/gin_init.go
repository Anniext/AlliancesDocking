package services

import (
	"AlliancesDocking/config"
	"AlliancesDocking/query"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type RouterGroup struct {
	*gin.RouterGroup
}

var (
	router       *gin.Engine
	api          *RouterGroup
	CacheManager CacheInterface
	TcpManager   TcpInterface
)

type CacheInterface interface {
	SetCache()
	LoadRoomData(c *gin.Context)
	GetRoomData()
}

type TcpInterface interface {
	SetTcp()
	Create(c *gin.Context)
	//Authentication(c *gin.Context)
}

func init() {
	RouterGroupInit()
}

func RouterGroupInit() {
	router = gin.Default()
	gin.SetMode(config.GetConfig().Dev.Router.Mode)
	//router.Use(DevCors()) //使用自定义的跨域中间件
	api = &RouterGroup{router.Group("/api")}
	CacheManager = &RouterGroup{api.Group("/cache")}
	TcpManager = &RouterGroup{api.Group("/tcp")}
}

func Serviceinit() {
	query.SetDefault(func() *gorm.DB {
		db, err := gorm.Open(mysql.Open(config.GetConfig().Dev.Dsn), &gorm.Config{})
		if err != nil {
			config.GetLog().Error.Println("Failed to connect to database:", err)
		}
		return db
	}()) //初始化数据库
	GroupInit()          // 初始化路由
	go StartMainServer() // 启动主服务
}

func GroupInit() {
	CacheManager.SetCache()
	TcpManager.SetTcp()
}

func StartMainServer() {
	err := http.ListenAndServe(config.GetConfig().Dev.Router.Host, router)
	config.GetLog().Info.Println(config.GetConfig().Dev.Router.Host)
	if err != nil {
		log.Println("Failed to start server: ", err)
		return
	} else {
		log.Println("Server started on port: ", config.GetConfig().Dev.Router.Host)
	}
}
