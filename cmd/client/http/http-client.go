package main

import (
	"Remote-Calculator/pkg/model"
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	for s, _ := r.ReadString('\n'); s != "exit\n"; s, _ = r.ReadString('\n') {
		s = strings.TrimRight(s, "\n")
		problem := model.NewProblem(s)
		buf := bytes.NewBuffer(nil)
		json.NewEncoder(buf).Encode(problem)
		res, err := http.Post("http://localhost:1234/solve", "application/json; charset=utf-8", buf)
		if err != nil {
			log.Fatal(err)
		}
		json.NewDecoder(res.Body).Decode(&problem)
		fmt.Println(problem.Result)
	}
}
