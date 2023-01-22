package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	log.Println("Launching server...")
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("Could not start listener", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Could not accept", err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	buffer := bufio.NewReader(conn)
	message, err := buffer.ReadString('\n')

	for {

		if err != nil {
			log.Panic("Not ending in a delim: ", err)
		}
		log.Println("Message Received:|", message, "|")
		conn.Write([]byte(message))
	}
}
