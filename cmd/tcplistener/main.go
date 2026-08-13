package main

import (
	"fmt"
	"net"
	"strings"
)

func getLinesChannel(c net.Conn) <-chan string {
	// create channel of type string
	ch := make(chan string)

	// go routine with parsing logic
	go func() {
		defer close(ch)
		buffer := ""

		for {
			// create an array of size 8 bytes
			data := make([]byte, 8)
			// read from tcp socket and populate 8 bytes from messages.txt
			count, err := c.Read(data)
			if err != nil {
				break //EOF
			}

			// convert bytes to string and add to buffer
			chunk := string(data[:count])
			buffer += chunk

			// split buffer with every \n found
			slices := strings.Split(buffer, "\n")

			// if a \n was found send line down channel and update buffer
			if len(slices) > 1 {
				for i := range len(slices) - 1 {
					ch <- slices[i]
					buffer = slices[i+1]
				}
			}
		}
	}()

	return ch
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Failed to start tcp listener")
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept tcp connection at desired port")
		} else {
			fmt.Println("Connection accepted!")
		}

		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Printf("read: %s\n", line)
		}

		conn.Close() //close listener once program ends
	}

}
