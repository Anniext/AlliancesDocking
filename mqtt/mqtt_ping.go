package mqtt

import (
	"AlliancesDocking/config"
	"AlliancesDocking/utils"
	"github.com/bytedance/sonic"
	"strings"
	"time"
)

type PingRequest struct {
	Data  PingRequestData
	Event string `json:"event"`
}

type PingRequestData struct {
	Body  PingRequestDataBody
	Code  int `json:"code"`
	MsgId int `json:"msg_id"`
}

type PingRequestDataBody struct {
	BinMd5         string `json:"bin_md5"`
	Date           string `json:"date"`
	Ip             string `json:"ip"`
	Kernel4Md5     string `json:"kernel4_md5"`
	Kernel4Version string `json:"kernel4_version"`
	Mac            string `json:"mac"`
	System         string `json:"system"`
	Time           string `json:"time"`
}

func RequestPing(ip string, c *Client, topic string) {
	res := PingRequest{
		Event: "ping",
		Data: PingRequestData{
			Code:  0,
			MsgId: 0,
			Body: PingRequestDataBody{
				Date:           utils.GetDateFromString(time.Now()), //获取当前日期
				Time:           utils.GetTimeFromString(time.Now()), //获取当前时间
				Ip:             ip,
				BinMd5:         "1234567890",
				Kernel4Md5:     "1234567890",
				Kernel4Version: "Virtual",
				Mac:            "1234567890",
				System:         "Virtual",
			},
		},
	}
	jsondata, err := sonic.Marshal(&res)
	if err == nil {
		topic := strings.Replace(topic, "+", ip, -1)
		err := c.Publish(topic, 0, false, jsondata)
		if err != nil {
			config.GetLog().Error.Println("mqtt publish error!!!")
			return
		} //发布消息
	}
}
