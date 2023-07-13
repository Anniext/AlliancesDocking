package data

import (
	"AlliancesDocking/model"
	"AlliancesDocking/query"
)

type AvtWorkOrder struct {
	model.AvtWorkOrder
	Building string `orm:"-" json:"building"`
	Floor    string `orm:"-" json:"floor"`
	Room     string `orm:"-" json:"room"`
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
