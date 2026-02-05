package request

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string // "http1.1"
	RequestTarget string // "/cats"
	Method        string //GET, POST etc
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		fmt.Println("Error reading request")
		return nil, err
	}

	req_string := string(req) //convert []bytes into string

	parts := strings.Split(req_string, "\r\n")

	req_line := parts[0]

	parseRequestLine(req_line)

	return nil, nil

}

func parseRequestLine(req string) ([]string, error) { //we are only parsing GET /coffee HTTP/1.1

	parts := strings.Split(req, " ") // split by CLRF registered nurse

	if len(parts) < 3 {
		fmt.Println("Invalid request line")
		return nil, nil
	}

	//first part is always the Method
	method := parts[0]

	req_target := parts[1]

	version := parts[2]

	for _, c := range method {
		if !unicode.IsUpper(c) {
			fmt.Println("Invalid Method")
			return nil, nil
		}
	}

	if version != "HTTP/1.1" {
		fmt.Println("Invalid HTTP version")
		return nil, nil
	}

	return []string{method, req_target, version}, nil

}
