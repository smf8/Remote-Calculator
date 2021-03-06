package main

import (
	"bufio"
	"net"
	"strings"
	"testing"
)

func BenchmarkCalculate(b *testing.B) {
	var str strings.Builder
	for i := 0; i < b.N; i++ {
		str.Reset()
		if (i % 4) == 0 {
			// summation
			str.WriteString(string(i))
			str.WriteString("+")
			str.WriteString(string(b.N - i))
		} else if (i % 4) == 1 {
			// substraction
			str.WriteString(string(i))
			str.WriteString("-")
			str.WriteString(string(b.N - i))
		} else if (i % 5) == 2 {
			// multiplication
			str.WriteString(string(i))
			str.WriteString("*")
			str.WriteString(string(b.N - i))
		} else {
			//division
			str.WriteString(string(b.N - i))
			str.WriteString("/")
			str.WriteString(string(i))
		}
		calculate(str.String())
	}
}
func TestCalculate(t *testing.T) {
	add := "2  +3 "
	sub := "4-1"
	multiply := " 3 * 22"
	div := "4 / 2"
	mod := " 8%6"
	if calculate(add).(int) != 5 {
		t.Error("Failed in addition")
	}
	if calculate(sub).(int) != 3 {
		t.Error("Failed in substraction")
	}
	if calculate(multiply).(int) != 66 {
		t.Error("Failed in multiplication")
	}
	if calculate(div).(float64) != 2.00 {
		t.Error("Failed in division")
	}
	if calculate(mod).(int) != 2 {
		t.Error("Failed in modulo")
	}
}

func TestServer(t *testing.T) {
	go main()
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		t.Error("Failed to connect to server ", err)
	}
	_, err1 := conn.Write([]byte("2*2\n"))
	if err1 != nil {
		t.Error("Failed to send data to server ", err1)
	}
	r := bufio.NewReader(conn)
	b, e := r.ReadString('\n')
	if e != nil {
		t.Error("Failed to read data from server ", e)
	}
	if b != "4\n" {
		t.Error("Failed in calculation")
	}
	defer conn.Close()
}
