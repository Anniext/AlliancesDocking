package main

import (
	"AlliancesDocking/config"
	"AlliancesDocking/data"
	services "AlliancesDocking/gin"
	"AlliancesDocking/system"
	"time"
)

func main() {
	//generate.GenGenerate()
	config.Configinit()
	// 初始化服务器
	services.Serviceinit()
	time.Sleep(200 * time.Millisecond)
	data.SystemDataInit()
	system.SysInit()
	select {}
}
