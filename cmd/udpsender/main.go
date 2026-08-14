package main

import (
	"log"
	"net"
	"io"
)

func main() {
	laddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		log.Println("Error resolving udp address")
	}

	conn, err := net.DialUDP("udp", laddr, nil)
	if err != nil {
		log.Println("Error establishing UDP connection")
	}
}
