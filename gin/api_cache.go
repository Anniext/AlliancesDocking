package services

import (
	"AlliancesDocking/data"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *RouterGroup) LoadRoomData(c *gin.Context) {
	_, err := data.LoadRoomData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, err)
}
