package data

import (
	"AlliancesDocking/config"
	"AlliancesDocking/model"
	"AlliancesDocking/query"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var CacheRoom *RoomMap

func init() {
	query.SetDefault(func() *gorm.DB {
		db, err := gorm.Open(mysql.Open(config.Configs.Dev.Dsn), &gorm.Config{})
		if err != nil {
			log.Println("Failed to connect to database:", err)
		}
		return db
	}())
}

func InitCacheData() {
	CacheRoom = NewRoomMap()

	_, err := LoadRoomData()
	if err != nil {
		return
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
	fmt.Println(CacheRoom)
	return count, err
}

// LoadWorkerOrder 加载工单数据，同步教室报修状态
func LoadWorkerOrder() {
	var list []*AvtWorkOrder
	err := query.AvtWorkOrder.Where(query.AvtWorkOrder.Status.Eq(0)).Or(query.AvtWorkOrder.Status.Eq(1)).Scan(&list)
	if err == nil {
		for _, wk := range list {
			CacheRoom.UpdateRoomRepairStatus(*wk.RoomID, 1)
		}
	}
}
