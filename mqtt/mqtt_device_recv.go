package mqtt

import (
	"AlliancesDocking/config"
	"AlliancesDocking/data"
	"github.com/bytedance/sonic"
	"log"
	"strconv"
	"strings"
)

type DevicePushRequest struct {
	Event string                `json:"event"`
	Data  DevicePushRequestData `json:"data"`
}

type DevicePushRequestData struct {
	Code  int                   `json:"code"`
	Body  DevicePushRequestBody `json:"body"`
	MsgId int                   `json:"msg_id"`
}

type DevicePushRequestBody struct {
	Devices []*DeviceValue `json:"devices"`
}

type DeviceValue struct {
	Id        int    `json:"id"`
	ValueType int    `json:"type"`
	Value     string `json:"value"`
}

func DevicePush(c *Client, topic string, msg []byte) {
	var req DevicePushRequest
	err := sonic.Unmarshal(msg, &req)
	ip := getCoreID(topic)
	if err != nil {
		log.Println("DevicePush Unmarshal Error :", err.Error())
	} else {
		if req.Event == "device_recv" {
			for _, v := range req.Data.Body.Devices {
				// 控制设备
				go RemoteConnect(ip, v)
				config.GetLog().Info.Println("DevicePush", v)
			}
			cnn := *data.CacheTcp.Get(ip).NetCoon
			buffer := make([]byte, 1024)
			n, err := cnn.Read(buffer)
			response := string(buffer[:n])
			if err != nil {
				log.Println("DevicePush Read Error :", err.Error())
				return
			}
			if response == "CONTROL_DEVICE  receive !Sunmnet" {
				rp := DevicePushRequest{
					Event: "device_push",
					Data: DevicePushRequestData{
						Code: 0,
						Body: DevicePushRequestBody{
							Devices: req.Data.Body.Devices,
						},
						MsgId: 0,
					},
				}
				jsondata, err := sonic.Marshal(&rp)
				if err != nil {
					log.Println("DevicePush Marshal Error :", err.Error())
					return
				}
				topic := strings.Replace(topic_device_push, "+", ip, -1)
				err = c.Publish(topic, 0, false, jsondata)
				if err != nil {
					config.GetLog().Error.Println("mqtt publish error!!!")
					return
				}
				// 更新设备到前端

			}

		}
	}
}

func RemoteConnect(ip string, deviceData *DeviceValue) {
	cnn := *data.CacheTcp.Get(ip).NetCoon
	joinnum := deviceData.Id
	switch {
	case joinnum == 1: // 上下课
		if deviceData.Value == "0" {
			_, err := cnn.Write([]byte("CONTROL_DEVICE3Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		} else if deviceData.Value == "1" {
			_, err := cnn.Write([]byte("CONTROL_DEVICE2Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		}
	case joinnum == 2: // 功放
		if deviceData.Value == "0" {
			_, err := cnn.Write([]byte("CONTROL_DEVICESOUND,0;Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		} else if deviceData.Value == "1" {
			_, err := cnn.Write([]byte("CONTROL_DEVICESOUND,1;Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		}
	case joinnum >= 3 && joinnum <= 7: // 投影仪
		if deviceData.Value == "0" {
			_, err := cnn.Write([]byte("CONTROL_DEVICEPROJECTOR,0,PROJECTOR_000" + strconv.Itoa(joinnum-2) + ";Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		} else if deviceData.Value == "1" {
			_, err := cnn.Write([]byte("CONTROL_DEVICEPROJECTOR,1,PROJECTOR_000" + strconv.Itoa(joinnum-2) + ";Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		}
	case joinnum >= 8 && joinnum <= 12: // 空调
		if deviceData.Value == "0" {
			_, err := cnn.Write([]byte("CONTROL_DEVICEAIRCONDITIONER,1.1.25,AIRCONDITIONER_00" + strconv.Itoa(joinnum-7) + ";Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		} else if deviceData.Value == "1" {
			_, err := cnn.Write([]byte("CONTROL_DEVICEAIRCONDITIONER,1.1.25,AIRCONDITIONER_00" + strconv.Itoa(joinnum-7) + ";Sunmnet"))
			if err != nil {
				log.Println("RemoteConnect Write Error :", err.Error())
				return
			}
		}
	}
}
