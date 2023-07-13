package data

func (m *RoomMap) UpdateRoomRepairStatus(rid int64, status int64) {
	m.lock.Lock()
	_, ok := m.data[rid]
	if ok {
		m.data[rid].LiveStatus = status
	}
	m.lock.Unlock()
}

func (m *RoomMap) Set(room *RoomStatus) {
	m.lock.Lock()
	m.data[room.ID] = room
	m.roomIP[room.IP] = room.ID
	m.lock.Unlock()
}

func (m *RoomMap) GetIP() []string {
	m.lock.RLock()
	defer m.lock.RUnlock()
	return func() []string {
		var ips []string
		for k := range m.roomIP {
			ips = append(ips, k)
		}
		return ips
	}()
}
