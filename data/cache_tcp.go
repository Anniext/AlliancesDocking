package data

import "sync"

type TcpMap struct {
	data map[string]*TcpNet
	lock sync.RWMutex
}

func NewTcpMap() *TcpMap {
	return &TcpMap{
		data: make(map[string]*TcpNet),
	}
}

func (m *TcpMap) Set(con *TcpNet) {
	m.lock.Lock()
	m.data[con.IP] = con
	m.lock.Unlock()
}

func (m *TcpMap) Get(ip string) *TcpNet {
	m.lock.RLock()
	defer func() { m.lock.RUnlock() }()
	return m.data[ip]
}
