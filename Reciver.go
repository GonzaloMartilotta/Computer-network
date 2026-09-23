package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "127.0.0.1:65432")
	if err != nil {
		// Hacer algo
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// Hacer algo
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		// Hacer algo
	}
	fmt.Println(string(buffer[:n]))
	defer conn.Close()
}
