package data

import (
	"AlliancesDocking/model"
	"AlliancesDocking/query"
	"fmt"
	"sync"
)

type RoomStatus struct {
	RoomEnvironmentData
	model.AvtRoom
	Id             int64
	Name           string
	FloorId        int64
	BuildingId     int64
	Status         int64 // 小于0 故障，大于0 ： 0-关，1-开
	InUse          int64
	LastPingTime   int64 // 上一次ping的时间戳
	CoreTime       int64
	CardUpdateTime int64
	IP             string
	K4Version      string
	K4MD5          string
	BinMD5         string
	LastBinMD5     string
	LastK4MD5      string
	System         string
	RepairStatus   int64               // 0- 未报修 ； 1-已报修
	LiveStatus     int64               // 0- 未直播 ； 1-直播中
	Equipments     []*model.AvtJoinnum // 如果删除设备，或者改变设备所在教室，那么先要在原来的教室里面移除设备，然后再在新的教室里面添加设备
}

type RoomEnvironmentData struct {
	// 根据设备类型 18-温度；19-湿度；20-pm2.5;21-照度；22-能耗;23-二氧化碳；24-噪音；25-甲醛；26-pm10；27-电流；28-电压
	// 以上设备类型，在每个教室只能出现一次，根据设备类型将joinum 的值，与对应的字段绑定
	Temperature  int `orm:"-" json:"temperature"`  //温度 *10
	Humidity     int `orm:"-" json:"humidity"`     // 湿度
	PM25         int `orm:"-" json:"pm25"`         // PM2.5
	Illumination int `orm:"-" json:"illumination"` // 照度
	Energy       int `orm:"-" json:"energy"`       // 能耗

	HCHO    int `orm:"-" json:"hcho"`    // 甲醛
	Noise   int `orm:"-" json:"noise"`   // 噪音
	PM10    int `orm:"-" json:"pm10"`    // PM10
	CO2     int `orm:"-" json:"co2"`     // 二氧化碳
	Current int `orm:"-" json:"current"` // 电流
	Voltage int `orm:"-" json:"voltage"` // 电压
}

type RoomMap struct {
	roomIP map[string]int64
	data   map[int64]*RoomStatus
	lock   sync.RWMutex
}

func NewRoomMap() *RoomMap {
	return &RoomMap{
		roomIP: make(map[string]int64),
		data:   make(map[int64]*RoomStatus),
	}
}

func LoadRoomData() (count int64, err error) {
	var roomList []*model.AvtRoom
	err = query.AvtRoom.Where(query.AvtRoom.IsDelete.Eq(0)).Scan(&roomList)
	count, err = query.AvtRoom.Where(query.AvtRoom.IsDelete.Eq(0)).Count()
	if err == nil {
		if count > 0 {
			for _, room := range roomList {
				rs := &RoomStatus{
					Id:         room.ID,
					Name:       *room.Name,
					IP:         *room.IP,
					FloorId:    *room.FloorID,
					BuildingId: *room.BuildingID,
					Status:     0,
					Equipments: make([]*model.AvtJoinnum, 0),
				}
				list, err := GetAvtEquipmentsByRoomId(room.ID)
				if err == nil {
					rs.Equipments = list
				}
				CacheRoom.Set(rs)
			}
			LoadWorkerOrder()
		} else {
			fmt.Println("系统没有教室数据！")
		}

	} else {
		fmt.Println("教室桩数据加载错误！")
	}
	return count, err
}
