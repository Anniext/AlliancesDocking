package config

import (
	"log"
	"os"
)

type JsonLogBody struct {
	Info    *log.Logger
	Error   *log.Logger
	Warning *log.Logger
}

func NewLog() string {
	return "./log/log.log"
}

// 写入配置的方法
func (c *OptionConfig) Write() {
	file, _ := os.Create(c.Option) //需要运行读取配置文件函数之后
	jsonLogBody.Info = log.New(file, "[Info] ", log.Ldate|log.Ltime|log.Lshortfile)
	jsonLogBody.Error = log.New(file, "[Error] ", log.Ldate|log.Ltime|log.Lshortfile)
	jsonLogBody.Warning = log.New(file, "[Warning] ", log.Ldate|log.Ltime|log.Lshortfile)
	log.Println("Log file created successfully, program starts running...")
	log.Println("<----------------------------->")
}

// Detection 检测日志方法文件大小,超过10M则重命名,创建新的日志文件
func (c *OptionConfig) Detection() {
	maxSize := 1024 * 1024 * 10 // 10M
	fileInfo, err := os.Stat(c.Option)
	if err != nil {
		jsonLogBody.Error.Println("Failed to get log file information:", err)
	}
	if fileInfo.Size() > int64(maxSize) {
		// 重命名日志文件
		err = os.Rename(c.Option, c.Option+".bak")
		if err != nil {
			jsonLogBody.Error.Println("Failed to rename log file:", err)
		}
		// 创建新的日志文件
		_, err := os.Create(c.Option)
		jsonLogBody.Info.Println("Log file too large, recreate new log file")
		if err != nil {
			jsonLogBody.Error.Println("Failed to create new log file:", err)
		}
	} else {
		log.Println("Log file size is normal, no need to recreate")
		jsonLogBody.Info.Println("Log file size is normal, no need to recreate")
	}
}
