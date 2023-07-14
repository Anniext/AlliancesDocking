package data

var (
	CacheRoom *RoomMap
	CacheTcp  *TcpMap
)

func SystemDataInit() {
	CacheRoom = NewRoomMap()
	CacheTcp = NewTcpMap()
}
