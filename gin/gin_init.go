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

var (
	router *gin.Engine
	api    *RouterGroup
	//FileManager RouterManager

	CacheManager CacheInterface
)

type CacheInterface interface {
	SetCache()
	LoadRoomData(c *gin.Context)
}

func init() {
	RouterGroupInit()
}

func RouterGroupInit() {
	router = gin.Default()
	router.Use(DevCors()) //使用自定义的跨域中间件
	api = &RouterGroup{router.Group("/api")}
	CacheManager = &RouterGroup{api.Group("/cache")}
}

func Serviceinit() {
	query.SetDefault(func() *gorm.DB {
		db, err := gorm.Open(mysql.Open(config.Configs.Dev.Dsn), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database:", err)
		}
		return db
	}()) //初始化数据库
	GroupInit()          // 初始化路由
	go StartMainServer() // 启动主服务

}

func StartMainServer() {
	err := http.ListenAndServe(config.Configs.Dev.AppPort, router)
	if err != nil {
		log.Println("Failed to start server: ", err)
	} else {
		log.Println("Server started on port: ", config.Configs.Dev.AppPort)
		return
	}
}

func GroupInit() {
	CacheManager.SetCache()
}

//
//func sadfs() {
//	router := gin.Default()
//
//	router.GET("/send-tcp", func(c *gin.Context) {
//		// 建立TCP连接
//		conn, err := net.Dial("tcp", "localhost:8080")
//		if err != nil {
//			c.String(500, "Error connecting to TCP server")
//			return
//		}
//		defer conn.Close()
//
//		// 设置长连接保持活跃的时间
//		conn.SetKeepAlive(true)
//		conn.SetKeepAlivePeriod(30 * time.Second)
//
//		// 发送数据
//		message := "Hello, server!"
//		_, err = conn.Write([]byte(message))
//		if err != nil {
//			c.String(500, "Error sending data")
//			return
//		}
//
//		// 接收响应
//		buffer := make([]byte, 1024)
//		n, err := conn.Read(buffer)
//		if err != nil {
//			c.String(500, "Error receiving response")
//			return
//		}
//
//		// 处理响应
//		response := string(buffer[:n])
//		c.String(200, response)
//	})
//
//	router.Run(":8080")
//}
