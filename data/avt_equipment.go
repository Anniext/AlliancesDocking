package data

import (
	"AlliancesDocking/config"
	"AlliancesDocking/model"
	"AlliancesDocking/query"
)

func GetAvtEquipmentsByRoomId(rid int64) (ml []*model.AvtJoinnum, err error) {
	err = query.AvtJoinnum.Where(query.AvtJoinnum.RoomID.Eq(rid), query.AvtJoinnum.ShowStatus.Eq(1)).Order(query.AvtJoinnum.Sort.Desc()).Scan(&ml)
	if err != nil {
		return nil, err
	}
	config.GetLog().Info.Println("GetAvtEquipmentsByRoomId", ml)
	return
}
