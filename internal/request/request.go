package request

import (
	"io"
	"log"
	"strings"
)

type Request struct {
    RequestLine RequestLine
}

type RequestLine struct {
    HttpVersion   string
    RequestTarget string
    Method        string
}

func requestFromReader(reader io.Reader) (*Request, error) {
	// read entire request as a single string
	request, err := io.ReadAll(reader)
	if err != nil {
		log.Fatal(err)
	}

}

func parseRequestLine(request string) (*string) {
	requestLine := strings.Split(request, "\r\n")
	if len(requestLine) > 0 {
		return &requestLine[0]
	}
	err := "Parsing request line could not be done"
	return &err
}
