package data

import (
	"AlliancesDocking/config"
	"log"
	"net/http"
)

var (
	CacheRoom *RoomMap
)

func SystemDataInit() {
	CacheRoom = NewRoomMap()

	_, err := LoadRoomData()
	if err != nil {
		return
	}

	err = GetRelaUnitFileData()
	if err != nil {
		return
	}
}
func GetRelaUnitFileData() (err error) {
	_, err = http.Get("http://127.0.0.1" + config.Configs.Dev.AppPort + "/api/cache/LoadRelaUnitFileData")
	if err != nil {
		log.Println("file_data请求失败:", err)
	}
	return
}
