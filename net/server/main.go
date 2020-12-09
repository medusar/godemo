package main

import (
	"bufio"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:5555")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println("error accept", err)
			continue
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	rd := bufio.NewReader(conn)
	for {
		line, _, err := rd.ReadLine()
		conn.Write(line)
		if err != nil {
			log.Println("failed to Read", err)
			break
		}
	}
}
