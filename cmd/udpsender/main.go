package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {

	port := "localhost:42069"
	addr, err := net.ResolveUDPAddr("udp", port)
	if err != nil {
		fmt.Println("Unable to resolve UDP address")
		return
	}

	conn, err := net.DialUDP("udp", nil, addr) // ladder: local address isuaually nil, raddr is the address we want to send to
	if err != nil {
		fmt.Println("Unable to establish UDP connection")
		return
	}

	defer conn.Close()

	r := bufio.NewReader(os.Stdin)

	for {

		fmt.Println(">")

		line, err := r.ReadString('\n') // single quotes: rune/bytes, double quotes: string
		if err != nil {
			fmt.Println(err)
			break
		}

		conn.Write([]byte(line)) //convert line into raw bytes

	}

}
