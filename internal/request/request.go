package request

import (
	"errors"
	"io"
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

func RequestFromReader(reader io.Reader) (*Request, error) {
	request, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	requestLine, err := parseRequestLine(string(request))
	if err != nil {
		return nil, err
	}

	return &Request{
		RequestLine: *requestLine,
	}, nil
}

func parseRequestLine(line string) (*RequestLine, error) {
	parts := strings.Split(line, "\r\n")
	if len(parts) == 0 {
		return nil, errors.New("improper request format")
	}

	fields := strings.Fields(parts[0])
	if len(fields) != 3 {
		return nil, errors.New("improper request line format")
	}

	httpVersion := strings.TrimPrefix(fields[2], "HTTP/")

	return &RequestLine{
		Method:        fields[0],
		RequestTarget: fields[1],
		HttpVersion:   httpVersion,
	}, nil
}
