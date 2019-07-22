package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
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
		wr := bufio.NewWriter(conn)

		// server must assign a new go routine for each client incoming messages
		go func() {
			for {
				// get string input from client
				m, err := r.ReadString('\n')
				if err != nil {
					log.Fatal(err, "on input messages")
				}
				// handling input
				b := calculate(m)
				switch b.(type) {
				case int:
					wr.WriteString(strconv.Itoa(b.(int)) + "\n")
					wr.Flush()
				case float64:
					wr.WriteString(fmt.Sprintf("%.2f", b.(float64)) + "\n")
					wr.Flush()
				}
				fmt.Println(b)
			}
		}()
	}
}

func calculate(s string) interface{} {
	var firstOperator, secondOperator int
	if strings.Contains(s, "-") {
		s := strings.Split(s, "-")
		firstOperator, _ = strconv.Atoi(strings.TrimSpace(s[0]))
		secondOperator, _ = strconv.Atoi(strings.TrimSpace(s[1]))
		return firstOperator - secondOperator
	}
	if strings.Contains(s, "+") {
		s := strings.Split(s, "+")
		firstOperator, _ = strconv.Atoi(strings.TrimSpace(s[0]))
		secondOperator, _ = strconv.Atoi(strings.TrimSpace(s[1]))
		return firstOperator + secondOperator
	}
	if strings.Contains(s, "*") {
		s := strings.Split(s, "*")
		firstOperator, _ = strconv.Atoi(strings.TrimSpace(s[0]))
		secondOperator, _ = strconv.Atoi(strings.TrimSpace(s[1]))
		return firstOperator * secondOperator
	}
	if strings.Contains(s, "/") {
		s := strings.Split(s, "/")
		firstOperator, _ = strconv.Atoi(strings.TrimSpace(s[0]))
		secondOperator, _ = strconv.Atoi(strings.TrimSpace(s[1]))
		return float64(firstOperator) / float64(secondOperator)
	}
	if strings.Contains(s, "%") {
		s := strings.Split(s, "%")
		firstOperator, _ = strconv.Atoi(strings.TrimSpace(s[0]))
		secondOperator, _ = strconv.Atoi(strings.TrimSpace(s[1]))
		return firstOperator % secondOperator
	}
	return nil
}
