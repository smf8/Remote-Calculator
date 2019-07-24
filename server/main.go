package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type client struct {
	conn       net.Conn
	writer     *bufio.Writer
	reader     *bufio.Reader
	income     chan string
	disconnect chan string
}

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
		// a client is connected successfully, creating a client type for it
		c := client{
			conn:   conn,
			writer: bufio.NewWriter(conn),
			reader: bufio.NewReader(conn),
			income: make(chan string),
		}
		// run listen in a indefinite go routine
		go listen(&c)

		// run respond indefinitely on another go routine
		go c.respond()
	}
}

// function to calculate a problem
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

// responding to incoming problems with their answer
func (c *client) respond() {
	for {
		select {
		case s := <-c.income:
			// message received, let's the answer back to client
			b := calculate(s)
			switch b.(type) {
			case int:
				c.writer.WriteString(strconv.Itoa(b.(int)) + "\n")
				c.writer.Flush()
			case float64:
				c.writer.WriteString(fmt.Sprintf("%.2f", b.(float64)) + "\n")
				c.writer.Flush()
			}
			fmt.Println(b)
		case <-c.disconnect:
			c.conn.Close()
			return
		}
	}
}

func listen(c *client) {
	for {
		// get string input from client
		m, err := c.reader.ReadString('\n')
		if err != nil {
			c.disconnect <- "exit"
			log.Fatal(err, " on input")
		}
		// no error, sending message to income channel
		c.income <- m
	}
}
