package services

import (
	"AlliancesDocking/config"
	"AlliancesDocking/data"
	errcode "AlliancesDocking/error"
	"AlliancesDocking/result"
	"AlliancesDocking/utils"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

func (r *RouterGroup) Create(c *gin.Context) {
	ip := c.Query("ip")
	tcp := &data.TcpNet{
		IP: ip,
		NetCoon: func(c *gin.Context) *net.Conn {
			conn, err := net.DialTimeout("tcp", ip+":"+config.GetConfig().Dev.Target.Port, 5*time.Second)
			if err != nil {
				config.GetLog().Error.Println(err)
				c.JSON(http.StatusInternalServerError, result.JsonError(errcode.TcpServerErr))
				return nil
			}
			err = conn.(*net.TCPConn).SetKeepAlive(true)
			if err != nil {
				config.GetLog().Error.Println(err)
				c.JSON(http.StatusInternalServerError, result.JsonError(errcode.TcpServerErr))
				return nil
			}
			_, err = conn.Write([]byte("GET_CHECKET" + utils.String2Md5("Sunmnet-"+config.GetConfig().Dev.Target.Key+"-"+ip) + "Sunmnet"))
			if err != nil {
				config.GetLog().Error.Println(err)
				return nil
			}
			buffer := make([]byte, 1024)
			n, err := conn.Read(buffer)
			response := string(buffer[:n])
			if err == nil && response == "GET_CHECKETAuthorization success !Sunmnet" {
				return &conn
			}
			return nil
		}(c),
	}
	data.CacheTcp.Set(tcp)
	if data.CacheTcp.Get(ip).NetCoon == nil {
		c.JSON(http.StatusNotFound, result.JsonError(errcode.TcpServerErr))
		return
	}
	c.JSON(http.StatusOK, result.JsonSuccess())
}
