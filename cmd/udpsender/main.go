package main

import (
	"bufio"
	"log"
	"net"
	"os"
	"fmt"
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

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf(">")

		line, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Error reading string from reader")
		}

		conn.Write([]byte(line))
	}
}
