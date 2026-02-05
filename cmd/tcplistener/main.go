package main

import (
	"fmt"
	"io"
	"net"
	"strings"
)

// create a string variable to hold contents of current line
// after reading 8 bytes, we split the data on new lines
// for each part except the last part, print line : read: current line + first part
// add the last part to the current line variable
// after the file ends we print whatever is left on current line

func getLinesChannel(f io.ReadCloser) <-chan string {
	//creates a channel of strings

	lines := make(chan string)
	//does reading inside a go routine

	go func() {

		defer func() {
			f.Close()
			defer close(lines)
		}()

		curr_line := ""
		buf := make([]byte, 8)

		for {
			n, err := f.Read(buf)
			if n > 0 {
				chunk := string(buf[:n])
				parts := strings.Split(chunk, "\n")

				for i := 0; i < len(parts)-1; i++ {
					curr_line += parts[i]
					fmt.Println(curr_line)
					lines <- curr_line //sends line to the channel
					curr_line = ""

				}
				curr_line += parts[len(parts)-1]
			}

			if err == io.EOF {
				if curr_line != "" {
					lines <- curr_line //insert any remaining strings into channel
				}
				return
			}

			if err != nil {
				return
			}

		}

	}()

	return lines

}

func main() {

	port := ":42069"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println("Unable to listen to port")
		return
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection: ", err)
			continue
		}

		fmt.Println("Connection accepted from:", conn)

		go func(c net.Conn) {
			for line := range getLinesChannel(c) {
				fmt.Println(line)
			}
		}(conn)
	}

}
