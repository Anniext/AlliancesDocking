package config

import (
	"encoding/json"
	"log"
	"os"
)

type JsonConfBody struct {
	Dev struct {
		AppName string `json:"app_name"`
		Dsn     string `json:"dsn"`
		Mqtt    struct {
			Host     string `json:"host"`
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"mqtt"`
		Router struct {
			Host string `json:"host"`
			Mode string `json:"mode"`
		} `json:"router"`
	} `json:"dev"`
}

func NewConfig() string {
	return "./cnf/config.json"
}

// 读取文件的方法
func (c *OptionConfig) Read() {
	// 读取json文件
	file, err := os.ReadFile(c.Option)
	if err != nil {
		log.Println("Failed to read configuration file:", err)
	} else {
		log.Println("Configuration file loaded successfully...")
	}
	// 解析json文件,序列化为对象
	err = json.Unmarshal(file, &jsonConfBody)
	if err != nil {
		log.Println("Failed to serialize json object:", err)
	} else {
		log.Println("serialized json object successfully...")
		log.Println("<----------------------------->")
	}
}
