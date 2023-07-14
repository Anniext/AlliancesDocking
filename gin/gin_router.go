package services

func (r *RouterGroup) SetCache() {
	r.GET("/LoadRoomData", CacheManager.LoadRoomData)
}

func (r *RouterGroup) SetTcp() {
	r.GET("/Create", TcpManager.Create)
}
