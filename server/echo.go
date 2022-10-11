package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:56789")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Print(err)
	}
	time.Sleep(2 * time.Second)
	m, err := conn.Write(buffer[:n])
	if err != nil {
		log.Print(err)
	}
	if n != m {
		log.Print(fmt.Errorf("read-write unequal length"))
	}
}
