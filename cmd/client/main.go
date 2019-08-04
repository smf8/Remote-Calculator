package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	inputString, _ := in.ReadString('\n')

	// connect to server
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err, "on connection")
	}

	// close the connection after client is finished
	defer conn.Close()
	// create input / output (stream?) to server
	output := bufio.NewWriter(conn)
	input := bufio.NewReader(conn)

	// create a go routine to check for incoming messages from server
	go func() {
		for {
			i, err := input.ReadString('\n')
			if err != nil {
				panic(err)
			}
			i = strings.TrimRight(i, "\n")
			fmt.Println(i)
		}
	}()

	for inputString != "q\n" {
		// send message to server
		output.WriteString(inputString)
		output.Flush()
		inputString, _ = in.ReadString('\n')
	}
}
