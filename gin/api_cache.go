package services

import (
	"AlliancesDocking/config"
	"AlliancesDocking/data"
	errcode "AlliancesDocking/error"
	"AlliancesDocking/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RouterGroup) LoadRoomData(c *gin.Context) {
	_, err := data.LoadRoomData()
	if err != nil {
		c.JSON(http.StatusBadRequest, result.JsonError(errcode.CacheErr))
	}
	c.JSON(http.StatusOK, result.JsonSuccess())
}

func (r *RouterGroup) GetRoomData() {
	_, err := http.Get("http://" + config.GetConfig().Dev.Router.Host + "/api/cache/LoadRoomData")
	if err != nil {
		config.GetLog().Error.Println("Failed to connect to database from avt_room:", err)
	}

}
