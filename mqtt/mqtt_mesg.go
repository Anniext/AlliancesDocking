package mqtt

import (
	"fmt"
	"strings"
)

func getCoreID(topic string) string {
	tps := strings.Split(topic, "/")
	if len(tps) > 2 {
		return tps[1]
	}
	return ""
}

func Subscribe() {
	err := MqttClient.Subscribe(HandleMessage, 0, topic_device_recv)
	if err != nil {
		fmt.Println("Subscribe error:", err)
		return
	}
}

// HandleMessage 消息处理回调函数
func HandleMessage(c *Client, topic string, msg []byte) {
	DevicePush(c, topic, msg)
}
