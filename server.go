package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	fmt.Println("Launching server...")
	listener, _ := net.Listen("tcp", ":8081")
	for {
		conn, err := listener.Accept()
		if err != nil {
			break
		}
		buffer := bufio.NewReader(conn)
		for {
			message, err := buffer.ReadString('\n')
			if err != nil {
				break
			}
			fmt.Print("Message Received:|", message, "|")
			conn.Write([]byte(message))
		}
	}
}
