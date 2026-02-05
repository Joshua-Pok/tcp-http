package request

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine  RequestLine
	Parser_state int
}

type RequestLine struct {
	HttpVersion   string // "http1.1"
	RequestTarget string // "/cats"
	Method        string //GET, POST etc
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	buf := make([]byte, 4096)
	n, err := reader.Read(buf) //reader.Read fills a byte size that we provide
	if err != nil {
		fmt.Println("Error reading request")
		return nil, err
	}

	req_string := string(req) //convert []bytes into string

	idx := strings.Index(req_string, "\r\n")
	if idx == -1 {
		return 0, nil
	}

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

func (r *Request) parse(data []byte) (int, error) {

	//accepts the next slice of bytes that need to be pparsed in the request struct
	// updates state of the parser
	// returns number of bytes it consumed and error if it encounteered one
}
