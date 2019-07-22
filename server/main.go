package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal(err, "on server setup")
	}
	// listen indefinitely for clients
	// use another go routine if server must do anything else
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err, "on client connection")
		}
		// conn is a net.Conn instance and we need to keep track of it's writer and reader inside another go routine
		// as net.Conn already implements Reader and Writer interfaces we use a bufio.Reader/Writer
		r := bufio.NewReader(conn)
		// wr := bufio.NewWriter(conn)

		// server must assign a new go routine for each client incoming messages
		go func() {
			for {
				// get string input from client
				m, err := r.ReadString('\n')
				if err != nil {
					log.Fatal(err, "on input messages")
				}
				// handling input
				fmt.Println(m)
			}
		}()
	}
}
