package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"time"
)

var (
	message = flag.String("message", "Dies ist ein Test", "echo-Nachricht")
)

func main() {
	flag.Parse()

	conn, err := net.DialTimeout("tcp", "localhost:56789", 3*time.Second)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	const maxBufferSize = 1024
	l := len(*message)
	if l > maxBufferSize {
		l = maxBufferSize
	}
	n, err := conn.Write([]byte(*message)[:l])
	if err != nil {
		log.Fatalf("failed to write to server %w", err)
	}
	ch := make(chan string)
	go func() {
		recv := make([]byte, n)
		_, err = conn.Read(recv)
		if err != nil {
			log.Printf("error receiving form server %w", err)
		}
		ch <- string(recv)
	}()
	timeout := time.After(time.Second)
	select {
	case s := <-ch:
		fmt.Printf("Received: %s\n", s)

	case <-timeout:
		log.Print("timeout receiving echo")
	}
}
