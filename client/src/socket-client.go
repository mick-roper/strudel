package main

import (
	"net"
)

type SocketClient struct {
	socket net.Conn
}
