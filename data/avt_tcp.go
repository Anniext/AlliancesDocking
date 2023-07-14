package data

import "net"

type TcpNet struct {
	NetCoon *net.Conn
	IP      string
}
