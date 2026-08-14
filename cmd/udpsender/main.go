package main

import (
	"fmt"
	"net"
	"io"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("Error resolving udp address")
	}

	conn, err := net.DialUDP("udp", laddr, nil)
	if err != nil {
		fmt.Println("Error establishing UDP connection")
	}
}
