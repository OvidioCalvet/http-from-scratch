package request

import (
	"errors"
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

	parsedLines := parseRequestLine(request)

}

func parseRequestLine(request string) (*RequestLine, error) {
	parsedLines := strings.Split(request, "\r\n")

	if len(parsedLines) == 3 {
		requestLine := RequestLine {
			HttpVersion: parsedLines[0],
			RequestTarget: parsedLines[1],
			Method: parsedLines[2],
		}

		return &requestLine
	}

	return errors.New("improper request format")
}
