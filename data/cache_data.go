package data

var (
	CacheRoom *RoomMap
)

func SystemDataInit() {
	CacheRoom = NewRoomMap()

	_, err := LoadRoomData()
	if err != nil {
		return
	}

}
