package http

import "net"

func init() {
	listener, _ := net.Listen("tcp", ":8080")

	for true {
		con, err := listener.Accept()
	}
}
