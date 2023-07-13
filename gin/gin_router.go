package services

import "github.com/gin-gonic/gin"

type RouterGroup struct {
	*gin.RouterGroup
}

func (r *RouterGroup) SetCache() {
	r.GET("/LoadRoomData", CacheManager.LoadRoomData)
}
