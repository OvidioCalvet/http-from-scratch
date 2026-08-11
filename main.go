package main

import (
	"fmt"
	"os"
	"strings"
)

func readFromFile() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Failed to open messages.txt")
	}

	buffer := ""

	for {
		// create an array of size 8 bytes
		data := make([]byte, 8)
		// read from file and populate 8 bytes from messages.txt
		count, err := file.Read(data)
		if err != nil {
			break //EOF
		}

		// convert bytes to string and add to buffer
		chunk := string(data[:count])
		buffer += chunk

		// split buffer with every \n found
		slices := strings.Split(buffer, "\n")

		// if a \n was found print line and update buffer
		if len(slices) > 1 {
			for i := range len(slices) - 1 {
				fmt.Printf("read: %s\n", slices[i])
				buffer = slices[i + 1]
			}
		}
	}
}

func main() {
	readFromFile()
}
