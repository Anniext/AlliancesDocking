package system

import (
	"AlliancesDocking/config"
	"AlliancesDocking/data"
	services "AlliancesDocking/gin"
	"AlliancesDocking/mqtt"
	"fmt"
	"log"
	"net/http"
)

func SysInit() {
	CacheInit()
	mqtt.MqttxServerInit()
	TcpInit()
	mqtt.Subscribe()
}

func CacheInit() {
	services.CacheManager.GetRoomData()
}

func TcpInit() {
	for _, ip := range data.CacheRoom.GetIP() {
		go GetIPData(ip)
	}
}

func GetIPData(ip string) {
	res, err := http.Get("http://" + config.GetConfig().Dev.Router.Host + "/api/tcp/Create?ip=" + ip)
	if err != nil {
		config.GetLog().Error.Println("Failed to connect to Tcp:", err)
		return
	}
	config.GetLog().Info.Println("GetIPData Cache Sucessful!")
	if res.StatusCode == 200 {
		go ListenHeart(ip)
	}
}

func ListenHeart(ip string) {
	for {
		conn := *data.CacheTcp.Get(ip).NetCoon
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		response := string(buffer[:n])
		if err != nil || response != "KEEPALIVESunmnet" {
			//config.GetLog().Error.Println("device is down", err)
			log.Println("device is down, ip is ", err, ip)
			config.GetLog().Error.Println("device is down, ip is ", err, ip)
			return
		} else {
			fmt.Println("device is alive")
			mqtt.RequestPing(ip, mqtt.MqttClient, mqtt.Topic_ping)
		}
	}
}
