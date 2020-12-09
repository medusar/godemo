package main

import (
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:5555")
	if err != nil {
		log.Fatal(err)
	}

	go loopRead(conn)

	b := []byte("hello\n")
	conn.Write(b)
	conn.Write(b)
	conn.Write(b)
	conn.Write(b)
	conn.Write(b)
	conn.Write(b)

	c := make(chan struct{})
	<-c
}

func loopRead(conn net.Conn) {
	for {
		tmp := make([]byte, 2048)
		n, err := conn.Read(tmp)
		if n > 0 {
			log.Println("data read:", string(tmp[:n]))
		}
		if err != nil {
			log.Fatal("failed to Read", err)
		}
	}
}
