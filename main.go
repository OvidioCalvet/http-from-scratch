package main

import (
	"fmt"
	"os"
)

func readFromFile() {
	file, err := os.Open("messages.txt")
	if err != nil {
		fmt.Println("Failed to open messages.txt")
	}

	for {
		data := make([]byte, 8)
		count, err := file.Read(data)
		if err != nil {
			break;
		}

		fmt.Printf("read: %s\n", string(data[:count]))
	}

}

func main() {
	readFromFile()
}
